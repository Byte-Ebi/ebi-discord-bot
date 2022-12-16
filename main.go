package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

var Token string

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	Token = os.Getenv("DISCORD_BOT_TOKEN")
}

func main() {
	fmt.Println(Token)
}
