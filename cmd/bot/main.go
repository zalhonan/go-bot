package main

import (
	"go-bot/internal/service/product"
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	token := os.Getenv("TOKEN")

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

	commander := NewCommander(bot, productService)

	for update := range updates {
		switch update.Message.Command() {
		case "help":
			commander.Help(update.Message)

		case "list":
			commander.List(update.Message)

		default:
			commander.DefaultBehavior(update.Message)
		}
	}
}

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

func (c *Commander) List(inpuMessage *tgbotapi.Message) {
	var productList = "Here is the list of prodcts\n\n"

	for _, v := range c.productService.List() {
		productList += v.Title + "\n"
	}

	msg := tgbotapi.NewMessage(inpuMessage.Chat.ID, productList)

	c.bot.Send(msg)
}

func (c *Commander) Help(inpuMessage *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(inpuMessage.Chat.ID, "thats a very helpful bot")

	c.bot.Send(msg)
}

func (c *Commander) DefaultBehavior(inpuMessage *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(inpuMessage.Chat.ID, inpuMessage.Text+" is my answer")

	c.bot.Send(msg)
}
