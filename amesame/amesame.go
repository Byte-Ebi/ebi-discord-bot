/*
Package amesame 簡單的 ping pong 機器人，在聊天室聽到 Ame 就會回傳 Same；聽到 Same 則會回傳 Ame
*/
package amesame

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

var err error

func RunBot(s *discordgo.Session) {

	// 註冊 event handler
	s.AddHandler(ameSameHandler)

	// 指定要接收的事件，在這邊只需要接收訊息事件
	s.Identify.Intents = discordgo.IntentsGuildMessages

	// 開啟一個和 Discord 連線的 websocket 並開始監聽
	err = s.Open()
	if err != nil {
		log.Fatal("error opening connection,", err)
		return
	}

	/*
		建立 channel 讓程式持續執行，直到接收到指定的 Unix 訊號
		channel buffer 設定 1 保證隨時可以接到 Signal 訊號
		SIGINT:  使用者發送INTR訊號(Ctrl+C)時觸發
		SIGTERM: 終止程序例如 kill pid 時發送
		os.Interrupt = SIGINT
	*/
	log.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc
}
