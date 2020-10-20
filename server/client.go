package server

import (
	"errors"
	"github.com/gorilla/websocket"
	"log"
	"time"
)

type Client struct {
	server *Server
	conn   *websocket.Conn
	send   chan []byte
}

func (c *Client) readPump() {
	var err error

	defer func() {
		c.server.kickClient(c, err)
	}()

	c.conn.SetReadLimit(maxMessageSize)

	_ = c.conn.SetReadDeadline(time.Now().Add(pongWait))

	c.conn.SetPongHandler(func(string) error {
		_ = c.conn.SetReadDeadline(time.Now().Add(pongWait))
		return nil
	})

	for {
		var message []byte
		_, message, err = c.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
				return
			}
			err = nil
			return
		}
		c.server.receive <- recvMessage{c, message}
	}
}

func (c *Client) writePump() {
	var err error
	ticker := time.NewTicker(pingPeriod)

	defer func() {
		ticker.Stop()
		c.server.kickClient(c, err)
	}()

	for {
		select {
		case message, ok := <-c.send:
			_ = c.conn.SetWriteDeadline(time.Now().Add(writeWait))

			if !ok {
				// The hub closed the channel.
				_ = c.conn.WriteMessage(websocket.CloseMessage, []byte{})
				err = errors.New("close")
				return
			}

			err = c.conn.WriteMessage(websocket.TextMessage, message)
			if err != nil {
				return
			}

		case <-ticker.C:
			_ = c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err = c.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}
