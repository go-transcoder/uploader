package services

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/go-transcoder/uploader/internal/application/command"
	"github.com/go-transcoder/uploader/internal/domain/entities"
	"github.com/go-transcoder/uploader/internal/domain/repositories"
	videos_lib "github.com/go-transcoder/videos-lib"
	"os"
)

type VideoService struct {
	unityOfWork repositories.UnityOfWork
}

func NewVideoService(unityOfWork repositories.UnityOfWork) *VideoService {
	return &VideoService{
		unityOfWork: unityOfWork,
	}
}

func (service *VideoService) Upload(uploadCommand *command.UploadVideoCommand) (*command.UploadVideoCommandResult, error) {
	service.unityOfWork.StartTransaction()
	videosRepository := service.unityOfWork.GetVideosRepo()

	defer service.unityOfWork.Rollback()

	cfg, err := config.LoadDefaultConfig(context.TODO())

	if err != nil {
		return nil, err
	}

	// Create the record in database to get the ID
	record, err := entities.NewVideo(uploadCommand.File.Filename)

	if err != nil {
		return nil, err
	}

	err = videosRepository.Create(record)

	if err != nil {
		return nil, err
	}

	filename, err := videos_lib.UploadMultipart(cfg, os.Getenv("INPUT_S3_BUCKET"), *uploadCommand.File, "uploads", fmt.Sprintf("%s.mp4", record.ID))

	if err != nil {
		return nil, err
	}

	record.Path = fmt.Sprintf("uploads/%s", filename)

	err = videosRepository.Update(record)
	if err != nil {
		return nil, err
	}

	// commit before exit
	service.unityOfWork.Commit()

	return &command.UploadVideoCommandResult{
		ID:    record.ID,
		Title: record.Title,
	}, nil
}
