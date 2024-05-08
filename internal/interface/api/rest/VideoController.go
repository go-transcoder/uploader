package rest

import (
	"fmt"
	"github.com/go-transcoder/uploader/internal/application/command"
	"github.com/go-transcoder/uploader/internal/application/interfaces"
	"github.com/labstack/echo/v4"
	"net/http"
)

type VideoController struct {
	service interfaces.VideoService
}

func NewVideoController(e *echo.Echo, service interfaces.VideoService) *VideoController {
	controller := &VideoController{
		service: service,
	}

	e.POST("/upload", controller.Upload)
	e.GET("/", controller.Index)

	return controller
}

func (uc *VideoController) Index(c echo.Context) error {
	return c.JSON(http.StatusOK, "Hello")
}

func (uc *VideoController) Upload(c echo.Context) error {

	file, err := c.FormFile("file")

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid File",
		})
	}

	result, err := uc.service.Upload(&command.UploadVideoCommand{
		File: file,
	})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": fmt.Sprintf("Error Uploading: %v", err),
		})
	}

	return c.JSON(http.StatusCreated, result)
}
