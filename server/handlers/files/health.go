package files

import (
	"context"
	"errors"
	"log/slog"
	"net/http"

	"github.com/labstack/echo/v5"
	healthservice "github.com/rvxt21/bucket-inventory/internal/service/health"
	"github.com/rvxt21/bucket-inventory/pkg/dto"
)

func (h *Handler) Live(c *echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{"status": "ok"})
}

func (h *Handler) Ready(c *echo.Context) error {
	ctx, cancel := context.WithTimeout(c.Request().Context(), h.healthCheckTimeout)
	defer cancel()

	resp := dto.ReadyResponse{Status: "ok"}
	status := http.StatusOK

	err := h.health.Ready(ctx)
	if err != nil {
		resp.Service = "s3"
		resp.Status = "unavailable"
		status = http.StatusServiceUnavailable

		h.log.ErrorContext(ctx, "readiness check failed", slog.Any("error", err))

		if errors.Is(err, healthservice.ErrPingDatabase) {
			resp.Service = "database"
		}
	}

	return c.JSON(status, resp)
}
