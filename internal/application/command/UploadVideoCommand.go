package command

import (
	"mime/multipart"
)

type UploadVideoCommand struct {
	File *multipart.FileHeader
}
