package handlers

import (
	"bytes"
	"log/slog"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v5"
	"github.com/rvxt21/bucket-inventory/internal/service"
	"github.com/rvxt21/bucket-inventory/pkg/dto"
)

type Handler struct {
	log       *slog.Logger
	service   service.Service
	validator *validator.Validate
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

	req := dto.UploadFile{
		Filename:    filename,
		ContentType: contentType,
		File:        src,
	}

	if err := h.validator.StructCtx(ctx, req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "validation error")
	}

	err = h.service.UploadFile(ctx, req)
	if err != nil {
		return err
	}

	return c.NoContent(http.StatusCreated)
}

func (h *Handler) GetFile(c *echo.Context) error {
	ctx := c.Request().Context()

	var req dto.GetFileRequest

	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid request body")
	}

	if err := h.validator.StructCtx(ctx, req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "validation error")
	}

	resp, err := h.service.GetFile(ctx, req)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "failed to get file")
	}

	contentType := "application/octet-stream"
	if resp.ContentType != nil {
		contentType = *resp.ContentType
	}

	return c.Stream(http.StatusOK, contentType, bytes.NewReader(resp.File))
}

func NewHandler(log *slog.Logger, service service.Service) *Handler {
	return &Handler{
		log:       log,
		service:   service,
		validator: validator.New(),
	}
}
