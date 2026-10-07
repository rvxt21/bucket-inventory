package files

import (
	"errors"
	"log/slog"
	"net/http"

	"github.com/labstack/echo/v5"
	filesservice "github.com/rvxt21/bucket-inventory/internal/service/files"
	"github.com/rvxt21/bucket-inventory/pkg/dto"
	"github.com/rvxt21/bucket-inventory/server/handlers"
)

func (h *Handler) GetFileByID(c *echo.Context) error {
	ctx := c.Request().Context()

	var req dto.GetFileRequest

	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, handlers.ErrBind.Wrap(err).Error())
	}

	if err := h.validator.StructCtx(ctx, &req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, handlers.ErrValidate.Wrap(err).Error())
	}

	file, err := h.service.GetFileByID(ctx, &req)
	if err != nil {
		if errors.Is(err, filesservice.ErrFileNotFound) {
			return echo.NewHTTPError(http.StatusNotFound, "file not found")
		}

		h.log.Error("failed to get file by id", slog.Any("error", err))

		return echo.NewHTTPError(http.StatusInternalServerError, "failed to get file by id")
	}

	return c.JSON(http.StatusOK, file)
}
