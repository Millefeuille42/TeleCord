package main

import (
	"encoding/json"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"io/ioutil"
	"os"
)

type discordContact struct {
	DiscordID  int `json:"discordID"`
	TelegramID int `json:"telegramID"`
}

func discordRegisterUser(discordID int, telegramID int) error {
	var path = fmt.Sprintf("./data/dToT/%d.json", discordID)
	contact := discordContact{
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

func discordTransmitMessage(message *discordgo.MessageCreate, session *discordgo.Session, bot *tgbotapi.BotAPI) error {
	var path = fmt.Sprintf("./data/dToT/%s.json", message.Author.ID)

	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		_, _ = session.ChannelMessageSend(message.ChannelID, "You must register using `/dest -[telegramID]`")
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

	sendMessage := fmt.Sprintf("FROM: %s (%s)\n\t%s", message.Author.Username, message.Author.ID, message.Content)
	msg := tgbotapi.NewMessage(int64(fileDataJson.TelegramID), sendMessage)
	_, err = bot.Send(msg)
	if err != nil {
		return err
	}
	return nil
}
