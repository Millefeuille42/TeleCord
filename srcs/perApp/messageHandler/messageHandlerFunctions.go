package messageHandler

import (
	"TeleCord/definitions"
	"TeleCord/teleCord"
	"TeleCord/utils"
	"github.com/bwmarrin/discordgo"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func TelegramMessageHandler(bot *tgbotapi.BotAPI) {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates, err := bot.GetUpdatesChan(u)

	utils.CheckError(err)

	for update := range updates {
		if update.Message == nil {
			continue
		}
		definitions.Socket.TelegramMessage = update
		teleCord.HandleMessage("telegram")
	}
}

func DiscordMessageHandler(session *discordgo.Session, message *discordgo.MessageCreate) {

	botID, err := session.User("@me")
	utils.CheckError(err)

	if botID.ID == message.Author.ID {
		return
	}
	definitions.Socket.DiscordMessage = message
	teleCord.HandleMessage("discord")
}
