package discord

import "encoding/json"

type Event struct {
	Opcode   Opcode          `json:"op"`
	Sequence int64           `json:"s"`
	Type     *EventType      `json:"t"`
	RawData  json.RawMessage `json:"d"`
}
