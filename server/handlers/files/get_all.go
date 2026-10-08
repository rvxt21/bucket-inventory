package files

import (
	"log/slog"
	"net/http"

	"github.com/labstack/echo/v5"
)

// GetFiles godoc
//
//	@Summary	List files with presigned download links
//	@Tags		files
//	@Produce	json
//	@Success	200	{array}		dto.FileResponse
//	@Failure	500	{object}	echo.HTTPError
//	@Router		/files [get]
func (h *Handler) GetFiles(c *echo.Context) error {
	ctx := c.Request().Context()

	files, err := h.service.GetFiles(ctx)
	if err != nil {
		h.log.Error("failed to get files", slog.Any("error", err))
		return echo.NewHTTPError(http.StatusInternalServerError, "failed to get files")
	}

	return c.JSON(http.StatusOK, files)
}
