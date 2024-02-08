package main

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sqs"
	"google.golang.org/protobuf/types/known/timestamppb"
	"io"
	"net/http"
	"os"
	"strconv"
)

func publishToQueue(svc *sqs.SQS, fileName string) error {
	queueUrl := os.Getenv("AWS_UPLOAD_QUEUE_URL")

	_, err := svc.SendMessage(&sqs.SendMessageInput{
		DelaySeconds: aws.Int64(10),
		MessageAttributes: map[string]*sqs.MessageAttributeValue{
			"FileName": &sqs.MessageAttributeValue{
				DataType:    aws.String("String"),
				StringValue: aws.String(fileName),
			},
			//"Author": &sqs.MessageAttributeValue{
			//	DataType:    aws.String("String"),
			//	StringValue: aws.String("John Grisham"),
			//},
			//"WeeksOn": &sqs.MessageAttributeValue{
			//	DataType:    aws.String("Number"),
			//	StringValue: aws.String("6"),
			//},
		},
		MessageBody: aws.String("Information about Video"),
		QueueUrl:    aws.String(queueUrl),
	})

	if err != nil {
		return err
	}

	return nil
}

func uploadFileToLocalStorage(r *http.Request) (fileName string, err error) {

	file, fileHeader, err := r.FormFile("file")

	if err != nil {
		return "", err
	}

	defer file.Close()

	// create the new file on the server
	uploadPath := os.Getenv("UPLOADER_APP_UPLOAD_PATH")
	fileName = strconv.FormatInt(timestamppb.Now().Seconds, 10) + "_" + fileHeader.Filename
	dst, err := os.Create(uploadPath + "/" + fileName)

	if err != nil {
		return "", err
	}

	defer dst.Close()

	// Copy the file data to the new file
	_, err = io.Copy(dst, file)
	if err != nil {
		return "", err
	}

	return fileName, nil
}
