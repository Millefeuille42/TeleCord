package main

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var socket socketStruct

func main() {
	args := os.Args

	if len(args) < 3 {
		fmt.Printf("%s telegramToken discordToken", args[0])
		return
	}

	socket = socketStruct{
		telegramSession: telegramSetup(args),
		discordSession:  discordSetup(args),
		discordMessage:  nil,
		telegramMessage: tgbotapi.Update{},
	}

	setupCloseHandler(socket.discordSession)
	hang()
}

func setupCloseHandler(session *discordgo.Session) {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		fmt.Println("\r- Ctrl+C pressed in Terminal")
		time.Sleep(2 * time.Second)
		_ = session.Close()
		os.Exit(0)
	}()
}
