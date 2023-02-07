package main

import (
	"ebi-discord-bot/drawing"
	"fmt"
	"log"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
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

	viper.SetConfigFile(".env")
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Fatal("error .env file not found.")
		} else {
			log.Fatal("error: ", err)
		}
	}

	BotToken = viper.GetString("DISCORD_BOT_TOKEN")
	GuildId = viper.GetString("GUILD_ID")
}

func main() {
	fmt.Println(logo)
	session, err := discordgo.New("Bot " + BotToken)
	if err != nil {
		log.Fatal("error creating Discord session,", err)
		return
	}
	defer session.Close()

	drawingBot(session)
}

// func amesamaBot(session *discordgo.Session) {
// 	amesame.RunBot(session)
// }

func drawingBot(session *discordgo.Session) {
	drawing.SetGuildID(GuildId)
	drawing.RunBot(session)
}
