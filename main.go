package main

import (
	"log"

	tgbotapi "github.com/Syfaro/telegram-bot-api"
)

func main() {

	bot, err := tgbotapi.NewBotAPI("")
	if err != nil {
		panic(err) // You should add better error handling than this!
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	bot.Debug = true // Has the library display every request and response.

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil { // ignore any non-Message Updates
			continue
		}

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
		msg.ReplyToMessageID = update.Message.MessageID

		if _, err := bot.Send(msg); err != nil {
			log.Panic(err)
		}
	}
}
