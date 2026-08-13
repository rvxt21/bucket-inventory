package handlers

import (
	"log/slog"

	"github.com/go-playground/validator/v10"
	"github.com/rvxt21/bucket-inventory/internal/service/file"
)

type Handler struct {
	log       *slog.Logger
	service   file.Service
	validator *validator.Validate
}

func NewHandler(log *slog.Logger, service file.Service) *Handler {
	return &Handler{
		log:       log,
		service:   service,
		validator: validator.New(),
	}
}
