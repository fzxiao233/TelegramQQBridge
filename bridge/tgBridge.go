package bridge

import (
	"fmt"
	"github.com/fzxiao233/TelegramQQBridge/adapter/Onebot"
	"github.com/fzxiao233/TelegramQQBridge/adapter/Telegram"
	"github.com/sirupsen/logrus"
	"regexp"
	"strings"
)

type TGBridge struct {
	*Telegram.TG
}

var imageReg *regexp.Regexp
var cqReg *regexp.Regexp

func init() {
	imageReg, _ = regexp.Compile(`url=(\S*)]`)
	cqReg, _ = regexp.Compile(`\[CQ:image\S*]`)
}

func (*TGBridge) parseMsg(qqMsg *Onebot.QQMsg) string {
	qqMessage := qqMsg.Message
	if strings.Contains(qqMessage, "CQ:image") {
		img := imageReg.FindAllString(qqMessage, -1)
		for i := 0; i < len(img); i++ {
			img[i] = strings.ReplaceAll(img[i], "url=", "")
		}
		qqMessage = cqReg.ReplaceAllString(qqMessage, "")
		qqMessage += strings.Join(img, "\n")
	}
	msg := fmt.Sprintf("%s: %s", qqMsg.Sender.Nickname, qqMessage)
	return msg
}

func (tg *TGBridge) PushMsg(qqMsg *Onebot.QQMsg) error {
	msg := tg.parseMsg(qqMsg)
	err := tg.SendMsg(msg)
	logrus.Infof("Push %s to Telegram:", msg)
	return err
}

func NewTGBridge() *TGBridge {
	return &TGBridge{Telegram.NewTG()}
}
