package process

import (
	"methodius-gpt/config"
	"testing"

	"gopkg.in/go-playground/assert.v1"
)

func TestGetItemFromDynamo(t *testing.T) {
	config.NewConfig()
	NewAwsSession()

	item, err := GetItemFromDynamo("99")
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, "99", item.ChatId)
}
