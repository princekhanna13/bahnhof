package main

import (
	"bahnhof/routes"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"net/http"
	"os"
)

func main() {
	server()
}
func bot() {
	bot, err := tgbotapi.NewBotAPI("5235943255:AAGKXu9H2ou_xn7YgPXyaIItVTqRe8Ax7EQ")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)
	for update := range updates {
		if update.Message.Text == "/home" { // If we got a message
			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

			msg := tgbotapi.NewMessage(update.Message.Chat.ID, routes.HomeRoute())
			msg.ReplyToMessageID = update.Message.MessageID

			bot.Send(msg)
		}
		if update.Message.Text == "/work" { // If we got a message
			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

			msg := tgbotapi.NewMessage(update.Message.Chat.ID, routes.WorkRoute())
			msg.ReplyToMessageID = update.Message.MessageID

			bot.Send(msg)
		}
	}
}

func server() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "9090"
	}
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(writer, "test")
	})
	http.HandleFunc("/route", routes.RouteFinder)
	http.HandleFunc("/start", func(writer http.ResponseWriter, request *http.Request) {
		go bot()
	})
	fmt.Print("test")
	http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
	return
}

//func parallelize(functions ...func()) {
//	var waitGroup sync.WaitGroup
//	waitGroup.Add(len(functions))
//
//	defer waitGroup.Wait()
//
//	for _, function := range functions {
//		go func(copy func()) {
//			defer waitGroup.Done()
//			copy()
//		}(function)
//	}
//}
