package process

import "methodius-transcribe/logger"

func ProcessMessage(msg string) error {
	file, err := TgBot.FileByID(msg)
	if err != nil {
		return err
	}

	readCloser, err := TgBot.File(&file)
	if err != nil {
		return err
	}
	logger.Log.Info("file downloaded from telegram")

	objectKey := msg

	err = UploadToBucket(readCloser, objectKey)
	if err != nil {
		return err
	}
	logger.Log.Info("file uploaded to s3")

	err = Transcribe(objectKey)
	if err != nil {
		return err
	}
	logger.Log.Info("transcription job created")

	//err = DeleteFromBucket(objectKey)
	//if err != nil {
	//	return err
	//}
	//logger.Log.Info("file deleted from s3")

	return nil
}
