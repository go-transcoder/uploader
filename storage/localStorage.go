package storage

import (
	"google.golang.org/protobuf/types/known/timestamppb"
	"io"
	"mime/multipart"
	"os"
	"strconv"
)

type localStorage struct {
	Path string
}

func (storage localStorage) UploadMultipart(file multipart.File, name string) (fileName string, err error) {
	defer file.Close()

	// update the filename
	fileName = strconv.FormatInt(timestamppb.Now().Seconds, 10) + "_" + name
	dst, err := os.Create(storage.Path + "/" + fileName)

	if err != nil {
		return "", err
	}
	defer dst.Close()

	_, err = io.Copy(dst, file)
	if err != nil {
		return "", err
	}

	return fileName, nil
}
