package main

import (
	"context"
	"errors"
	"methodius-gpt/config"
	"methodius-gpt/logger"
	"methodius-gpt/process"
	"strconv"
	"strings"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	tele "gopkg.in/telebot.v3"
)

func main() {
	lambda.Start(handler)
}

func handler(ctx context.Context, event events.S3Event) error {
	err := config.NewConfig()
	if err != nil {
		println(err.Error())
		return err
	}

	logger.Log = logger.New(config.Conf.Logger.Level)

	if len(event.Records) == 0 {
		logger.Log.Error("no event received")
		return errors.New("no event received")
	}

	pref := tele.Settings{
		Token:  config.Conf.TgToken,
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}

	b, err := tele.NewBot(pref)
	if err != nil {
		logger.Log.Error(err.Error())
	}

	err = process.NewAwsSession()
	if err != nil {
		logger.Log.Error(err.Error())
		return err
	}

	voiceId := event.Records[0].S3.Object.Key
	voiceId = strings.Replace(voiceId, "transcripts/transcript-", "", 1)
	voiceId = strings.Replace(voiceId, ".json", "", 1)

	object, err := process.DownloadFromS3(config.Conf.S3BucketName, event.Records[0].S3.Object.Key)
	if err != nil {
		logger.Log.Error(err.Error())
		return err
	}
	logger.Log.Info("transcript downloaded from s3")

	err = process.DeleteFromBucket(config.Conf.S3BucketName, event.Records[0].S3.Object.Key)
	if err != nil {
		logger.Log.Error(err.Error()) //TODO remove return, log warn
		return err
	}
	logger.Log.Info("transcript deleted from s3")

	err = process.DeleteFromBucket(config.Conf.S3BucketName, voiceId+".ogg")
	if err != nil {
		logger.Log.Error(err.Error())
		return err
	}
	logger.Log.Info("audio file deleted from s3")

	query, err := process.ExtractStringFromFile(object)
	if err != nil {
		logger.Log.Error(err.Error())
		return err
	}

	item, err := process.GetItemFromDynamo(voiceId)
	if err != nil {
		logger.Log.Error(err.Error())
		return err
	}
	logger.Log.Info("item from dynamo received")

	answer, err := process.GptQuerier(query, item.ChatId)
	if err != nil {
		logger.Log.Error(err.Error())
		return err
	}
	logger.Log.Info("gpt queried successfully")

	chatId, err := strconv.ParseInt(item.ChatId, 10, 64)
	if err != nil {
		logger.Log.Error(err.Error())
		return err
	}
	loadingMsgId, err := strconv.Atoi(item.LoadingMessageId)
	if err != nil {
		logger.Log.Error(err.Error())
		return err
	}

	b.Delete(&tele.Message{ID: loadingMsgId, Chat: &tele.Chat{ID: chatId}})

	_, err = b.Send(&tele.Chat{ID: chatId}, answer)
	if err != nil {
		logger.Log.Error(err.Error())
		return err
	}
	logger.Log.Info("telegram message sent")

	return nil
}
