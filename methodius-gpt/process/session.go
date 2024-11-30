package process

import (
	"methodius-gpt/config"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
)

var AwsSession *session.Session

func NewAwsSession() error {
	var err error
	AwsSession, err = session.NewSession(&aws.Config{
		Region: aws.String(config.Conf.AwsRegion)},
	)
	if err != nil {
		return err
	}

	return nil
}
