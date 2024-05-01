package postgres

import (
	"github.com/google/uuid"
	"time"
)

type Video struct {
	ID          uuid.UUID `gorm:"primaryKey"`
	Title       string
	Url         string
	Path        string
	IsProcessed bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
