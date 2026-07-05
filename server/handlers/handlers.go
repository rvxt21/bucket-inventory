package handlers

import (
	"log/slog"
	"net/http"

	"github.com/labstack/echo/v5"
	"github.com/rvxt21/bucket-inventory/internal/service"
	"github.com/rvxt21/bucket-inventory/pkg/dto"
)

type Handler struct {
	log     *slog.Logger
	service service.Service
}

func (h *Handler) UploadFile(c *echo.Context) error {
	ctx := c.Request().Context()

	file, err := c.FormFile("file")
	if err != nil {
		return err
	}

	src, err := file.Open()
	if err != nil {
		return err
	}

	defer src.Close()

	contentType := file.Header.Get("Content-Type")
	filename := file.Filename

	uploadReq := dto.UploadFile{
		Filename:    filename,
		ContentType: contentType,
		File:        src,
	}

	err = h.service.UploadFile(ctx, uploadReq)
	if err != nil {
		return err
	}

	return c.NoContent(http.StatusCreated)
}

func NewHandler(log *slog.Logger, service service.Service) *Handler {
	return &Handler{
		log:     log,
		service: service,
	}
}
