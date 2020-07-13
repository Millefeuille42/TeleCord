package main

import (
	"github.com/bwmarrin/discordgo"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type socketStruct struct {
	telegramSession *tgbotapi.BotAPI
	telegramMessage tgbotapi.Update
	discordSession  *discordgo.Session
	discordMessage  *discordgo.MessageCreate
}

type messageStruct struct {
	senderID         int
	senderName       string
	messageContent   string
	attachmentsLinks []string
}

type contact struct {
	FromID     int    `json:"fromID"`
	ToID       int    `json:"toID"`
	ToPlatform string `json:"toPlatform"`
}
