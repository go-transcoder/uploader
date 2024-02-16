package services

import (
	"github.com/go-transcoder/uploader/storage"
	"mime/multipart"
)

func Upload(storage storage.Storage, file multipart.File, path string, name string) (string, error) {
	fileName, err := storage.UploadMultipart(file, path, name)

	if err != nil {
		return "", err
	}

	return fileName, nil
}
