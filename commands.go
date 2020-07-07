package main

import (
	"encoding/json"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func register(origin string, message messageStruct, socket socketStruct) error {
	path := fmt.Sprintf("./data/%s/%d.json", origin, message.senderID)
	arg := strings.Split(message.messageContent, "-")
	if len(arg) <= 2 {
		_ = sendMessage(origin, "Invalid number of arguments", message.senderID, socket, nil)
		return nil
	}

	arg[1] = strings.TrimSpace(arg[1])
	dest, _ := strconv.Atoi(arg[1])
	contact := contact{
		FromID:     message.senderID,
		ToID:       dest,
		ToPlatform: arg[2],
	}
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
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
	err = sendMessage(origin, fmt.Sprintf("You are now talking to %s on %s", arg[1], arg[2]),
		message.senderID, socket, nil)
	if err != nil {
		return err
	}
	return nil
}

func sendMessage(platform, text string, dest int, socket socketStruct, attachments []string) error {
	attachmentsList := ""

	if attachments != nil {
		for _, att := range attachments {
			attachmentsList = fmt.Sprintf("%s\n%s", attachmentsList, att)
		}
	}
	text = fmt.Sprintf("%s\n%s", text, attachmentsList)

	switch platform {
	case "telegram":
		msg := tgbotapi.NewMessage(int64(dest), text)
		_, err := bot.Send(msg)
		if err != nil {
			return err
		}
	case "discord":
		dmChan, err := socket.discordSession.UserChannelCreate(fmt.Sprintf("%d", dest))
		if err != nil {
			return err
		}
		_, err = socket.discordSession.ChannelMessageSend(dmChan.ID, text)
		if err != nil {
			return err
		}
	}
	return nil
}

func transmitMessage(origin string, message messageStruct, socket socketStruct) error {
	var path = fmt.Sprintf("./data/%s/%d.json", origin, message.senderID)

	message.messageContent = fmt.Sprintf("FROM: %s (%d - %s)\n\t%s", message.senderName, message.senderID,
		origin, message.messageContent)

	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		_ = sendMessage(origin, "You must register using /dest -ID -Platform", message.senderID, socket, nil)
		return nil
	}
	fileData, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	fileDataJson := contact{}
	err = json.Unmarshal(fileData, &fileDataJson)
	if err != nil {
		return err
	}
	if fileDataJson.ToPlatform == "discord" {
		message.messageContent = fmt.Sprintf("```%s```", message.messageContent)
	}
	err = sendMessage(fileDataJson.ToPlatform, message.messageContent, fileDataJson.ToID, socket, message.attachmentsLinks)
	if err != nil {
		return err
	}
	return nil
}
