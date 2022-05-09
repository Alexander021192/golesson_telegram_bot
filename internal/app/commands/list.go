package commands

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (c *Commander) List(inputMsg *tgbotapi.Message) {
	products := c.productService.List()
	outputList := "Here all products:\n\n"
	for _, el := range products {
		outputList += el.Title + "\n"
	}

	msg := tgbotapi.NewMessage(inputMsg.Chat.ID, outputList)
	c.bot.Send(msg)
}