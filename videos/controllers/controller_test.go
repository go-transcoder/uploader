package controllers

import (
	"bytes"
	"fmt"
	"github.com/go-transcoder/uploader/videos/services"
	"io"
	"mime/multipart"
	"net/http/httptest"
	"os"
	"testing"
)

func TestIndex(t *testing.T) {
	// make sure that the template exists
	templatePath := "../templates/upload.html"
	// the output file should exist
	_, err := os.Stat(templatePath)

	if os.IsNotExist(err) {
		t.Fatalf("File %s is expected to be in path. error: %v", templatePath, err)
	}
}

func TestPost(t *testing.T) {

	// mocking the request
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

	// Create a new HTTP request with the form data
	req := httptest.NewRequest("POST", "/videos", &buf)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	// Create a ResponseRecorder to record the response
	//rr := httptest.NewRecorder()

	// Call the videos handler function directly and pass in the request and response recorder
	// instead  of Post(rr, req) I want to mock it and not using the exact method
	if err := req.ParseForm(); err != nil {
		t.Fatalf("ParseForm() err: %v", err)
	}

	requestFile, fileHeader, err := req.FormFile("file")

	if err != nil {
		t.Fatalf("Error getting file from form err: %v", requestFile)
	}

	defer requestFile.Close()

	if fmt.Sprintf("/tmp/%s", fileHeader.Filename) != tmpFile.Name() {
		t.Fatalf("Error getting file name for %s expected %s", fileHeader.Filename, tmpFile.Name())
	}

	mockedStorage := mockStorage{}
	_, err = services.Upload(mockedStorage, requestFile, "uploads", fileHeader.Filename)

	if err != nil {
		t.Fatalf("Error uploading File err: %v", err)
	}
}

type mockStorage struct {
}

func (mockStorage mockStorage) UploadMultipart(file multipart.File, path string, fileName string) (string, error) {
	// this a mock return the filename and no error
	return fileName, nil
}
