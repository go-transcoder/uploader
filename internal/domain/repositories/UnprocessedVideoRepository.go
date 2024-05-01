package repositories

import (
	"github.com/go-transcoder/uploader/internal/domain/entities"
	"github.com/google/uuid"
)

type UnprocessedVideoRepository interface {
	Create(video *entities.UnprocessedVideo) error
	FindById(id uuid.UUID) (*entities.UnprocessedVideo, error)
	//FindAll() ([]*entities.UnprocessedVideo, error)
	//Update(product *entities.UnprocessedVideo) error
	//Delete(id uuid.UUID) error
}
