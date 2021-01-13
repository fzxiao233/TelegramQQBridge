package Telegram

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

func (t *TG) SendMsg(msg string) error {
	tgmsg := tgbotapi.NewMessage(364985867, msg)
	_, err := t.Send(tgmsg)
	return err
}
