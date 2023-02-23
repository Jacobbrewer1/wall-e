package discord

import "log"

func (s *Session) listen() {
	log.Println("listening to new messages")

	for {
		messageType, message, err := s.connection.ReadMessage()
		if err != nil {
			log.Println("reading websocket message:", err)
			return
		}

		log.Println(messageType)
		log.Println(string(message))

		select {

		case <-CurrentSession.stop:
			log.Println("stop message received: stopped listening")
			return
		default:
			//s.onEvent(messageType, message)
		}
	}
}
