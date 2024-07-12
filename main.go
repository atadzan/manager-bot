package main

import (
	"flag"
	"log"

	"github.com/atadzan/bv-manager-bot/config"
	"github.com/atadzan/bv-manager-bot/consumer"
	"github.com/atadzan/bv-manager-bot/processor"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	appCfg := config.MustLoadConfig(mustConfigPath())

	bot, err := tgbotapi.NewBotAPI(appCfg.BotToken)
	if err != nil {
		panic(err)
	}

	bot.Debug = true

	eventProcessor := processor.New(appCfg.Proxies)
	eventConsumer := consumer.New(bot, eventProcessor)

	eventConsumer.Start()
}

func mustConfigPath() string {
	cfgPath := flag.String(
		"config",
		"",
		"config path",
	)

	flag.Parse()

	if *cfgPath == "" {
		log.Fatal("config path is not specified")
	}

	return *cfgPath
}
