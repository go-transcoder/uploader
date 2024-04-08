package services

import (
	"context"
	"fmt"
	"github.com/go-transcoder/uploader/infrastructure/db"
	"github.com/go-transcoder/uploader/storage"
	"github.com/go-transcoder/uploader/videos/repositories"
	"log"
	"mime/multipart"
)

func Upload(storage storage.Storage, file multipart.File, path string, name string) (string, error) {

	db := db.GetInstance()
	ctx := context.Background()

	// I need the ID of the video after we create it in the database.
	// we want to store the video by ID

	// Get a Tx for making transaction requests.
	tx, err := db.Connection.BeginTx(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	// Defer a rollback in case anything fails.
	defer tx.Rollback()

	var videoId string

	err = tx.QueryRowContext(ctx, `INSERT INTO videos(title, path, url) VALUES($1, $2, $3) RETURNING id`, name, "uploads", "").Scan(&videoId)
	if err != nil {
		log.Fatal(err)
	}

	fileName, err := storage.UploadMultipart(file, path, fmt.Sprintf("%s.mp4", videoId))

	if err != nil {
		log.Fatal(err)
	}

	// Commit the transaction.
	if err = tx.Commit(); err != nil {
		log.Fatal(err)
	}
	return fileName, nil
}

func GetVideos() []repositories.Video {
	return repositories.GetVideos()
}
