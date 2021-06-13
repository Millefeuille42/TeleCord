package messageParser

import (
	"github.com/Millefeuille42/TeleCord/srcs/definitions"
	"strconv"
)

func TelegramParseMessage() (definitions.MessageStruct, error) {
	var attachments []string = nil
	var message definitions.MessageStruct

	if definitions.Socket.TelegramMessage.Message.Document != nil {
		attachments = make([]string, 0)
		attachment, err := definitions.Socket.TelegramSession.GetFileDirectURL(definitions.Socket.TelegramMessage.
			Message.Document.FileID)
		if err != nil {
			return definitions.MessageStruct{SenderID: uint64(definitions.Socket.TelegramMessage.Message.From.ID)}, err
		}
		attachments = append(attachments, attachment)
	}
	message = definitions.MessageStruct{
		SenderID:         uint64(definitions.Socket.TelegramMessage.Message.From.ID),
		SenderName:       definitions.Socket.TelegramMessage.Message.From.UserName,
		MessageContent:   definitions.Socket.TelegramMessage.Message.Text,
		AttachmentsLinks: attachments,
	}
	return message, nil
}

func DiscordParseMessage() (definitions.MessageStruct, error) {
	var attachments []string = nil
	var message definitions.MessageStruct

	authorID, _ := strconv.Atoi(definitions.Socket.DiscordMessage.Author.ID)
	if len(definitions.Socket.DiscordMessage.Attachments) > 0 {
		attachments = make([]string, 0)
		for _, att := range definitions.Socket.DiscordMessage.Attachments {
			attachments = append(attachments, att.URL)
		}
	}
	message = definitions.MessageStruct{
		SenderID:         uint64(authorID),
		SenderName:       definitions.Socket.DiscordMessage.Author.Username,
		MessageContent:   definitions.Socket.DiscordMessage.Content,
		AttachmentsLinks: attachments,
	}
	return message, nil
}
