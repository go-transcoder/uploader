package postgres

import (
	"github.com/go-transcoder/uploader/internal/domain/entities"
	"github.com/go-transcoder/uploader/internal/domain/repositories"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type GormUnprocessedVideoRepository struct {
	db *gorm.DB
}

func NewGormUnprocessedVideoRepo(db *gorm.DB) repositories.UnprocessedVideoRepository {
	return &GormUnprocessedVideoRepository{
		db: db,
	}
}

func (repo *GormUnprocessedVideoRepository) Create(u *entities.UnprocessedVideo) error {
	if err := repo.db.Create(u).Error; err != nil {
		return err
	}

	stored, err := repo.FindById(u.ID)

	if err != nil {
		return err
	}

	*u = *stored

	return nil

}

func (repo *GormUnprocessedVideoRepository) FindById(id uuid.UUID) (*entities.UnprocessedVideo, error) {
	var v entities.UnprocessedVideo

	if err := repo.db.Find(&v, id).Error; err != nil {
		return nil, err
	}

	return &v, nil
}
