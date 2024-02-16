package storage

import "mime/multipart"

type Storage interface {
	UploadMultipart(file multipart.File, path string, name string) (filename string, err error)
}
