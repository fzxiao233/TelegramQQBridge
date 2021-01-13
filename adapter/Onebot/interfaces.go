package Onebot

type QQMsg struct {
	Time        int64  `json:"time"`
	PostType    string `json:"post_type"`
	MessageType string `json:"message_type"`
	SubType     string `json:"sub_type"`
	MessageId   int64  `json:"message_id"`
	Message     string `json:"message"`
	Sender      struct {
		UserId   int64  `json:"user_id"`
		Nickname string `json:"nickname"`
	} `json:"sender"`
}
