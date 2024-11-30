package process

import (
	"errors"
	"methodius-transcribe/config"
	"methodius-transcribe/logger"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

func ReceiveMessageFromQueue(sqsEvent events.SQSEvent) error {
	var err error
	AwsSession, err = session.NewSession(&aws.Config{
		Region: aws.String(config.Conf.AwsRegion)},
	)
	if err != nil {
		logger.Log.Fatal(err.Error())
	}

	sqsClient := sqs.New(AwsSession)

	getQueueUrlOutput, err := sqsClient.GetQueueUrl(&sqs.GetQueueUrlInput{
		QueueName: aws.String(config.Conf.SqsQueueName),
	})
	if err != nil {
		return err
	}

	if len(sqsEvent.Records) == 0 {
		return errors.New("no message received from queue")
	}

	for _, msg := range sqsEvent.Records {
		logger.Log.Info("message %s received", msg.Body)

		err = ProcessMessage(msg.Body)
		if err != nil {
			return err
		}

		_, err = sqsClient.DeleteMessage(&sqs.DeleteMessageInput{
			QueueUrl:      getQueueUrlOutput.QueueUrl,
			ReceiptHandle: &msg.ReceiptHandle,
		})
		if err != nil {
			return err
		}

		logger.Log.Info("message deleted from queue")
	}

	return nil
}
