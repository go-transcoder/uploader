package entities

import (
	"errors"
	"github.com/google/uuid"
	"time"
)

type UnprocessedVideo struct {
	ID        uuid.UUID
	Name      string
	Path      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewUnprocessedVideo(path, name string) (*UnprocessedVideo, error) {
	// validation
	if path == "" || name == "" {
		return nil, errors.New("invalid details")
	}

	return &UnprocessedVideo{
		ID:   uuid.New(),
		Name: name,
		Path: path,
	}, nil
}
