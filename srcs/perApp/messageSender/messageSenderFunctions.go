package messageSender

import (
	"TeleCord/definitions"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func DiscordSendMessage(text string, dest int) error {
	dmChan, err := definitions.Socket.DiscordSession.UserChannelCreate(fmt.Sprintf("%d", dest))
	if err != nil {
		return err
	}
	_, err = definitions.Socket.DiscordSession.ChannelMessageSend(dmChan.ID, text)
	if err != nil {
		return err
	}
	return nil
}

func TelegramSendMessage(text string, dest int) error {
	msg := tgbotapi.NewMessage(int64(dest), text)
	_, err := definitions.Socket.TelegramSession.Send(msg)
	if err != nil {
		return err
	}
	return nil
}
