package main

import (
	"ebi-discord-bot/timemachine"
	"fmt"
	"log"
	"os"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

var logo = `
 ____          _          _____  _      _ 
| __ )  _   _ | |_  ___  | ____|| |__  (_)
|  _ \ | | | || __|/ _ \ |  _|  | '_ \ | |
| |_) || |_| || |_|  __/ | |___ | |_) || |
|____/  \__, | \__|\___| |_____||_.__/ |_|
        |___/                             
`

var (
	BotToken string
	GuildId  string
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}

	BotToken = os.Getenv("DISCORD_BOT_TOKEN")
	GuildId = os.Getenv("GUILD_ID")
}

func main() {
	dg, err := discordgo.New("Bot " + BotToken)
	if err != nil {
		log.Fatal("error creating Discord session,", err)
		return
	}
	fmt.Println(logo)

	defer dg.Close()

	timemachine.SetBotToken(BotToken)
	timemachine.SetGuildID(GuildId)
	timemachine.RunBot(dg)
}
