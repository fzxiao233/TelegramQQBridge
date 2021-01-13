package Onebot

type QQ struct {
	*QQReceiver
}

func NewQQ(toTGChan chan *QQMsg) *QQ {
	return &QQ{QQReceiver: NewQQReceiver(toTGChan)}
}
