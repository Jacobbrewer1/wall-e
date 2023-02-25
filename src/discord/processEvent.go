package discord

import (
	"encoding/json"
	"log"
)

func (s *Session) handleMessage() {
	for msg := range s.messages {
		var event Event
		if err := json.Unmarshal(msg, &event); err != nil {
			log.Println(err)
			continue
		}

		if event.Opcode.Eqauls(OpcodeHeartbeatACK) {
			// Heartbeat response, do nothing
			continue
		}

		log.Println(string(msg))

		if event.Type == nil {
			log.Println("event type is nil, cannot process")
			continue
		}

		if !s.handlers.Has(*event.Type) {
			// handler for event not defined
			continue
		}

		handler := *s.handlers[*event.Type]
		// todo : unmarshal event here and then process into the handler defined above
		handler(event.RawData)
	}
}
