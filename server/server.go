package server

import (
	"encoding/json"
	"errors"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"ocgcore"
	"time"
)

type recvMessage struct {
	c *Client
	m []byte
}

type Server struct {
	config     Config
	unregister chan *Client
	register   chan *Client
	receive    chan recvMessage
	clients    map[*Client]bool

	core  *ocgcore.OcgCore
	duels map[*Client]duelInfo
}

type Config struct {
	Address      string
	CorePath     string
	CardReader   ocgcore.CardReader
	ScriptReader ocgcore.ScriptReader
}

func NewServer(c Config) *Server {
	core, err := ocgcore.NewOcgCore(c.CorePath)
	if err != nil {
		panic(err)
	}
	return &Server{
		core:       core,
		config:     c,
		unregister: make(chan *Client),
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

func (s *Server) runServer() {
	for {
		select {
		case client := <-s.register:
			s.clients[client] = true

		case client := <-s.unregister:
			if _, ok := s.clients[client]; ok {
				delete(s.clients, client)
				close(client.send)
			}
			if d, ok := s.duels[client]; ok {
				delete(s.duels, client)
				if d.done != nil {
					close(d.done)
				}
			}

		case msg := <-s.receive:
			c := msg.c
			var m jsonMessage
			err := json.Unmarshal(msg.m, &m)
			if err != nil {
				s.unregister <- c
				break
			}

			switch m.Action {
			case "create_duel":
				duel := s.core.CreateDuel(ocgcore.CreateDuelOptions{
					Seed:         0,
					Mode:         ocgcore.DuelModeMR5,
					CardReader:   s.config.CardReader,
					ScriptReader: s.config.ScriptReader,
				})
				s.createDuel(c, duel)
				_ = s.sendClient(c, "create_duel", resultDuelCreation{Success: true})

			case "start_duel":
				err := s.startDuel(c)
				if err != nil {
					s.unregister <- c
					break
				}

			case "duel_response":
				//s.duelResponse(m.Payload)

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
				_ = s.sendClient(c, "message", m1)
			case _, _ = <-duel.done:
				break outer
			}
		}
		close(duel.done)
		duel.duel.Destroy()
	}()
	return nil
}

func (s *Server) duelResponse(c *Client, payload json.RawMessage) {

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
