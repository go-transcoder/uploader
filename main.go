package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"net/http"
	"os"
)

func upload(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "GET":
		http.ServeFile(w, r, "./upload.html")
	case "POST":
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
		}

		fileName, err := uploadFileToLocalStorage(r)

		if err != nil {
			http.Error(w, fmt.Sprintf("Error uploading: %v", err), http.StatusBadRequest)
			return
		}

		fmt.Fprintf(w, "File Uploaded successfully.")

		go uploadToS3(fileName)

		fmt.Fprintf(w, "File successfully pushed to queue.")
	}
}

func main() {
	fmt.Println("Hello World")

	http.HandleFunc("/", upload)

	err := http.ListenAndServe(":4000", nil)
	if err != nil {
		fmt.Printf("error starting server: %v", err)
	}
}

func uploadToS3(fileName string) {
	sess := session.Must(session.NewSession())
	s3 := S3{
		Uploader:   s3manager.NewUploader(sess),
		BucketName: os.Getenv("INPUT_S3_BUCKET"),
		FilePath:   os.Getenv("UPLOADER_APP_UPLOAD_PATH"),
	}

	err := s3.uploadToS3(fileName)
	if err != nil {
		fmt.Printf("error uploading to s3: %v", err)
	}
}
