package postgres

import (
	"github.com/go-transcoder/uploader/internal/domain/entities"
	"github.com/go-transcoder/uploader/internal/domain/repositories"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type GormVideosRepository struct {
	db *gorm.DB
}

func NewGormVideosRepository(db *gorm.DB) repositories.VideosRepository {
	return &GormVideosRepository{db: db}
}

func (repo *GormVideosRepository) Create(video *entities.Video) error {
	dbVideo := ToDBVideo(video)

	if err := repo.db.Create(dbVideo).Error; err != nil {
		return err
	}

	storedVideo, err := repo.FindById(dbVideo.ID)

	if err != nil {
		return err
	}

	*video = *storedVideo

	return nil
}

func (repo *GormVideosRepository) FindById(id uuid.UUID) (*entities.Video, error) {
	var dbVideo Video

	if err := repo.db.Find(&dbVideo, id).Error; err != nil {
		return nil, err
	}

	return FromDBVideo(&dbVideo)
}

func (repo *GormVideosRepository) FindAll() ([]*entities.Video, error) {
	var dbVideos []Video
	var err error

	if err := repo.db.Find(&dbVideos).Error; err != nil {
		return nil, err
	}

	videos := make([]*entities.Video, len(dbVideos))
	for i, dbVideo := range dbVideos {
		videos[i], err = FromDBVideo(&dbVideo)
		if err != nil {
			return nil, err
		}
	}
	return videos, nil
}

func (repo *GormVideosRepository) Update(video *entities.Video) error {
	dbVideo := ToDBVideo(video)
	err := repo.db.Model(&Video{}).Where("id = ?", dbVideo.ID).Updates(dbVideo).Error
	if err != nil {
		return err
	}

	// Read row from DB to never return different data than persisted
	storedVideo, err := repo.FindById(dbVideo.ID)
	if err != nil {
		return err
	}

	// Map back to domain entity
	*video = *storedVideo

	return nil
}

func (repo *GormVideosRepository) Delete(id uuid.UUID) error {
	return repo.db.Delete(&Video{}, id).Error
}
