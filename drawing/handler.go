package drawing

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

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
			holomen, quote := draw(uid)
			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: fmt.Sprintf("> 今日**%v**給 <@%s> 的建議是 \n > 「**%v**」", holomen, uid, quote),
				},
			})
		},
	}
)

func draw(uid string) (string, string) {
	holomen := getHolomen(uid)
	quote := getQuote(uid)
	return holomen, quote
}
