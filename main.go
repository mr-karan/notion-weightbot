package main

import (
	"log"
	"os"
	"strconv"
	"time"

	tb "gopkg.in/tucnak/telebot.v2"
)

func main() {
	// Check for required env vars.
	token, ok := os.LookupEnv("TELEGRAM_BOT_TOKEN")
	if !ok {
		log.Fatal("TELEGRAM_BOT_TOKEN not set")
	}
	path, ok := os.LookupEnv("WEIGHTBOT_CSV_FILE")
	if !ok {
		log.Fatal("WEIGHTBOT_CSV_FILE not set")
	}
	dbID, ok := os.LookupEnv("NOTION_DB_ID")
	if !ok {
		log.Fatal("NOTION_DB_ID not set")
	}
	bearerToken, ok := os.LookupEnv("NOTION_API_TOKEN")
	if !ok {
		log.Fatal("NOTION_API_TOKEN not set")
	}

	// Initialise Telegram bot.
	b, err := tb.NewBot(tb.Settings{
		Token:  token,
		Poller: &tb.LongPoller{Timeout: 10 * time.Second},
	})
	if err != nil {
		log.Fatal(err)
	}

	// Register bot handlers.
	b.Handle("/hello", func(m *tb.Message) {
		b.Send(m.Sender, "Hello World!")
	})
	b.Handle("/record", func(m *tb.Message) {
		weight, err := strconv.ParseFloat(m.Payload, 64)
		if err != nil {
			log.Print(err)
			b.Send(m.Sender, "Oops invalid payload. Please enter a number!")
			return
		}
		date := m.Time().Format("2006-01-02")
		go saveToCSV(date, weight, path)
		go saveToNotion(date, weight, dbID, bearerToken)
		b.Send(m.Sender, "Saved mesaurement! Have a great day ahead.")
	})

	// Start listening for updates.
	b.Start()
}
