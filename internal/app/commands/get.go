package commands

import (
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (c *Commander) Get(inpuMessage *tgbotapi.Message) {
	args := inpuMessage.CommandArguments()

	i, err := strconv.Atoi(args)

	var msg tgbotapi.MessageConfig

	if err != nil || i > len(c.productService.List()) {
		msg = tgbotapi.NewMessage(inpuMessage.Chat.ID, "no such product honey")
	} else {
		product := c.productService.List()[i]

		msg = tgbotapi.NewMessage(inpuMessage.Chat.ID, "want get some? "+product.Title)
	}

	c.bot.Send(msg)
}
