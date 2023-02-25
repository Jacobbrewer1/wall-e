package discord

import "fmt"

type HandlerFunc func(any)

func (s *Session) AddHandler(eventType EventType, handler HandlerFunc) error {
	if s.handlers.Has(eventType) {
		return fmt.Errorf("eventType %s already defined", eventType)
	}

	s.handlers.Set(eventType, &handler)
	return nil
}
