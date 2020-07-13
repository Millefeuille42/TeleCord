package main

import "strconv"

func telegramParseMessage(socket socketStruct) (messageStruct, error) {
	var attachments []string = nil
	var message messageStruct

	if socket.telegramMessage.Message.Document != nil {
		attachments = make([]string, 0)
		attachment, err := socket.telegramSession.GetFileDirectURL(socket.telegramMessage.Message.Document.FileID)
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
	return message, nil
}

func discordParseMessage(socket socketStruct) (messageStruct, error) {
	var attachments []string = nil
	var message messageStruct

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
	return message, nil
}
