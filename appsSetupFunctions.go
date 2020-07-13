package main

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func discordSetup(args []string) *discordgo.Session {
	discordBot, err := discordgo.New("Bot " + args[2])
	checkError(err)
	err = discordBot.Open()
	checkError(err)
	discordBot.AddHandler(discordMessageHandler)
	fmt.Printf("%-15sOK\n", "Discord:")
	return discordBot
}

func telegramSetup(args []string) *tgbotapi.BotAPI {
	bot, _ := tgbotapi.NewBotAPI(args[1])
	bot.Debug = false
	fmt.Printf("%-15sOK\n", "Telegram:")
	go telegramMessageHandler(bot)
	return bot
}
