package main

import (
	"github.com/bwmarrin/discordgo"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func telegramMessageHandler(bot *tgbotapi.BotAPI) {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates, err := bot.GetUpdatesChan(u)

	checkError(err)

	for update := range updates {
		if update.Message == nil {
			continue
		}
		socket.telegramMessage = update
		handlerRouter("telegram")
	}
}

func discordMessageHandler(session *discordgo.Session, message *discordgo.MessageCreate) {

	botID, err := session.User("@me")
	checkError(err)

	if botID.ID == message.Author.ID {
		return
	}
	socket.discordMessage = message
	handlerRouter("discord")
}
