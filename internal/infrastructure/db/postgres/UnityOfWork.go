package postgres

import (
	"github.com/go-transcoder/uploader/internal/domain/repositories"
	"gorm.io/gorm"
)

type UnityOfWork struct {
	db *gorm.DB
	tx *gorm.DB
}

func NewUnityOfWork(db *gorm.DB) repositories.UnityOfWork {
	return &UnityOfWork{
		db: db,
	}
}

func (uof *UnityOfWork) StartTransaction() {
	uof.tx = uof.db.Begin()
}

func (uof *UnityOfWork) GetVideosRepo() repositories.VideosRepository {
	return NewGormVideosRepository(uof.db)
}

func (uof *UnityOfWork) Rollback() {
	uof.tx.Rollback()
}

func (uof *UnityOfWork) Commit() {
	uof.tx.Commit()
}
