package discord

import "encoding/json"

// Event is what all events should inherit from. The structure of the events will be
// defined under the RawData part of this struct
type Event struct {
	Opcode   Opcode          `json:"op"`
	Sequence int64           `json:"s"`
	Type     *EventType      `json:"t"`
	RawData  json.RawMessage `json:"d"`
}
