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
	discord.NewSession("token here").Start()

	<-make(chan any)
}
