package bot

import (
	"log"
	"github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func Run() {
	bot, err := tgbotapi.NewBotAPI("8541911212:AAE1m1NhtGxwnCDxdHsTDdDUEiswwqSsbPg")
	if err != nil {
		log.Panic(err)
	}
	
	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
    u.Timeout = 60

    updates := bot.GetUpdatesChan(u)
	for update := range updates {
        if update.Message == nil {
            continue
        }

        log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

        // Respond to messages
        msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")

        switch update.Message.Text {
        case "/start":
            msg.Text = "Hello! I'm your bot. How can I help you?"
        case "/help":
            msg.Text = "Available commands:\n/start - Start the bot\n/help - Show this message"
        default:
            msg.Text = "You said: " + update.Message.Text
        }

        bot.Send(msg)
    }
}