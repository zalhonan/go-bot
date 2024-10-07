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

	for update := range updates {
		switch update.Message.Command() {
		case "help":
			commandHelp(bot, update.Message)

		case "list":
			commandList(bot, update.Message, productService)

		default:
			defaultBehavior(bot, update.Message)
		}
	}
}

func commandList(bot *tgbotapi.BotAPI, inpuMessage *tgbotapi.Message, productService *product.Service) {
	var productList = "Here is the list of prodcts\n\n"

	for _, v := range productService.List() {
		productList += v.Title + "\n"
	}

	msg := tgbotapi.NewMessage(inpuMessage.Chat.ID, productList)

	bot.Send(msg)
}

func commandHelp(bot *tgbotapi.BotAPI, inpuMessage *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(inpuMessage.Chat.ID, "thats a very helpful bot")

	bot.Send(msg)
}

func defaultBehavior(bot *tgbotapi.BotAPI, inpuMessage *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(inpuMessage.Chat.ID, inpuMessage.Text+" is my answer")

	bot.Send(msg)
}
