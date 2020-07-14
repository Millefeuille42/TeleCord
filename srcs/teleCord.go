package main

import (
	"fmt"
	"github.com/Millefeuille42/TeleCord/definitions"
	"github.com/Millefeuille42/TeleCord/perApp/appSetup"
	"github.com/Millefeuille42/TeleCord/utils"
	"github.com/bwmarrin/discordgo"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	args := os.Args

	if len(args) < 3 {
		fmt.Printf("%s telegramToken discordToken", args[0])
		return
	}

	definitions.Socket = definitions.SocketStruct{
		TelegramSession: appSetup.TelegramSetup(args),
		DiscordSession:  appSetup.DiscordSetup(args),
		DiscordMessage:  nil,
		TelegramMessage: tgbotapi.Update{},
	}

	setupCloseHandler(definitions.Socket.DiscordSession)
	utils.Hang()
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
