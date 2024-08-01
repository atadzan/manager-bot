package main

import (
	"flag"
	"log"

	"github.com/atadzan/bv-manager-bot/consumer"
	"github.com/atadzan/bv-manager-bot/processor"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {

	bot, err := tgbotapi.NewBotAPI(mustToken())
	if err != nil {
		panic(err)
	}

	commands := tgbotapi.NewSetMyCommands([]tgbotapi.BotCommand{
		{
			Command:     "/help",
			Description: "Help info",
		},
		{
			Command:     "/list_proxies",
			Description: "List proxies",
		},
		{
			Command:     "/check_proxies",
			Description: "Check proxies",
		},
		{
			Command:     "/update_proxies",
			Description: "Update proxies",
		},
		{
			Command:     "/clear_list",
			Description: "Clear proxy list",
		},
		{
			Command:     "/update_passwords",
			Description: "Update passwords",
		},
	}...)

	if _, err = bot.Request(commands); err != nil {
		log.Println(err)
	}
	bot.Debug = true

	eventProcessor := processor.New()
	eventConsumer := consumer.New(bot, eventProcessor)

	eventConsumer.Start()
}

func mustToken() string {
	tgBotToken := flag.String(
		"tg-bot-token",
		"",
		"telegram bot token",
	)

	flag.Parse()

	if *tgBotToken == "" {
		log.Fatal("tg-bot-token is not specified")
	}

	return *tgBotToken
}
