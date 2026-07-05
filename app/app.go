package app

import (
	"github.com/rvxt21/bucket-inventory/config"
	"github.com/rvxt21/bucket-inventory/internal/service"
	"github.com/rvxt21/bucket-inventory/internal/storage"
	"github.com/rvxt21/bucket-inventory/logger"
	"github.com/rvxt21/bucket-inventory/server"
	"github.com/rvxt21/bucket-inventory/server/handlers"
	"go.uber.org/fx"
)

func App(cfg *config.Config) *fx.App {
	return fx.New(
		fx.Supply(
			cfg,
		),

		fx.Provide(
			logger.NewLogger,
			server.NewServer,
			storage.NewS3,
			fx.Annotate(service.NewS3Service, fx.As(new(service.Service))),
			handlers.NewHandler,
		),

		fx.Invoke(invokeHooks),
	)
}

type hooks struct {
	fx.In

	Server *server.Server
	S3     *storage.S3
}

func invokeHooks(lc fx.Lifecycle, h hooks) {
	lc.Append(fx.Hook{OnStart: h.Server.Start, OnStop: h.Server.Stop})
	lc.Append(fx.Hook{OnStart: h.S3.Start, OnStop: h.S3.Stop})
}
