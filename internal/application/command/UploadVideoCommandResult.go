package command

import "github.com/google/uuid"

type UploadVideoCommandResult struct {
	ID    uuid.UUID
	Title string
}
