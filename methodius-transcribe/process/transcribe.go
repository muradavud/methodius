package process

import (
	"fmt"
	"methodius-transcribe/config"

	"github.com/aws/aws-sdk-go/aws"
	transcribe "github.com/aws/aws-sdk-go/service/transcribeservice"
)

func Transcribe(objectKey string) error {
	transcribeClient := transcribe.New(AwsSession)

	_, err := transcribeClient.StartTranscriptionJob(&transcribe.StartTranscriptionJobInput{
		LanguageCode: aws.String("en-US"), //TODO: make it configurable
		Media: &transcribe.Media{
			MediaFileUri: aws.String(fmt.Sprintf("s3://%s/%s", config.Conf.S3BucketName, objectKey+".ogg")),
		},
		TranscriptionJobName: aws.String(fmt.Sprintf("transcribe-job-%s", objectKey)),
		OutputBucketName:     aws.String(config.Conf.S3BucketName),
		OutputKey:            aws.String(fmt.Sprintf("transcripts/transcript-%s", objectKey+".json")),
	})
	if err != nil {
		return err
	}

	return nil
}
