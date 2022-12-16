package amesame

import (
	"github.com/bwmarrin/discordgo"
)

func ameSameHandler(s *discordgo.Session, m *discordgo.MessageCreate) {

	// 忽略來自機器人自己發的訊息，避免無限自問自答
	if m.Author.ID == s.State.User.ID {
		return
	}

	if m.Content == "Ame" {
		s.ChannelMessageSend(m.ChannelID, "Same")
	}
	if m.Content == "Same" {
		s.ChannelMessageSend(m.ChannelID, "Ame")
	}
}
