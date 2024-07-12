package consumer

import (
	"log"

	"github.com/atadzan/bv-manager-bot/messages"
	"github.com/atadzan/bv-manager-bot/processor"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Consumer interface {
	Start()
}

type consumer struct {
	bot       *tgbotapi.BotAPI
	processor processor.Processor
}

func New(bot *tgbotapi.BotAPI, processor processor.Processor) Consumer {
	return &consumer{
		bot:       bot,
		processor: processor,
	}
}

func (c *consumer) Start() {
	updateConfig := tgbotapi.NewUpdate(0)

	updateConfig.Timeout = 30

	updates := c.bot.GetUpdatesChan(updateConfig)

	for update := range updates {
		var (
			msg string
			err error
		)
		if update.Message == nil {
			continue
		} else {
			switch update.Message.Text {
			case messages.CMDHelp:
				msg = messages.Help
			case messages.CMDListProxies:
				msg = c.processor.ListProxies()
			case messages.CMDCheckProxies:
				msg = c.processor.CheckProxies()
			case messages.CMDUpdatePasswords:
				msg = messages.UpdatePasswords
			default:
				msg = messages.UnknownCMD
			}

			tgMsg := tgbotapi.NewMessage(update.Message.Chat.ID, msg)

			tgMsg.ReplyToMessageID = update.Message.MessageID
			if _, err = c.bot.Send(tgMsg); err != nil {
				log.Println(err)
			}
		}

	}
	return
}
