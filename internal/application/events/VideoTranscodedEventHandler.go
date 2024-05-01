package events

import (
	"encoding/json"
	"fmt"
	"github.com/go-transcoder/uploader/internal/domain/events"
	"github.com/go-transcoder/uploader/internal/domain/repositories"
	"github.com/google/uuid"
	"strings"
)

type VideoTranscodedEventHandler struct {
	Consumer events.EventConsumerService
	Repo     repositories.VideosRepository
}

type ProcessedVideoEventMessage struct {
	VideoTitle string `json:"video_title"`
}

func NewVideoTranscodedEventHandler(Consumer events.EventConsumerService, Repo repositories.VideosRepository) *VideoTranscodedEventHandler {
	return &VideoTranscodedEventHandler{
		Consumer: Consumer,
		Repo:     Repo,
	}
}

func (h *VideoTranscodedEventHandler) Process() error {

	m, err := h.Consumer.ReadMessage()

	if err != nil {
		return err
	}
	message := *m
	var p ProcessedVideoEventMessage
	err = json.Unmarshal(message.Message, &p)

	if err != nil {
		return err
	}

	// Get the ID from the video title
	parts := strings.Split(p.VideoTitle, ".")
	videoId := parts[0]

	fmt.Printf("videoId %s", videoId)

	videoUUID, err := uuid.Parse(videoId)
	if err != nil {
		return err
	}
	video, err := h.Repo.FindById(videoUUID)
	video.IsProcessed = true

	err = h.Repo.Update(video)
	if err != nil {
		return err
	}

	return nil
}
