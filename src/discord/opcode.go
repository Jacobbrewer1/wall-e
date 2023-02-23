package discord

import (
	"fmt"
	"go/types"
	"strconv"
	"wall-e/src/custom"
)

type Opcode int

func (o *Opcode) MarshalText() (text []byte, err error) {
	str := strconv.Itoa(int(*o))
	return []byte(str), nil
}

func (o *Opcode) UnmarshalText(text []byte) error {
	i, err := strconv.Atoi(string(text))
	if err != nil {
		return err
	}
	*o = Opcode(i)
	return nil
}

func (o *Opcode) MarshalBSON() ([]byte, error) {
	return []byte(o.String()), nil
}

func (o *Opcode) UnmarshalBSON(bytes []byte) (err error) {
	*o, err = OpcodeFromString(string(bytes))
	return err
}

func (o *Opcode) Scan(src any) (err error) {
	*o, err = OpcodeFromString(fmt.Sprintf("%s", src))
	return err
}

func (o *Opcode) Underlying() types.Type {
	return o
}

func (o *Opcode) String() string {
	values := opcodeMap.Keys()
	return values[*o]
}

func (o *Opcode) Display() string {
	return displayFormatter(o)
}

var opcodeMap = custom.Map[string, Opcode]{
	"DISPATCH":              OpcodeDispatch,
	"HEARTBEAT":             OpcodeHeartbeat,
	"IDENTIFY":              OpcodeIdentify,
	"PRESENCE_UPDATE":       OpcodePresenceUpdate,
	"VOICE_STATE_UPDATE":    OpcodeVoiceStateUpdate,
	"RESUME":                OpcodeResume,
	"RECONNECT":             OpcodeReconnect,
	"REQUEST_GUILD_MEMBERS": OpcodeRequestGuildMembers,
	"INVALID_SESSION":       OpcodeInvalidSession,
	"HELLO":                 OpcodeHello,
	"HEARTBEAT_ACK":         OpcodeHeartbeatACK,
}

const (
	OpcodeDispatch Opcode = iota
	OpcodeHeartbeat
	OpcodeIdentify
	OpcodePresenceUpdate
	OpcodeVoiceStateUpdate
	OpcodeResume
	OpcodeReconnect
	OpcodeRequestGuildMembers
	OpcodeInvalidSession
	OpcodeHello
	OpcodeHeartbeatACK
)

func OpcodeFromString(text string) (Opcode, error) {
	parsedStringValue := parseText(text)

	if !opcodeMap.Has(parsedStringValue) {
		return Opcode(-1), fmt.Errorf(errorTextEnumNotFound, text)
	}

	return *opcodeMap.Get(parsedStringValue), nil
}
