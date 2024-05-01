package request

import (
	"github.com/go-transcoder/uploader/internal/application/command"
	"mime/multipart"
)

type UploadRequest struct {
	File *multipart.File
}

func (r *UploadRequest) ToUploadVideoCommand() (*command.UploadVideoCommand, error) {
	return &command.UploadVideoCommand{
		File: r.File,
	}, nil
}
