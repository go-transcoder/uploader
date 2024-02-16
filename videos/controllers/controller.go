package controllers

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/go-transcoder/uploader/storage"
	"github.com/go-transcoder/uploader/videos/services"
	"net/http"
	"os"
)

func Index(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./videos/templates/upload.html")
}

func Post(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
	}

	file, fileHeader, err := r.FormFile("file")

	if err != nil {
		fmt.Printf("error getting file: %v", err)
	}

	cfg, err := config.LoadDefaultConfig(context.TODO())

	s3Storage := storage.S3{
		S3Client:   s3.NewFromConfig(cfg),
		BucketName: os.Getenv("INPUT_S3_BUCKET"),
	}
	go func() {
		fileName, err := services.Upload(&s3Storage, file, "uploads", fileHeader.Filename)

		if err != nil {
			fmt.Printf("error uploading to s3: %v", err)
		}

		fmt.Printf("File %v has been successfully uploaded to bucket\n", fileName)
	}()

	fmt.Fprintf(w, "File successfully pushed to bucket.")

}
