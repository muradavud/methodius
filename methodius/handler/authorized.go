package handler

import (
	"methodius/bot"

	tele "gopkg.in/telebot.v3"
)

func OnAuthorized(c tele.Context) error {

	c.Send(bot.KV["authorized"])

	bot.SendLanguageInline(c)

	return nil
}
