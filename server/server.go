package server

import (
	"context"
	"errors"
	"log/slog"
	"net"
	"net/http"

	"github.com/labstack/echo/v5"
	"github.com/rvxt21/bucket-inventory/config"
)

type Server struct {
	server *http.Server
	config *config.Config
	logger *slog.Logger
}

func (s *Server) Start(_ context.Context) error {
	router := echo.New()

	addr := net.JoinHostPort(s.config.Host, s.config.Port)
	s.server = &http.Server{Addr: addr, Handler: router, ReadTimeout: s.config.ReadTimeout}

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

func NewServer(cfg *config.Config, log *slog.Logger) *Server {
	return &Server{
		config: cfg,
		logger: log,
	}
}
