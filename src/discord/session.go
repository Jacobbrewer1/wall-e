package discord

import (
	"encoding/json"
	"github.com/Jacobbrewer1/websocket"
	"runtime"
	"sync"
	"wall-e/src/custom"
)

type Session struct {
	sync.RWMutex

	// Identify is sent during initial handshake with the discord gateway.
	// https://discord.com/developers/docs/topics/gateway#identify
	Identify Identify

	// Event handlers
	handlers custom.Map[EventType, *HandlerFunc]

	// The websocket connection.
	connection *websocket.Conn

	// stop is used to stop the threads that maintain the connection
	stop custom.Bool

	// messages is a channel that is used to allow for the system to process multiple messages at the same time, concurrency yano
	messages chan []byte

	// stores session ID of current Gateway connection
	sessionID string

	resumeGatewayUrl string
}

func (s *Session) SetBotActivity(botActivity Activity) error {
	data, err := json.Marshal(updateStatusOp{
		Op: OpcodePresenceUpdate,
		Data: UpdateStatusData{
			Activities: []*Activity{
				&botActivity,
			},
		},
	})
	if err != nil {
		return err
	}

	if err := s.connection.WriteMessage(websocket.TextMessage, data); err != nil {
		return err
	}

	return nil
}

func NewSession(token string) *Session {
	s := &Session{
		Identify: Identify{
			Token: token,
			Properties: IdentifyProperties{
				OS: runtime.GOOS,
			},
			Intents: IntentsAll,
		},
		stop:     false,
		messages: make(chan []byte, eventChannelBuffer),
	}

	return s
}
