package interfaces

import "github.com/go-transcoder/uploader/internal/application/command"

type VideoService interface {
	Upload(uploadCommand *command.UploadVideoCommand) (*command.UploadVideoCommandResult, error)
}
