package bridge

import (
	"github.com/fzxiao233/TelegramQQBridge/adapter/Onebot"
	"github.com/sirupsen/logrus"
)

type Bridge struct {
	toTGChan chan *Onebot.QQMsg
	*TGBridge
	*Onebot.QQ
	//toQQChan chan int
}

func (b *Bridge) Run() {
	go b.QQ.MsgReader()
	for QQMsg := range b.toTGChan {
		err := b.PushMsg(QQMsg)
		if err != nil {
			logrus.Warn("Push msg:", err)
		}
	}
}

func NewBridge() *Bridge {
	tgBridge := NewTGBridge()
	toTGChan := make(chan *Onebot.QQMsg)
	return &Bridge{
		toTGChan: toTGChan,
		TGBridge: tgBridge,
		QQ:       Onebot.NewQQ(toTGChan),
	}
}
