package files

import (
	"log/slog"
	"net/http"

	"github.com/labstack/echo/v5"
)

func (h *Handler) GetFiles(c *echo.Context) error {
	ctx := c.Request().Context()

	files, err := h.service.GetFiles(ctx)
	if err != nil {
		h.log.Error("failed to get files", slog.Any("error", err))
		return echo.NewHTTPError(http.StatusInternalServerError, "failed to get files")
	}

	return c.JSON(http.StatusOK, files)
}
