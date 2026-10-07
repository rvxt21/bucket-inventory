package files

import (
	"log/slog"

	"github.com/go-playground/validator/v10"
	filesservice "github.com/rvxt21/bucket-inventory/internal/service/files"
)

type Handler struct {
	log       *slog.Logger
	service   filesservice.Service
	validator *validator.Validate
}

func NewHandler(log *slog.Logger, service filesservice.Service) *Handler {
	return &Handler{
		log:       log,
		service:   service,
		validator: validator.New(),
	}
}
