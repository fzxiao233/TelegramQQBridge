package Onebot

import (
	"github.com/fzxiao233/TelegramQQBridge/Global"
	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
	"github.com/tidwall/gjson"
	"net/url"
)

type QQReceiver struct {
	*websocket.Conn
	toTGChan chan *QQMsg
}

func (r *QQReceiver) msgFilter(qqMsg *QQMsg) bool {
	if qqMsg.PostType != "message" || qqMsg.MessageType != "private" {
		return false
	}
	return true
}

func (r *QQReceiver) msgParser(rawEvent []byte) *QQMsg {
	qqMsg := &QQMsg{
		Time:        gjson.GetBytes(rawEvent, "time").Int(),
		PostType:    gjson.GetBytes(rawEvent, "post_type").Str,
		MessageType: gjson.GetBytes(rawEvent, "message_type").Str,
		SubType:     gjson.GetBytes(rawEvent, "sub_type").Str,
		MessageId:   gjson.GetBytes(rawEvent, "message_id").Int(),
		Message:     gjson.GetBytes(rawEvent, "message").Str,
		Sender: struct {
			UserId   int64  `json:"user_id"`
			Nickname string `json:"nickname"`
		}{
			UserId:   gjson.GetBytes(rawEvent, "sender.user_id").Int(),
			Nickname: gjson.GetBytes(rawEvent, "sender.nickname").Str,
		},
	}
	//if err != nil {
	//	logrus.Warn("unmarshal msg:", err)
	//}
	return qqMsg
}

func (r *QQReceiver) MsgReader() {
	defer r.Conn.Close()
	for {
		_, msg, err := r.Conn.ReadMessage()
		logrus.Infof("MsgReader: %s", msg)
		if err != nil {
			logrus.Warn("Read message:", err)
			continue
		}
		qqMsg := r.msgParser(msg)
		if qqMsg == nil {
			continue
		}
		if !r.msgFilter(qqMsg) {
			continue
		}
		r.toTGChan <- qqMsg
	}

}

func NewQQReceiver(toTGChan chan *QQMsg) *QQReceiver {
	u := url.URL{Scheme: "ws", Host: Global.Config.WsServer, Path: "/event"}
	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		logrus.Fatal("dial:", err)
	}
	logrus.Info("Onebot Connected")
	return &QQReceiver{c, toTGChan}
}
