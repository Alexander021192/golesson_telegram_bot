package main

import (
	"io/ioutil"
	"log"

	"github.com/Alexander021192/bot/internal/app/commands"
	"github.com/Alexander021192/bot/internal/service/product"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	// for debug
	content, err := ioutil.ReadFile("./.env")
	if err != nil {
		log.Fatal(err)
	}
	token := string(content)
	// fmt.Println(token)

	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.UpdateConfig{
		Timeout: 60,
	}

	updates := bot.GetUpdatesChan(u)

	productService := product.NewService()

	commander := commands.NewCommander(bot, productService)

	page := 1

	for update := range updates {
		commander.HandleUpdate(update, &page)
	}
}
