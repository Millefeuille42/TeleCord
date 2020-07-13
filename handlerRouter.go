package main

import (
	"fmt"
	"strings"
)

func messageParser(origin string) (messageStruct, error) {
	switch origin {
	case "telegram":
		return telegramParseMessage(socket)
	case "discord":
		return discordParseMessage(socket)
	}
	return messageStruct{}, nil
}

func handlerRouter(origin string) {
	message, err := messageParser(origin)
	if err != nil {
		_ = sendMessage(origin, err.Error(), message.senderID, nil)
		return
	}

	if strings.HasPrefix(message.messageContent, "/") {
		var err error = nil

		if strings.HasPrefix(message.messageContent, "/dest") {
			err = register(origin, message)
		}
		if strings.HasPrefix(message.messageContent, "/myID") {
			err = sendMessage(origin, fmt.Sprintf("Your ID:\n\t%d", message.senderID), message.senderID, nil)
		}
		if strings.HasPrefix(message.messageContent, "/myDest") {
			err = getDest(origin, message)
		}
		if err != nil {
			_ = sendMessage(origin, err.Error(), message.senderID, nil)
		}
		return
	}

	err = transmitMessage(origin, message)
	if err != nil {
		_ = sendMessage(origin, err.Error(), message.senderID, nil)
	}
}
