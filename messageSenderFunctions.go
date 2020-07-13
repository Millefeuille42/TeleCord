package main

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func discordSendMessage(text string, dest int) error {
	dmChan, err := socket.discordSession.UserChannelCreate(fmt.Sprintf("%d", dest))
	if err != nil {
		return err
	}
	_, err = socket.discordSession.ChannelMessageSend(dmChan.ID, text)
	if err != nil {
		return err
	}
	return nil
}

func telegramSendMessage(text string, dest int) error {
	msg := tgbotapi.NewMessage(int64(dest), text)
	_, err := socket.telegramSession.Send(msg)
	if err != nil {
		return err
	}
	return nil
}
