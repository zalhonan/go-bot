package commands

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func (c *Commander) Help(inpuMessage *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(inpuMessage.Chat.ID, "thats a very helpful bot")

	c.bot.Send(msg)
}