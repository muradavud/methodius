package process

import (
	"context"
	"methodius-gpt/config"

	openai "github.com/sashabaranov/go-openai"
)

func GptQuerier(query string, chatId string) (string, error) {
	client := openai.NewClient(config.Conf.OpenAIToken)

	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model:       openai.GPT3Dot5Turbo,
			MaxTokens:   100,
			Temperature: 0.7,
			User:        chatId, //TODO
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: query,
				},
			},
		},
	)
	if err != nil {
		return "", err
	}

	return resp.Choices[0].Message.Content, nil
}
