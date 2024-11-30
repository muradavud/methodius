package process

import (
	"methodius-transcribe/config"
	"time"

	tele "gopkg.in/telebot.v3"
)

var TgBot *tele.Bot

func NewBot() error {
	pref := tele.Settings{
		Token:  config.Conf.TgToken,
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}

	var err error
	TgBot, err = tele.NewBot(pref)
	if err != nil {
		return err
	}

	return nil
}
