package process

import (
	"errors"
	"methodius-gpt/config"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

type Item struct {
	VoiceId          string `json:"voice_id"`
	ChatId           string `json:"chat_id"`
	LoadingMessageId string `json:"loading_message_id"`
	Language         string `json:"language"`
}

func GetItemFromDynamo(voiceId string) (*Item, error) {
	dbClient := dynamodb.New(AwsSession)

	input := &dynamodb.GetItemInput{
		TableName: aws.String(config.Conf.TableName),
		Key: map[string]*dynamodb.AttributeValue{
			"voice_id": {
				S: aws.String(voiceId),
			},
		},
	}

	result, err := dbClient.GetItem(input)
	if err != nil {
		return nil, err
	}

	if result.Item == nil {
		return nil, errors.New("item not found")
	}

	var item Item
	err = dynamodbattribute.UnmarshalMap(result.Item, &item)
	if err != nil {
		return nil, err
	}

	return &item, nil
}
