package discord

import (
	"github.com/Jacobbrewer1/websocket"
	"runtime"
	"sync"
)

var CurrentSession *Session

type eventHandler func()

type Session struct {
	sync.RWMutex

	// Identify is sent during initial handshake with the discord gateway.
	// https://discord.com/developers/docs/topics/gateway#identify
	Identify Identify

	// Event handlers
	handlers map[EventHandlerType]*eventHandler

	// The websocket connection.
	connection *websocket.Conn

	// stop is used to stop the threads that maintain the connection
	stop chan interface{}

	// stores session ID of current Gateway connection
	sessionID string
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
	}

	CurrentSession = s
	return s
}