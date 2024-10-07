package commands

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func (c *Commander) DefaultBehavior(inpuMessage *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(inpuMessage.Chat.ID, inpuMessage.Text+" is my answer")

	c.bot.Send(msg)
}