package commands

import (
	"go-bot/internal/service/product"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Commander struct {
	bot            *tgbotapi.BotAPI
	productService *product.Service
}

func NewCommander(bot *tgbotapi.BotAPI, productService *product.Service) *Commander {
	return &Commander{
		bot:            bot,
		productService: productService,
	}
}

func (c *Commander) RunCommand(message *tgbotapi.Message) {
	switch message.Command() {
	case "help":
		c.Help(message)

	case "list":
		c.List(message)

	case "get":
		c.Get(message)

	default:
		c.DefaultBehavior(message)
	}
}
