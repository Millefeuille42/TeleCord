package appSetup

import (
	"TeleCord/perApp/messageHandler"
	"TeleCord/utils"
	"fmt"
	"github.com/bwmarrin/discordgo"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func DiscordSetup(args []string) *discordgo.Session {
	discordBot, err := discordgo.New("Bot " + args[2])
	utils.CheckError(err)
	err = discordBot.Open()
	utils.CheckError(err)
	discordBot.AddHandler(messageHandler.DiscordMessageHandler)
	fmt.Printf("%-15sOK\n", "Discord:")
	return discordBot
}

func TelegramSetup(args []string) *tgbotapi.BotAPI {
	bot, _ := tgbotapi.NewBotAPI(args[1])
	bot.Debug = false
	fmt.Printf("%-15sOK\n", "Telegram:")
	go messageHandler.TelegramMessageHandler(bot)
	return bot
}
