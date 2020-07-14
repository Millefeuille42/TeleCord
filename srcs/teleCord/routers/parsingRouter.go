package routers

import (
	"github.com/Millefeuille42/TeleCord/definitions"
	"github.com/Millefeuille42/TeleCord/perApp/messageParser"
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
