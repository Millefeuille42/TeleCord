package main

import (
	"encoding/json"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"io/ioutil"
	"os"
)

type telegramContact struct {
	TelegramID	int `json:"telegramID"`
	DiscordID	int `json:"discordID"`
}

func telegramRegisterUser(telegramID int, discordID int) error {
	var path = fmt.Sprintf("./data/tToD/%d.json", telegramID)
	contact := telegramContact{
		TelegramID: telegramID,
		DiscordID:  discordID,
	}

	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		fmt.Println("\tData file not found")

		file, err := os.Create(path)
		if err != nil {
			return err
		}
		defer file.Close()

	}

	jsonData, err := json.MarshalIndent(contact, "", "\t")
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(path, jsonData, 0644)
	if err != nil {
		return err
	}
	return nil
}

func telegramTransmitMessage(message *tgbotapi.Message, session *discordgo.Session, bot *tgbotapi.BotAPI) error {
	var path = fmt.Sprintf("./data/tToD/%d.json", message.From.ID)

	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		msg := tgbotapi.NewMessage(message.Chat.ID, "You must register using `/dest -[discordID]`")
		_, _ = bot.Send(msg)
		return nil
	}

	fileData, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	fileDataJson := telegramContact{}
	err = json.Unmarshal(fileData, &fileDataJson)
	if err != nil {
		return err
	}

	dmChan, err := session.UserChannelCreate(fmt.Sprintf("%d", fileDataJson.DiscordID))
	sendMessage := fmt.Sprintf("```FROM: %s\n\t%s```", message.From.UserName, message.Text)
	if err != nil {
		return err
	}
	_, err = session.ChannelMessageSend(dmChan.ID, sendMessage)
	if err != nil {
		return err
	}
	return nil
}