package consumer

import (
	"fmt"
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
		var err error
		if update.Message == nil {
			continue
		}

		var tgMsg = tgbotapi.NewMessage(update.Message.Chat.ID, "")

		if update.Message.IsCommand() {
			switch update.Message.Command() {
			case messages.CMDHelp, messages.CMDStart:
				tgMsg.Text = messages.Help
			case messages.CMDListProxies:
				tgMsg.Text = c.processor.ListProxies()
			case messages.CMDCheckProxies:
				tgMsg.Text = c.processor.CheckProxies()
			case messages.UpdateProxies:
				tgMsg.Text = messages.UpdateProxiesMsg
				tgMsg.ReplyMarkup = tgbotapi.ForceReply{ForceReply: true, Selective: true}
			case messages.CMDUpdatePasswords:
				tgMsg.Text = messages.UpdatePasswords
			case messages.CMDClearProxyList:
				tgMsg.Text = c.processor.ClearProxyList()
			default:
				tgMsg.Text = messages.UnknownCMD
			}

		} else if update.Message.ReplyToMessage != nil && update.Message.ReplyToMessage.Text == messages.UpdateProxiesMsg {
			tgMsg.Text = c.processor.UpdateProxies(update.Message.Text)

		} else {
			tgMsg.Text = fmt.Sprintf("%s is not command", update.Message.Text)
		}

		if _, err = c.bot.Send(tgMsg); err != nil {
			log.Println(err)
		}

	}
	return
}
