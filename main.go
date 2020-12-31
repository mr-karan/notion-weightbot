package main

import (
	"encoding/csv"
	"log"
	"os"
	"strconv"
	"time"

	tb "gopkg.in/tucnak/telebot.v2"
)

func main() {
	token, ok := os.LookupEnv("TELEGRAM_BOT_TOKEN")
	if !ok {
		log.Fatal("TELEGRAM_BOT_TOKEN not set")
	}
	path, ok := os.LookupEnv("WEIGHTBOT_CSV_FILE")
	if !ok {
		log.Fatal("WEIGHTBOT_CSV_FILE not set")
	}

	b, err := tb.NewBot(tb.Settings{
		Token:  token,
		Poller: &tb.LongPoller{Timeout: 10 * time.Second},
	})
	if err != nil {
		log.Fatal(err)
		return
	}

	b.Handle("/hello", func(m *tb.Message) {
		b.Send(m.Sender, "Hello World!")
	})

	b.Handle("/record", func(m *tb.Message) {
		_, err := strconv.ParseFloat(m.Payload, 32)
		if err != nil {
			log.Print(err)
			b.Send(m.Sender, "oops invalid payload. Please enter a number!")
			return
		}
		file, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		var data [][]string
		data = append(data, []string{m.Time().Format("2006-01-02"), m.Payload})

		w := csv.NewWriter(file)
		w.WriteAll(data)
		b.Send(m.Sender, "Saved mesaurement! Have a great day ahead.")
	})

	b.Start()
}
