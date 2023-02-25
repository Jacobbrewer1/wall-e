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
	}
}
