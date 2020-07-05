package main

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"syscall"
	"time"
)

var bot *tgbotapi.BotAPI

func main() {
	args := os.Args

	if len(args) < 3 {
		fmt.Printf("%s telegramToken discordToken", args[0])
		return
	}
	bot, _ = tgbotapi.NewBotAPI(args[1])
	bot.Debug = true
	fmt.Printf("Authorized on account %s\n", bot.Self.UserName)

	discordBot, err := discordgo.New("Bot " + args[2])
	checkError(err)
	fmt.Println("Discord bot created")
	discordBot.AddHandler(messageHandler)
	err = discordBot.Open()
	checkError(err)
	fmt.Println("Discord Bot up and running")

	setupCloseHandler(discordBot)
	telegramMessageHandler(bot, discordBot)
}

func setupCloseHandler(session *discordgo.Session) {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		fmt.Println("\r- Ctrl+C pressed in Terminal")
		time.Sleep(2 * time.Second)
		_ = session.Close()
		os.Exit(0)
	}()
}

func telegramMessageHandler(bot *tgbotapi.BotAPI, session *discordgo.Session) {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates, err := bot.GetUpdatesChan(u)

	checkError(err)

	for update := range updates {
		if update.Message == nil {
			continue
		}
		fmt.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
		fmt.Println(update.Message)

		if strings.HasPrefix(update.Message.Text, "/") {
			if strings.HasPrefix(update.Message.Text, "/dest") {
				arg := strings.Split(update.Message.Text, "-")
				if len(arg) > 1 {
					id, _ := strconv.Atoi(arg[1])
					err := telegramRegisterUser(update.Message.From.ID, id)
					if err != nil {
						msg := tgbotapi.NewMessage(update.Message.Chat.ID, fmt.Sprintf("An Error Occurred %s", err.Error()))
						_, _ = bot.Send(msg)
					}
				} else {
					msg := tgbotapi.NewMessage(update.Message.Chat.ID, "You must provide a UserID")
					_, _ = bot.Send(msg)
				}
			}
			if strings.HasPrefix(update.Message.Text, "/myID") {
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, fmt.Sprintf("Your Telegram ID:\n\t%d", update.Message.From.ID))
				_, _ = bot.Send(msg)
			}
			continue
		}

		err := telegramTransmitMessage(update.Message, session, bot)
		if err != nil {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, fmt.Sprintf("An Error Occurred %s", err.Error()))
			_, _ = bot.Send(msg)
		}
	}
}

func messageHandler(session *discordgo.Session, message *discordgo.MessageCreate) {

	botID, err := session.User("@me")
	checkError(err)

	if botID.ID == message.Author.ID {
		return
	}

	if strings.HasPrefix(message.Content, "/") {
		if strings.HasPrefix(message.Content, "/dest") {
			arg := strings.Split(message.Content, "-")
			if len(arg) > 1 {
				tID, _ := strconv.Atoi(arg[1])
				dID, _ := strconv.Atoi(message.Author.ID)
				err := discordRegisterUser(dID, tID)
				if err != nil {
					msg := fmt.Sprintf("An Error Occurred %s", err.Error())
					_, _ = session.ChannelMessageSend(message.ChannelID, msg)
				}
			} else {
				_, _ = session.ChannelMessageSend(message.ChannelID, "You must provide a UserID")
			}
		}
		if strings.HasPrefix(message.Content, "/myID") {
			msg := fmt.Sprintf("Your Discord ID:\n\t%s", message.Author.ID)
			_, _ = session.ChannelMessageSend(message.ChannelID, msg)
		}
		return
	}
	err = discordTransmitMessage(message, session, bot)
	if err != nil {
		msg := fmt.Sprintf("An Error Occurred %s", err.Error())
		_, _ = session.ChannelMessageSend(message.ChannelID, msg)
	}
}
