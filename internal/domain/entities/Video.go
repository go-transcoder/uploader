package entities

import (
	"errors"
	"github.com/google/uuid"
	"time"
)

type Video struct {
	ID          uuid.UUID
	Title       string
	Url         string
	Path        string
	IsProcessed bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func NewVideo(title string) (*Video, error) {

	if title == "" {
		return nil, errors.New("invalid videos details")
	}

	return &Video{
		ID:          uuid.New(),
		Title:       title,
		Url:         "",
		Path:        "",
		IsProcessed: false,
	}, nil
}
