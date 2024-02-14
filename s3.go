package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/aws/aws-sdk-go/service/s3/s3manager/s3manageriface"
	"os"
)

type S3 struct {
	Uploader   s3manageriface.UploaderAPI
	BucketName string
	FilePath   string
}

func (s3 *S3) uploadToS3(fileName string) error {
	path := s3.FilePath + "/" + fileName
	destination := fmt.Sprintf("uploads/%s", fileName)
	_, err := os.Stat(path)

	if os.IsNotExist(err) {
		return err
	}

	f, err := os.Open(path)

	result, err := s3.Uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(s3.BucketName),
		Key:    aws.String(destination),
		Body:   f,
	})

	if err != nil {
		return err
	}

	fmt.Printf("file uploaded to, %s\n", result.Location)

	return nil
}
