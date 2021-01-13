package main

import "github.com/fzxiao233/TelegramQQBridge/bridge"

func main() {
	b := bridge.NewBridge()
	b.Run()
}
