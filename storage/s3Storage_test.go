package storage

import (
	"bytes"
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/gofor-little/env"
	"io"
	"mime/multipart"
	"os"
	"testing"
)

func init() {
	// Set environment variables here
	// Load an .env file and set the key-value pairs as environment variables.
	if err := env.Load("../.env.test"); err != nil {
		panic(err)
	}
}

// Testing the UploadMultipart with localstack
func TestS3_UploadMultipart(t *testing.T) {
	cfg, err := config.LoadDefaultConfig(context.TODO())

	if err != nil {
		t.Fatalf("Error aws LoadingDefaultConfig err: %v", err)
	}
	s3Storage := S3{
		S3Client: s3.NewFromConfig(cfg, func(o *s3.Options) {
			o.UsePathStyle = true
		}),
		BucketName: os.Getenv("INPUT_S3_BUCKET"),
	}

	// create the multipart file so we can upload to S3
	content := "This is a test file."
	tmpFile, err := os.CreateTemp("", "example")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpFile.Name()) // Clean up

	if _, err := tmpFile.Write([]byte(content)); err != nil {
		t.Fatal(err)
	}
	if err := tmpFile.Close(); err != nil {
		t.Fatal(err)
	}

	// Create a new multipart form
	var buf bytes.Buffer
	writer := multipart.NewWriter(&buf)

	// Add the file to the form
	part, err := writer.CreateFormFile("file", tmpFile.Name())
	if err != nil {
		t.Fatal(err)
	}

	// Copy the file content to the form
	file, err := os.Open(tmpFile.Name())
	if err != nil {
		t.Fatal(err)
	}
	defer file.Close()

	if _, err := io.Copy(part, file); err != nil {
		t.Fatal(err)
	}

	writer.Close()

	fileName, err := s3Storage.UploadMultipart(file, "uploads", "example")

	if err != nil {
		t.Fatalf("Error uploading the file err: %v", err)
	}

	fmt.Printf("File Uploaded to s3: %s", fileName)
}
