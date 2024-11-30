package config

import (
	_ "embed"

	"gopkg.in/yaml.v2"
)

var Conf *Config

type (
	Config struct {
		Debug          bool   `yaml:"debug"`
		TgToken        string `yaml:"tg_token"`
		OpenAIToken    string `yaml:"openai_token"`
		SqsUrl         string `yaml:"sqs_url"`
		SqsQueueName   string `yaml:"sqs_queue_name"`
		TableName      string `yaml:"dynamodb_table_name"`
		AwsRegion      string `yaml:"aws_region"`
		AwsAccessKeyId string `yaml:"aws_access_key_id"`
		AwsAccessKey   string `yaml:"aws_secret_access_key"`
		S3BucketName   string `yaml:"s3_bucket_name"`

		Logger Logger `yaml:"logger"`
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
