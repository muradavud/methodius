package main

import (
	"methodius/bot"
	"methodius/config"
	"methodius/handler"
	"methodius/logger"
	"methodius/process"

	"methodius/middleware"

	tele "gopkg.in/telebot.v3"
)

func main() {
	err := config.NewConfig()
	if err != nil {
		panic(err)
	}

	logger.Log = logger.New(config.Conf.Logger.Level)

	err = process.NewAwsSession()
	if err != nil {
		logger.Log.Fatal(err.Error())
	}

	err = bot.NewBot()
	if err != nil {
		logger.Log.Fatal(err.Error())
	}

	bot.B.Handle("/start", handler.Start)
	bot.B.Handle(tele.OnVoice, handler.OnVoice, middleware.Auth)
	bot.B.Handle(tele.OnText, handler.OnText, middleware.Auth)
	bot.B.Handle(tele.OnCallback, handler.OnCallback)

	logger.Log.Info("bot started") //TODO add more logging everywhere
	bot.B.Start()
}
