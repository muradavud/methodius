package handler

import (
	_ "embed"
	"fmt"
	"methodius/bot"
	"methodius/logger"
	"methodius/process"
	"time"

	tele "gopkg.in/telebot.v3"
)

func OnVoice(c tele.Context) error {
	defer timeTrack(time.Now(), "Voice")

	chat, err := process.GetChatFromDynamo(fmt.Sprint(c.Message().Chat.ID))
	if err != nil {
		logger.Log.Error(err.Error())
		return err
	}

	err = process.SendMessageToQueue(c.Message().Voice.File.FileID, chat.Language)
	if err != nil {
		logger.Log.Error(err.Error())
		return err
	}

	msg, err := bot.SendLoading(c)
	if err != nil {
		logger.Log.Error(err.Error())
		return err
	}

	err = process.UploadVoiceQueryToDynamo(c.Message().Voice.FileID,
		c.Message().Chat.ID, msg.ID, "EN") //TODO rollback if fails, delete from queue
	if err != nil {
		logger.Log.Error(err.Error())
		return err
	}

	return nil
}

func UserJoined(c tele.Context) error {
	return c.Send("hey")

}

func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	logger.Log.Pretty("%s took %s", name, elapsed)
}
