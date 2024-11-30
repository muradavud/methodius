package bot

import (
	"methodius/config"
	"time"

	tele "gopkg.in/telebot.v3"
)

var B *tele.Bot
var KV map[string]string

func NewBot() error {
	pref := tele.Settings{
		Token:  config.Conf.TgToken,
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}

	var err error

	KV, err = parseKeyValueFile()
	if err != nil {
		return err
	}

	B, err = tele.NewBot(pref)
	if err != nil {
		return err
	}

	return nil
}
