package process

import (
	"io"
	"methodius-transcribe/config"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
)

func UploadToBucket(readCloser io.ReadCloser, objectKey string) error {
	file, err := os.CreateTemp("", "temp_*.ogg")
	if err != nil {
		return err
	}
	defer func() {
		readCloser.Close()
		file.Close()
		os.Remove(file.Name())
	}()

	_, err = io.Copy(file, readCloser)
	if err != nil {
		return err
	}

	fileStat, err := file.Stat()
	if err != nil {
		return err
	}

	_, err = file.Seek(0, 0)
	if err != nil {
		return err
	}

	s3Client := s3.New(AwsSession)

	_, err = s3Client.PutObject(&s3.PutObjectInput{
		Body:          file,
		Bucket:        aws.String(config.Conf.S3BucketName),
		Key:           aws.String(objectKey + ".ogg"),
		ContentLength: aws.Int64(fileStat.Size()),
	})
	if err != nil {
		return err
	}

	return nil
}

func DeleteFromBucket(objectKey string) error {
	s3Client := s3.New(AwsSession)

	_, err := s3Client.DeleteObject(&s3.DeleteObjectInput{
		Bucket: aws.String(config.Conf.S3BucketName),
		Key:    aws.String(objectKey + ".ogg"),
	})
	if err != nil {
		return err
	}

	return nil
}
