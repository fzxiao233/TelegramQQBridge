package Telegram

import (
	"github.com/fzxiao233/TelegramQQBridge/Global"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/sirupsen/logrus"
	"log"
	"net/http"
	"net/url"
)

type TG struct {
	*tgbotapi.BotAPI
	updatesChan tgbotapi.UpdatesChannel
}

func NewTG() *TG {
	tgClient := &http.Client{}
	if Global.Config.UseProxy {
		proxyUrl, err := url.Parse("http://" + Global.Config.Proxy)
		if err != nil {
			logrus.Panic("Proxy setting error")
		}
		tgClient = &http.Client{Transport: &http.Transport{Proxy: http.ProxyURL(proxyUrl)}}
	}
	bot, err := tgbotapi.NewBotAPIWithClient(Global.Config.TGBotToken, tgClient)
	if err != nil {
		log.Panic(err)
	}
	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates, err := bot.GetUpdatesChan(u)
	//for update := range updates {
	//	if update.Message == nil { // ignore any non-Message Updates
	//		continue
	//	}
	//	log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
	//	msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
	//	msg.ReplyToMessageID = update.Message.MessageID
	//	bot.Send(msg)
	//}
	return &TG{bot, updates}
}
