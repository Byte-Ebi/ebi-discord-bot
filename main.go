package main

import (
	"log"
	"os"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"

	"ebi-discord-bot/amesame"
)

var BotToken string

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}

	BotToken = os.Getenv("DISCORD_BOT_TOKEN")
}

func main() {
	dg, err := discordgo.New("Bot " + BotToken)
	if err != nil {
		log.Fatal("error creating Discord session,", err)
		return
	}

	defer dg.Close()

	amesame.RunBot(dg)
}
