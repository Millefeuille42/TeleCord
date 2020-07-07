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

var bot *tgbotapi.BotAPI

func main() {
	args := os.Args

	if len(args) < 3 {
		fmt.Printf("%s telegramToken discordToken", args[0])
		return
	}
	bot, _ = tgbotapi.NewBotAPI(args[1])
	bot.Debug = true
	fmt.Printf("Authorized on account %s\n", bot.Self.UserName)

	discordBot, err := discordgo.New("Bot " + args[2])
	checkError(err)
	fmt.Println("Discord bot created")
	discordBot.AddHandler(messageHandler)
	err = discordBot.Open()
	checkError(err)
	fmt.Println("Discord Bot up and running")

	setupCloseHandler(discordBot)
	telegramMessageHandler(bot, discordBot)
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

func telegramMessageHandler(bot *tgbotapi.BotAPI, session *discordgo.Session) {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates, err := bot.GetUpdatesChan(u)

	checkError(err)

	for update := range updates {
		if update.Message == nil {
			continue
		}
		socket := socketStruct{
			telegramMessage: update,
			discordSession:  session,
			discordMessage:  &discordgo.MessageCreate{},
		}
		handlerRouter("telegram", socket)
	}
}

func messageHandler(session *discordgo.Session, message *discordgo.MessageCreate) {

	botID, err := session.User("@me")
	checkError(err)

	if botID.ID == message.Author.ID {
		return
	}
	socket := socketStruct{
		telegramMessage: tgbotapi.Update{},
		discordSession:  session,
		discordMessage:  message,
	}
	handlerRouter("discord", socket)
}
