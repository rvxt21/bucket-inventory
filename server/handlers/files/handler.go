package files

import (
	"log/slog"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/rvxt21/bucket-inventory/config"
	filesservice "github.com/rvxt21/bucket-inventory/internal/service/files"
	healthservice "github.com/rvxt21/bucket-inventory/internal/service/health"
)

type Handler struct {
	log                *slog.Logger
	service            filesservice.Service
	health             healthservice.Service
	healthCheckTimeout time.Duration
	validator          *validator.Validate
}

func NewHandler(cfg *config.Config, log *slog.Logger, service filesservice.Service, health healthservice.Service) *Handler {
	return &Handler{
		log:                log,
		service:            service,
		health:             health,
		healthCheckTimeout: cfg.HealthCheckTimeout,
		validator:          validator.New(),
	}
}
