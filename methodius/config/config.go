package config

import (
	_ "embed"

	"gopkg.in/yaml.v2"
)

var Conf *Config

type (
	Config struct {
		Debug                 bool   `yaml:"debug"`
		TgToken               string `yaml:"tg_token"`
		UserPassword          string `yaml:"user_password"`
		SqsQueueName          string `yaml:"sqs_queue_name"`
		VoiceQueriesTableName string `yaml:"voice_queries_table_name"`
		ChatsTableName        string `yaml:"chats_table_name"`
		AwsRegion             string `yaml:"aws_region"`
		AwsAccessKeyId        string `yaml:"aws_access_key_id"`
		AwsAccessKey          string `yaml:"aws_secret_access_key"`
		Logger                Logger `yaml:"logger"`
	}

	Logger struct {
		Level string `yaml:"log_level"`
	}
)

//go:embed config.yml
var yml []byte

func NewConfig() error {
	Conf = &Config{}

	err := yaml.Unmarshal(yml, Conf)
	if err != nil {
		return err
	}

	return nil
}
