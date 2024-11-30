package process

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
)

func DownloadFromS3(bucket string, key string) (*s3.GetObjectOutput, error) {
	s3Client := s3.New(AwsSession)

	params := &s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	}

	resp, err := s3Client.GetObject(params)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
