package repositories

import (
	"github.com/go-transcoder/uploader/infrastructure/db"
	"log"
)

type Video struct {
	ID          string `db:"id"`
	Title       string `db:"title"`
	Path        string `db:"path"`
	Url         string `db:"url"`
	IsTranscode bool   `db:"is_transcode"`
	CreatedAt   string `db:"created_at"`
	UpdatedAt   string `db:"updated_at"`
}

func GetVideos() []Video {
	var db = db.GetInstance()

	rows, err := db.Connection.Query("SELECT * FROM videos")

	if err != nil {
		log.Fatal(err)
	}
	var videos []Video
	for rows.Next() {
		var video Video
		if err := rows.Scan(&video.ID, &video.Title, &video.Path, &video.Url, &video.IsTranscode, &video.CreatedAt, &video.UpdatedAt); err != nil {
			log.Fatal(err)
		}
		videos = append(videos, video)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	return videos
}
