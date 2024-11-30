package handler

import (
	"fmt"
	"methodius/bot"
	"methodius/logger"
	"methodius/process"

	tele "gopkg.in/telebot.v3"
)

func Start(c tele.Context) error {
	err := process.UploadChatToDynamo(fmt.Sprint(c.Chat().ID), c.Message().Sender.Username,
		c.Message().Sender.FirstName, c.Message().Sender.LastName,
		"", bot.KV["default_langauge"], false) //TODO default language
	if err != nil {
		logger.Log.Error(err.Error())
		return err
	}

	c.Send(bot.KV["start"])

	return nil
}
