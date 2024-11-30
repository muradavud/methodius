package handler

import (
	"fmt"
	"methodius/bot"
	"methodius/logger"
	"methodius/process"

	tele "gopkg.in/telebot.v3"
)

func OnCallback(c tele.Context) error {
	res := c.Callback()

	languageCode := res.Data[1:]

	chat, err := process.GetChatFromDynamo(fmt.Sprint(c.Chat().ID))
	if err != nil {
		logger.Log.Error(err.Error())
		return err
	}

	chat.Language = languageCode //TODO is language code valid?

	err = process.UploadChatToDynamo(chat.ChatId, chat.Username, chat.FirstName,
		chat.LastName, chat.History,
		chat.Language, chat.IsAuthorized) //TODO update instead of upload
	if err != nil {
		logger.Log.Error(err.Error())
		return err
	}

	c.Send(bot.KV[languageCode])

	return c.Respond()
}
