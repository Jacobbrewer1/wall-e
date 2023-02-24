package discord

import (
	"context"
	"encoding/json"
	"github.com/Jacobbrewer1/websocket"
	"log"
	"net/http"
	"time"
)

type heartbeatOp struct {
	Op   Opcode `json:"op"`
	Data int64  `json:"d"`
}

type helloOp struct {
	HeartbeatInterval time.Duration `json:"heartbeat_interval"`
}

const disocrdWebsocketUrl = "wss://gateway.discord.gg"

func (s *Session) Start() error {
	header := http.Header{}
	header.Add("accept-encoding", "zlib")
	wsApi, _, err := websocket.DefaultDialer.DialContext(context.Background(), disocrdWebsocketUrl, header)
	if err != nil {
		return err
	}

	s.connection = wsApi
	_, msg, err := CurrentSession.connection.ReadMessage()
	if err != nil {
		return err
	}

	var e Event
	if err := json.Unmarshal(msg, &e); err != nil {
		return err
	}

	var h helloOp
	if err := json.Unmarshal(e.RawData, &h); err != nil {
		return err
	}

	if err := s.identify(); err != nil {
		s.stop <- struct{}{}
		return err
	}

	// todo : read initial message here

	go s.heartbeat(h.HeartbeatInterval)
	go s.listen()

	return nil
}

func (s *Session) heartbeat(interval time.Duration) {
	heartbeatJson, _ := json.Marshal(heartbeatOp{
		Op: OpcodeHeartbeat,
	})

	ticker := time.NewTicker(interval * time.Millisecond)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			if err := s.connection.WriteMessage(websocket.TextMessage, heartbeatJson); err != nil {
				// Most likely connection closed, reconnect
				log.Println("websocket error:", err)
				log.Println("reconnecting")
				if err := CurrentSession.Start(); err != nil {
					log.Println(err)
				}
				return
			}

			log.Println("ping success")
		case <-s.stop:
			log.Println("stop received, stopping heartbeat")
			return
		}
	}
}
