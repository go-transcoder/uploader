package repositories

import (
	"github.com/go-transcoder/uploader/internal/domain/entities"
	"github.com/google/uuid"
)

type VideosRepository interface {
	Create(video *entities.Video) error
	FindById(id uuid.UUID) (*entities.Video, error)
	FindAll() ([]*entities.Video, error)
	Update(product *entities.Video) error
	Delete(id uuid.UUID) error
}
