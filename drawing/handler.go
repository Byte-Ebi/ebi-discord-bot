package drawing

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/bwmarrin/discordgo"
)

var (
	seed int64
)

func setSeed(s string) {
	i, err := strconv.Atoi(s)
	if err != nil {
		log.Fatalf("[error] Contvert user id: %v", err)
		return
	}
	seed = int64(i) + int64(time.Now().YearDay())
}

var (
	// 註冊指令
	commands = []*discordgo.ApplicationCommand{
		{
			Name:        "about",
			Description: "關於這個 Bot",
		},
		{
			Name:        "drawing",
			Description: "抽一支籤",
		},
	}

	// 指令對應的處理方式
	commandHandlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
		"about": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: "沒什麼用的抽籤，每天 0:00(UTC+8) 重置",
				},
			})
		},
		"drawing": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			uid := i.Member.User.ID
			holomen := draw(uid)
			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: fmt.Sprintf("> <@%s> 抽到了: %v\n", uid, holomen),
				},
			})
		},
	}
)

func draw(uid string) string {
	setSeed(uid)
	holomen := getHolomen(seed)
	return holomen
}
