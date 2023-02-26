package discord

import (
	"log"
)

const (
	numberOfWorkers    = 5
	eventChannelBuffer = numberOfWorkers * 5
)

func (s *Session) listen() {
	log.Println("listening to new messages")

	for i := 0; i < numberOfWorkers; i++ {
		go s.handleMessage()
	}

	for {
		_, message, err := s.connection.ReadMessage()
		if err != nil {
			log.Println("reading websocket message:", err)
			return
		}

		if s.stop {
			log.Println("listen: stop received")
			return
		}

		s.messages <- message
	}
}
