package main

import (
	"fmt"
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
	bot := discord.NewSession("token here")
	if err := bot.Start(); err != nil {
		log.Println(err)
		return
	}

	err := bot.SetBotActivity(discord.Activity{
		Name: fmt.Sprintf("Satisfactory"),
		Type: discord.ActivityTypeWatching,
		URL:  "",
	})
	if err != nil {
		log.Println(err)
		return
	}

	<-make(chan any)
}
