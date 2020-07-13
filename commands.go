package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func register(origin string, message messageStruct) error {
	path := fmt.Sprintf("./data/%s/%d.json", origin, message.senderID)
	arg := strings.Split(message.messageContent, "-")
	if len(arg) <= 3 {
		_ = sendMessage(origin, "Invalid number of arguments", message.senderID, nil)
		return nil
	}

	arg[1] = strings.TrimSpace(arg[1])
	arg[2] = strings.TrimSpace(arg[2])
	arg[2] = strings.ToLower(arg[2])
	arg[3] = strings.TrimSpace(arg[3])
	dest, _ := strconv.Atoi(arg[1])
	contact := contact{
		FromID:     message.senderID,
		ToID:       dest,
		ToPlatform: arg[2],
		ToName:     arg[3],
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
	err = sendMessage(origin, fmt.Sprintf("You are now talking to %s (%s) on %s", arg[3], arg[1], arg[2]),
		message.senderID, nil)
	if err != nil {
		return err
	}
	return nil
}

func getDest(origin string, message messageStruct) error {
	var path = fmt.Sprintf("./data/%s/%d.json", origin, message.senderID)

	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		_ = sendMessage(origin, "You must register using /dest -ID -Platform -Nickname",
			message.senderID, nil)
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

	response := fmt.Sprintf("Currently talking to %s (%d) on %s",
		fileDataJson.ToName, fileDataJson.ToID, fileDataJson.ToPlatform)
	err = sendMessage(origin, response, message.senderID, nil)
	if err != nil {
		return err
	}
	return err
}

func sendMessage(platform, text string, dest int, attachments []string) error {
	attachmentsList := ""

	if attachments != nil {
		for _, att := range attachments {
			attachmentsList = fmt.Sprintf("%s\n%s", attachmentsList, att)
		}
	}
	text = fmt.Sprintf("%s\n%s", text, attachmentsList)

	switch platform {
	case "telegram":
		return telegramSendMessage(text, dest)
	case "discord":
		return discordSendMessage(text, dest)
	}
	return nil
}

func transmitMessage(origin string, message messageStruct) error {
	var path = fmt.Sprintf("./data/%s/%d.json", origin, message.senderID)

	message.messageContent = fmt.Sprintf("FROM: %s (%d - %s)\n\t%s", message.senderName, message.senderID,
		origin, message.messageContent)

	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		_ = sendMessage(origin, "You must register using /dest -ID -Platform -Nickname",
			message.senderID, nil)
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
	err = sendMessage(fileDataJson.ToPlatform, message.messageContent, fileDataJson.ToID, message.attachmentsLinks)
	if err != nil {
		return err
	}
	return nil
}
