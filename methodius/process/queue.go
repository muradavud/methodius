package process

import (
	"methodius/config"
	"methodius/logger"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sqs"
)

func SendMessageToQueue(message string, languageCode string) error {
	sqsClient := sqs.New(AwsSession)

	queueURL, err := sqsClient.GetQueueUrl(&sqs.GetQueueUrlInput{
		QueueName: aws.String(config.Conf.SqsQueueName),
	})
	if err != nil {
		logger.Log.Error(err.Error())
		return err
	}

	attributes := map[string]*sqs.MessageAttributeValue{
		"LanguageCode": {
			DataType:    aws.String("String"),
			StringValue: aws.String(languageCode),
		},
	}

	params := &sqs.SendMessageInput{
		MessageBody:       aws.String(message),
		QueueUrl:          queueURL.QueueUrl,
		MessageAttributes: attributes,
	}

	resp, err := sqsClient.SendMessage(params)
	if err != nil {
		logger.Log.Error(err.Error())
		return err
	}

	logger.Log.Info("Success", aws.StringValue(resp.MessageId))

	return nil
}
