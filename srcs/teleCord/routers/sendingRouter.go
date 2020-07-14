package routers

import (
	"fmt"
	"github.com/Millefeuille42/TeleCord/perApp/messageSender"
)

func SendingRouter(platform, text string, dest int, attachments []string) error {
	attachmentsList := ""

	if attachments != nil {
		for _, att := range attachments {
			attachmentsList = fmt.Sprintf("%s\n%s", attachmentsList, att)
		}
	}
	text = fmt.Sprintf("%s\n%s", text, attachmentsList)

	switch platform {
	case "telegram":
		return messageSender.TelegramSendMessage(text, dest)
	case "discord":
		return messageSender.DiscordSendMessage(text, dest)
	}
	return nil
}
