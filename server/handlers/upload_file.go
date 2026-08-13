package handlers

import (
	"errors"
	"io"
	"log/slog"
	"net/http"

	"github.com/labstack/echo/v5"
	"github.com/rvxt21/bucket-inventory/pkg/dto"
)

const sniffLen = 512

func (h *Handler) UploadFile(c *echo.Context) error {
	ctx := c.Request().Context()

	file, err := c.FormFile("file")
	if err != nil {
		h.log.Error("failed to read form file", slog.Any("error", err))
		return echo.NewHTTPError(http.StatusBadRequest, "file is required")
	}

	src, err := file.Open()
	if err != nil {
		h.log.Error("failed to open form file", slog.Any("error", err))
		return err
	}

	defer src.Close()

	filename := file.Filename

	contentType, err := detectContentType(src)
	if err != nil {
		h.log.Error("failed to detect content type", slog.String("filename", filename), slog.Any("error", err))
		return err
	}

	req := dto.UploadFile{
		Filename:    filename,
		ContentType: contentType,
		Size:        file.Size,
		File:        src,
	}

	if err := h.validator.StructCtx(ctx, req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "validation error")
	}

	_, err = h.service.UploadFile(ctx, req)
	if err != nil {
		h.log.Error("failed to upload file", slog.String("filename", filename), slog.Any("error", err))
		return err
	}

	return c.NoContent(http.StatusCreated)
}

func detectContentType(src io.ReadSeeker) (string, error) {
	buf := make([]byte, sniffLen)

	n, err := src.Read(buf)
	if err != nil && !errors.Is(err, io.EOF) {
		return "", err
	}

	_, err = src.Seek(0, io.SeekStart)
	if err != nil {
		return "", err
	}

	return http.DetectContentType(buf[:n]), nil
}
