package routers

import (
	"TeleCord/definitions"
	"TeleCord/perApp/messageParser"
)

func ParsingRouter(origin string) (definitions.MessageStruct, error) {
	switch origin {
	case "telegram":
		return messageParser.TelegramParseMessage()
	case "discord":
		return messageParser.DiscordParseMessage()
	}
	return definitions.MessageStruct{}, nil
}
