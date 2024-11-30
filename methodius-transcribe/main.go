package main

import (
	"methodius-transcribe/config"
	"methodius-transcribe/logger"
	"methodius-transcribe/process"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(handler)
}

func handler(sqsEvent events.SQSEvent) error {
	err := config.NewConfig()
	if err != nil {
		println(err.Error())
		return err
	}

	logger.Log = logger.New(config.Conf.Logger.Level)

	err = process.NewAwsSession()
	if err != nil {
		logger.Log.Fatal(err.Error())
	}

	err = process.NewBot()
	if err != nil {
		logger.Log.Error(err.Error())
		return err
	}

	err = process.ReceiveMessageFromQueue(sqsEvent)
	if err != nil {
		logger.Log.Error(err.Error())
		return err
	}

	logger.Log.Info("process ended successfully")

	return nil
}
