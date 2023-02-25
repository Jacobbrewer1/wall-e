package discord

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/Jacobbrewer1/websocket"
	"log"
	"net/http"
	"time"
	"wall-e/src/custom"
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
	log.Println("bot starting")

	header := http.Header{}
	header.Add("accept-encoding", "zlib")
	wsApi, _, err := websocket.DefaultDialer.DialContext(context.Background(), disocrdWebsocketUrl, header)
	if err != nil {
		return err
	}

	s.connection = wsApi
	_, msg, err := s.connection.ReadMessage()
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

	if err := s.InitialRead(); err != nil {
		return err
	}

	go s.heartbeat(h.HeartbeatInterval)
	go s.listen()

	return nil
}

func (s *Session) resume() {
	// todo : implement this
	panic("not implemented yet")
}

type readyOp struct {
	Type   EventType `json:"t"`
	S      int       `json:"s"`
	Opcode Opcode    `json:"op"`
	Data   struct {
		V            int `json:"v"`
		UserSettings any `json:"user_settings"`
		User         struct {
			Verified      bool        `json:"verified"`
			Username      string      `json:"username"`
			MfaEnabled    bool        `json:"mfa_enabled"`
			Id            string      `json:"id"`
			Flags         int         `json:"flags"`
			Email         string      `json:"email"`
			DisplayName   string      `json:"display_name"`
			Discriminator string      `json:"discriminator"`
			Bot           custom.Bool `json:"bot"`
			Avatar        string      `json:"avatar"`
		} `json:"user"`
		SessionType      string `json:"session_type"`
		SessionId        string `json:"session_id"`
		ResumeGatewayUrl string `json:"resume_gateway_url"`
		Relationships    []any  `json:"relationships"`
		PrivateChannels  []any  `json:"private_channels"`
		Presences        []any  `json:"presences"`
		Guilds           []struct {
			Unavailable bool   `json:"unavailable"`
			Id          string `json:"id"`
		} `json:"guilds"`
		GuildJoinRequests    []any    `json:"guild_join_requests"`
		GeoOrderedRtcRegions []string `json:"geo_ordered_rtc_regions"`
		Application          struct {
			Id    string `json:"id"`
			Flags int    `json:"flags"`
		} `json:"application"`
		Trace []string `json:"_trace"`
	} `json:"d"`
}

func (s *Session) InitialRead() error {
	_, startup, err := s.connection.ReadMessage()
	if err != nil {
		return err
	}

	log.Println(string(startup))

	var readyEvent readyOp
	if err := json.Unmarshal(startup, &readyEvent); err != nil {
		return err
	}

	if readyEvent.Opcode != OpcodeDispatch {
		return fmt.Errorf("unexpected opcode received: expected 0, got %d", readyEvent.Opcode)
	} else if !readyEvent.Type.Eqauls(EventTypeReady) {
		return fmt.Errorf("unexpected event type: got %s, expected %s", EventTypeHello, EventTypeReady)
	}

	s.sessionID = readyEvent.Data.SessionId
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
				s.errChan <- err
				return
			}

			log.Println("ping success")
		case <-s.stop:
			log.Println("stop received, stopping heartbeat")
			return
		}
	}
}
