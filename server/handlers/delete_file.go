package handlers

import (
	"errors"
	"log/slog"
	"net/http"

	"github.com/labstack/echo/v5"
	"github.com/rvxt21/bucket-inventory/internal/service/files"
	"github.com/rvxt21/bucket-inventory/pkg/dto"
)

func (h *Handler) DeleteFile(c *echo.Context) error {
	ctx := c.Request().Context()

	var req dto.DeleteFileRequest

	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, ErrBind.Wrap(err).Error())
	}

	if err := h.validator.StructCtx(ctx, &req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, ErrValidate.Wrap(err).Error())
	}

	err := h.service.DeleteFile(ctx, &req)
	if err != nil {
		if errors.Is(err, files.ErrFileNotFound) {
			return echo.NewHTTPError(http.StatusNotFound, "file not found")
		}

		h.log.Error("failed to delete file by id", slog.Any("error", err))

		return echo.NewHTTPError(http.StatusInternalServerError, "failed to delete file by id")
	}

	return c.NoContent(http.StatusNoContent)
}
