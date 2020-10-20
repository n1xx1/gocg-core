package server

import (
	"encoding/json"
	"errors"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"ocgcore"
	"ocgcore/database"
	"time"
)

type recvMessage struct {
	c *Client
	m []byte
}

type unregisterErr struct {
	c *Client
	e error
}

type Server struct {
	config     Config
	unregister chan unregisterErr
	register   chan *Client
	receive    chan recvMessage
	clients    map[*Client]bool
	duels      map[*Client]duelInfo
}

type Config struct {
	Address      string
	ScriptReader ocgcore.ScriptReader
	Database     database.CardDatabase
}

func NewServer(c Config) *Server {
	return &Server{
		config:     c,
		unregister: make(chan unregisterErr),
		register:   make(chan *Client),
		receive:    make(chan recvMessage),
		clients:    map[*Client]bool{},
	}
}

func (s *Server) Run() error {
	go s.runServer()

	mux := http.NewServeMux()
	mux.HandleFunc("/", s.serveWebsocket)
	return http.ListenAndServe(s.config.Address, mux)
}

type duelInfo struct {
	duel *ocgcore.OcgDuel
	done chan bool
}

type jsonMessage struct {
	Action  string          `json:"action"`
	Payload json.RawMessage `json:"payload"`
}

func (s *Server) sendClient(c *Client, action string, v interface{}) error {
	data, err := json.Marshal(v)
	if err != nil {
		return err
	}
	data2, err := json.Marshal(jsonMessage{
		Action:  action,
		Payload: data,
	})
	if err != nil {
		return err
	}

	c.send <- data2
	return nil
}

func (s *Server) kickClient(c *Client, err error) {
	s.unregister <- unregisterErr{
		c: c,
		e: err,
	}
}

func (s *Server) sendClientRaw(c *Client, action string, data json.RawMessage) error {
	data2, err := json.Marshal(jsonMessage{
		Action:  action,
		Payload: data,
	})
	if err != nil {
		return err
	}

	c.send <- data2
	return nil
}

func (s *Server) runServer() {
	for {
		select {
		case client := <-s.register:
			s.clients[client] = true

		case u := <-s.unregister:
			if _, ok := s.clients[u.c]; ok {
				log.Print("kicking client ", u.e)

				_ = s.destroyDuel(u.c)
				_ = u.c.conn.Close()
				delete(s.clients, u.c)
				close(u.c.send)
			}

		case msg := <-s.receive:
			c := msg.c
			var m jsonMessage
			err := json.Unmarshal(msg.m, &m)
			if err != nil {
				s.kickClient(c, err)
				break
			}

			switch m.Action {
			case "card":
				var msg messageCard
				if err := json.Unmarshal(m.Payload, &msg); err != nil {
					s.kickClient(c, err)
					break
				}
				_ = s.sendClient(c, "card", s.config.Database[msg.Card].Card)

			case "create_duel":
				duel := ocgcore.CreateDuel(ocgcore.CreateDuelOptions{
					Seed: 0,
					Mode: ocgcore.DuelModeMR5,
					CardReader: func(code uint32) (raw ocgcore.RawCardData) {
						if card, ok := s.config.Database[code]; ok {
							raw = card.Raw
						}
						return
					},
					ScriptReader: s.config.ScriptReader,
				})
				s.createDuel(c, duel)
				_ = s.sendClient(c, "create_duel", resultDuelCreation{Success: true})
			case "start_duel":
				err := s.startDuel(c)
				if err != nil {
					s.kickClient(c, err)
					break
				}
				// TODO: feedback
			case "duel_response":
				err := s.duelResponse(c, m.Payload)
				if err != nil {
					s.kickClient(c, err)
					break
				}
				// TODO: feedback
			case "quit_duel":
				err := s.destroyDuel(c)
				if err != nil {
					s.kickClient(c, err)
					break
				}
				// TODO: feedback
			}
		}
	}
}

func (s *Server) createDuel(c *Client, duel *ocgcore.OcgDuel) {
	s.duels[c] = duelInfo{duel: duel}
}

func (s *Server) startDuel(c *Client) error {
	duel, ok := s.duels[c]
	if !ok {
		return errors.New("first create the duel")
	}

	duel.done = make(chan bool)
	messages := duel.duel.Start()
	go func() {
	outer:
		for {
			select {
			case m1, ok := <-messages:
				if !ok {
					break outer
				}

				m2, _ := ocgcore.MessageToJSON(m1)
				err := s.sendClientRaw(c, "message", m2)

				if err != nil {
					s.kickClient(c, err)
					break outer
				}
			case _, _ = <-duel.done:
				break outer
			}
		}
		_ = s.destroyDuel(c)
	}()
	return nil
}

func (s *Server) destroyDuel(c *Client) error {
	duel, ok := s.duels[c]
	if !ok {
		return errors.New("first create the duel")
	}

	close(duel.done)
	duel.duel.Destroy()
	return nil
}

func (s *Server) duelResponse(c *Client, payload json.RawMessage) error {
	duel, ok := s.duels[c]
	if !ok {
		return errors.New("first create the duel")
	}
	resp, err := ocgcore.JSONToResponse(payload)
	if err != nil {
		return err
	}
	duel.duel.SendResponse(resp)
	return nil
}

func (s *Server) serveWebsocket(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	client := &Client{
		server: s,
		conn:   conn,
		send:   make(chan []byte),
	}
	client.server.register <- client

	go client.writePump()
	go client.readPump()
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

const (
	writeWait      = 10 * time.Second
	pongWait       = 60 * time.Second
	pingPeriod     = (pongWait * 9) / 10
	maxMessageSize = 512
)
