package postgres

import "github.com/go-transcoder/uploader/internal/domain/entities"

// TODO: we can remove the mapper, but keeping it as an example

func ToDBVideo(video *entities.Video) *Video {
	var v = &Video{
		Title:       video.Title,
		Url:         video.Url,
		Path:        video.Path,
		IsProcessed: video.IsProcessed,
		CreatedAt:   video.CreatedAt,
		UpdatedAt:   video.UpdatedAt,
	}

	v.ID = video.ID

	return v
}

func FromDBVideo(dbProduct *Video) (*entities.Video, error) {
	var v = &entities.Video{
		Title:       dbProduct.Title,
		Url:         dbProduct.Url,
		Path:        dbProduct.Path,
		IsProcessed: dbProduct.IsProcessed,
		CreatedAt:   dbProduct.CreatedAt,
		UpdatedAt:   dbProduct.UpdatedAt,
	}
	v.ID = dbProduct.ID

	return v, nil
}
