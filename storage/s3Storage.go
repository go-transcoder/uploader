package storage

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"io"
	"mime/multipart"
	"os"
)

type S3 struct {
	S3Client   *s3.Client
	BucketName string
}

func (s3Api *S3) UploadMultipart(file multipart.File, path string, fileName string) (string, error) {
	defer file.Close()

	// first videos to tmp
	tmpFile, err := os.CreateTemp("", fileName)

	if err != nil {
		return fileName, err
	}
	defer os.Remove(tmpFile.Name()) // Clean up

	// copy the multipart to the tmp file
	_, err = io.Copy(tmpFile, file)

	if err != nil {
		return "", err
	}

	destination := fmt.Sprintf("%s/%s", path, fileName)

	f, err := os.Open(tmpFile.Name())
	defer f.Close()

	_, err = s3Api.S3Client.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket: aws.String(s3Api.BucketName),
		Key:    aws.String(destination),
		Body:   f,
	})

	if err != nil {
		return fileName, err
	}

	return fileName, nil
}
