package commands

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func (c *Commander) List(inpuMessage *tgbotapi.Message) {
	var productList = "Here is the list of prodcts\n\n"

	for _, v := range c.productService.List() {
		productList += v.Title + "\n"
	}

	msg := tgbotapi.NewMessage(inpuMessage.Chat.ID, productList)

	c.bot.Send(msg)
}