package main

import (
	"fmt"
	"strconv"
	"strings"
)

func messageParser(origin string, socket socketStruct) (messageStruct, error) {
	var attachments []string = nil
	var message messageStruct

	switch origin {
	case "telegram":
		if socket.telegramMessage.Message.Document != nil {
			attachments = make([]string, 0)
			attachment, err := bot.GetFileDirectURL(socket.telegramMessage.Message.Document.FileID)
			if err != nil {
				return messageStruct{senderID: socket.telegramMessage.Message.From.ID}, err
			}
			attachments = append(attachments, attachment)
		}
		message = messageStruct{
			senderID:         socket.telegramMessage.Message.From.ID,
			senderName:       socket.telegramMessage.Message.From.UserName,
			messageContent:   socket.telegramMessage.Message.Text,
			attachmentsLinks: attachments,
		}
	case "discord":
		authorID, _ := strconv.Atoi(socket.discordMessage.Author.ID)
		if len(socket.discordMessage.Attachments) > 0 {
			attachments = make([]string, 0)
			for _, att := range socket.discordMessage.Attachments {
				attachments = append(attachments, att.URL)
			}
		}
		message = messageStruct{
			senderID:         authorID,
			senderName:       socket.discordMessage.Author.Username,
			messageContent:   socket.discordMessage.Content,
			attachmentsLinks: attachments,
		}
	}
	return message, nil
}

func handlerRouter(origin string, socket socketStruct) {
	message, err := messageParser(origin, socket)
	if err != nil {
		_ = sendMessage(origin, err.Error(), message.senderID, socket, nil)
		return
	}

	if strings.HasPrefix(message.messageContent, "/") {
		var err error = nil

		if strings.HasPrefix(message.messageContent, "/dest") {
			err = register(origin, message, socket)
		}
		if strings.HasPrefix(message.messageContent, "/myID") {
			err = sendMessage(origin, fmt.Sprintf("Your ID:\n\t%d", message.senderID), message.senderID, socket, nil)
		}
		if err != nil {
			_ = sendMessage(origin, err.Error(), message.senderID, socket, nil)
		}
		return
	}
	err = transmitMessage(origin, message, socket)
	if err != nil {
		_ = sendMessage(origin, err.Error(), message.senderID, socket, nil)
	}
}
