package main

import (
	"log"
	"wall-e/src/discord"
)

func init() {
	initLogging()
}

func initLogging() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
}

func main() {
	if err := discord.NewSession("token here").Start(); err != nil {
		log.Println(err)
		return
	}

	<-make(chan any)
}
