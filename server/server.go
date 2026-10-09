package server

import (
	"context"
	"errors"
	"log/slog"
	"net"
	"net/http"

	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
	"github.com/rvxt21/bucket-inventory/config"
	_ "github.com/rvxt21/bucket-inventory/open_api" // registers the generated spec
	filehandlers "github.com/rvxt21/bucket-inventory/server/handlers/files"
)

const bytesInMB = 1 << 20

type Server struct {
	server   *http.Server
	config   *config.Config
	logger   *slog.Logger
	handlers *filehandlers.Handler
}

func (s *Server) Start(_ context.Context) error {
	router := echo.New()

	router.Use(middleware.Recover())

	addr := net.JoinHostPort(s.config.HTTP.Host, s.config.HTTP.Port)
	s.server = &http.Server{Addr: addr, Handler: router, ReadTimeout: s.config.ReadTimeout}

	router.POST("/files", s.handlers.UploadFile, middleware.BodyLimit(s.config.MaxUploadMB*bytesInMB))
	router.GET("/files", s.handlers.GetFiles)
	router.GET("/files/:id", s.handlers.GetFileByID)
	router.DELETE("/files/:id", s.handlers.DeleteFile)

	router.GET("/readyz", s.handlers.Ready)
	router.GET("/healthz", s.handlers.Live)

	go func() {
		err := s.server.ListenAndServe()
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			s.logger.Error("failed to start server", slog.Any("error", err.Error()))
		}
	}()

	return nil
}

func (s *Server) Stop(ctx context.Context) error {
	return s.server.Shutdown(ctx)
}

func NewServer(cfg *config.Config, log *slog.Logger, h *filehandlers.Handler) *Server {
	return &Server{
		config:   cfg,
		logger:   log,
		handlers: h,
	}
}
