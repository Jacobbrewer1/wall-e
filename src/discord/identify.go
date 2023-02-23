package discord

import (
	"encoding/json"
	"github.com/Jacobbrewer1/websocket"
	"log"
)

type IdentifyOperation struct {
	Op   Opcode   `json:"op"`
	Data Identify `json:"d"`
}

type Identify struct {
	Token          string              `json:"token"`
	Properties     IdentifyProperties  `json:"properties"`
	Compress       bool                `json:"compress"`
	LargeThreshold int                 `json:"large_threshold"`
	Shard          *[2]int             `json:"shard,omitempty"`
	Presence       GatewayStatusUpdate `json:"presence,omitempty"`
	Intents        Intent              `json:"intents"`
}

type IdentifyProperties struct {
	OS              string `json:"os"`
	Browser         string `json:"browser"`
	Device          string `json:"device"`
	Referer         string `json:"referer"`
	ReferringDomain string `json:"referring_domain"`
}

type GatewayStatusUpdate struct {
	Since  int    `json:"since"`
	Status string `json:"status"`
	AFK    bool   `json:"afk"`
}

func (s *Session) identify() {
	log.Println("identifying")

	op := IdentifyOperation{
		Op:   OpcodeIdentify,
		Data: s.Identify,
	}

	data, _ := json.Marshal(op)

	if err := s.connection.WriteMessage(websocket.TextMessage, data); err != nil {
		log.Println(err)
		s.stop <- struct{}{}
		return
	}

	log.Println("identified")
}
