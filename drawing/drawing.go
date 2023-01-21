/*
Package drawing 透過 slash command 呼叫指令進行無意義的每日抽籤
*/

package drawing

import (
	"log"
	"os"
	"os/signal"

	"github.com/bwmarrin/discordgo"
)

var (
	err     error
	guildID string
)

func SetGuildID(s string) {
	guildID = s
}

func RunBot(s *discordgo.Session) {
	err = s.Open()
	if err != nil {
		log.Fatalf("[error] Cannot open the session: %v", err)
	}
	defer s.Close()

	addCommands(s)

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	log.Println("[info] Press Ctrl+C to exit")
	<-stop

	removeCommands(s)
	log.Println("[info] Gracefully shutting down.")
}

func addCommands(s *discordgo.Session) {
	log.Println("[info] Adding commands...")
	registeredCommands := make([]*discordgo.ApplicationCommand, len(commands))
	for i, v := range commands {
		cmd, err := s.ApplicationCommandCreate(s.State.User.ID, guildID, v)
		if err != nil {
			log.Panicf("[error] Cannot create '%v' command: %v", v.Name, err)
		}
		registeredCommands[i] = cmd
	}
	s.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if h, ok := commandHandlers[i.ApplicationCommandData().Name]; ok {
			h(s, i)
		}
	})
}

func removeCommands(s *discordgo.Session) {
	if RemoveCommands {
		log.Println("[info] Removing commands...")
		registeredCommands, err := s.ApplicationCommands(s.State.User.ID, guildID)
		if err != nil {
			log.Fatalf("[error] Could not fetch registered commands: %v", err)
		}

		for _, v := range registeredCommands {
			err := s.ApplicationCommandDelete(s.State.User.ID, guildID, v.ID)
			if err != nil {
				log.Panicf("[error] Cannot delete '%v' command: %v", v.Name, err)
			}
		}
	}
}
