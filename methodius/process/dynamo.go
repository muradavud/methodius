package process

import (
	"errors"
	"fmt"
	"methodius/config"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

type VoiceQuery struct {
	VoiceId          string `json:"voice_id"`
	ChatId           string `json:"chat_id"`
	LoadingMessageId string `json:"loading_message_id"`
	Language         string `json:"language"`
}

type Chat struct {
	ChatId       string `json:"chat_id"`
	Username     string `json:"username"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	History      string `json:"history"`
	Language     string `json:"language"`
	IsAuthorized bool   `json:"is_authorized"`
}

func UploadVoiceQueryToDynamo(voiceId string, chatId int64,
	loadingMessageId int, language string) error {

	dbClient := dynamodb.New(AwsSession) //TODO dynamo and sqs clients to global and init in main

	entry := VoiceQuery{
		VoiceId:          voiceId,
		ChatId:           fmt.Sprint(chatId),
		LoadingMessageId: fmt.Sprint(loadingMessageId),
		Language:         language,
	}

	av, err := dynamodbattribute.MarshalMap(entry)
	if err != nil {
		return err
	}

	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String(config.Conf.VoiceQueriesTableName),
	}

	_, err = dbClient.PutItem(input)
	if err != nil {
		return err
	}

	return nil
}

func UploadChatToDynamo(chatId string, username string, firstName string,
	lastName string, history string,
	language string, isAuthorized bool) error {

	dbClient := dynamodb.New(AwsSession)

	entry := Chat{
		ChatId:       chatId,
		Username:     username,
		FirstName:    firstName,
		LastName:     lastName,
		History:      history,
		Language:     language,
		IsAuthorized: isAuthorized,
	}

	av, err := dynamodbattribute.MarshalMap(entry)
	if err != nil {
		return err
	}

	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String(config.Conf.ChatsTableName),
	}

	_, err = dbClient.PutItem(input)
	if err != nil {
		return err
	}

	return nil
}

func GetChatFromDynamo(chatId string) (*Chat, error) {
	dbClient := dynamodb.New(AwsSession)

	input := &dynamodb.GetItemInput{
		TableName: aws.String(config.Conf.ChatsTableName),
		Key: map[string]*dynamodb.AttributeValue{
			"chat_id": {
				S: aws.String(chatId),
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

	var item Chat
	err = dynamodbattribute.UnmarshalMap(result.Item, &item)
	if err != nil {
		return nil, err
	}

	return &item, nil
}
