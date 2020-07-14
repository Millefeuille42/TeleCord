package definitions

import (
	"github.com/bwmarrin/discordgo"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type SocketStruct struct {
	TelegramSession *tgbotapi.BotAPI
	TelegramMessage tgbotapi.Update
	DiscordSession  *discordgo.Session
	DiscordMessage  *discordgo.MessageCreate
}

type MessageStruct struct {
	SenderID         int
	SenderName       string
	MessageContent   string
	AttachmentsLinks []string
}

type Contact struct {
	FromID     int    `json:"fromID"`
	ToID       int    `json:"toID"`
	ToPlatform string `json:"toPlatform"`
	ToName     string `json:"toName"`
}
