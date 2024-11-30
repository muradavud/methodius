package process

import (
	"context"
	"methodius-gpt/config"
	"testing"

	openai "github.com/sashabaranov/go-openai"
)

func TestGptQuerier(t *testing.T) {
	config.NewConfig()
	query := "Say Hello!"

	client := openai.NewClient(config.Conf.OpenAIToken)

	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model:       openai.GPT3Dot5Turbo,
			MaxTokens:   100,
			Temperature: 0.7,
			User:        "default", //TODO
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: query,
				},
			},
		},
	)
	if err != nil {
		t.Errorf("Error: %v", err)
	}

	t.Log(resp.Choices[0].Message.Content)
}
