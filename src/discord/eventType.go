package discord

import (
	"fmt"
	"go/types"
	"strconv"
	"wall-e/src/custom"
)

type EventType int

func (e *EventType) MarshalText() (text []byte, err error) {
	str := strconv.Itoa(int(*e))
	return []byte(str), nil
}

func (e *EventType) UnmarshalText(text []byte) (err error) {
	*e, err = EventTypeFromString(string(text))
	return err
}

func (e *EventType) MarshalJSON() (text []byte, err error) {
	return e.MarshalText()
}

func (e *EventType) UnmarshalJSON(text []byte) error {
	return e.UnmarshalText(text)
}

func (e *EventType) MarshalBSON() ([]byte, error) {
	return []byte(e.String()), nil
}

func (e *EventType) UnmarshalBSON(bytes []byte) (err error) {
	*e, err = EventTypeFromString(string(bytes))
	return err
}

func (e *EventType) Scan(src any) (err error) {
	*e, err = EventTypeFromString(fmt.Sprintf("%s", src))
	return err
}

func (e *EventType) Underlying() types.Type {
	return e
}

func (e *EventType) String() string {
	values := eventTypeMap.Keys()
	return values[*e]
}

func (e *EventType) Display() string {
	return displayFormatter(e)
}

func (e *EventType) Eqauls(eventType EventType) custom.Bool {
	return *e == eventType
}

func (e *EventType) IsIn(eventTypes ...EventType) custom.Bool {
	if eventTypes == nil {
		return false
	}

	for _, opcode := range eventTypes {
		if e.Eqauls(opcode) {
			return true
		}
	}

	return false
}

var eventTypeMap = custom.Map[string, EventType]{
	"HELLO":           EventTypeHello,
	"READY":           EventTypeReady,
	"MESSAGE_CREATE":  EventTypeMessageCreate,
	"GUILD_CREATE":    EventTypeGuildCreate,
	"TYPING_START":    EventTypeTypingStart,
	"PRESENCE_UPDATE": EventTypePresenceUpdate,
}

const (
	EventTypeHello          EventType = iota // EventTypeHello is what is received on connection
	EventTypeReady                           // EventTypeReady is the initial message that is read from the websocket to justify the identity of the bot
	EventTypeMessageCreate                   // EventTypeMessageCreate is for when the bot joins a new server
	EventTypeGuildCreate                     // EventTypeGuildCreate is for when the bot joins a new guild (server). This is also received upon new startups when each guild is made available to the bot
	EventTypeTypingStart                     // EventTypeTypingStart is used for when a user starts typing
	EventTypePresenceUpdate                  // EventTypePresenceUpdate is their current state on a guild. This event is sent when a user's presence or info, such as name or avatar, is updated.
)

func EventTypeFromString(text string) (EventType, error) {
	parsedStringValue := parseText(text)

	if !eventTypeMap.Has(parsedStringValue) {
		return EventType(-1), fmt.Errorf("eventType: "+errorTextEnumNotFound, text)
	}

	return *eventTypeMap.Get(parsedStringValue), nil
}
