package Global

import (
	"github.com/hjson/hjson-go"
	"github.com/sirupsen/logrus"
	"io/ioutil"
)

type MainConfig struct {
	WsServer   string
	UseProxy   bool
	Proxy      string
	TGBotToken string
}

var DefaultConfig = `
{
	"WsServer": "127.0.0.1:6700"
	"UseProxy": false
	"Proxy": ""
	"TGBotToken": ""
}
`

var Config *MainConfig

func init() {
	if !IsExistFile("./config.hjson") {
		ioutil.WriteFile("config.hjson", []byte(DefaultConfig), 0644)
		logrus.Panic("Generate default config...\n Program exits.")
	}
	b, err := ioutil.ReadFile("./config.hjson")
	if err != nil {
		logrus.Panic("Can't find config")
	}
	err = hjson.Unmarshal(b, Config)
	if err != nil {
		logrus.Panic("Loading config failed")
	}
}
