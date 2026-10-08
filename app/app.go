package app

import (
	"github.com/rvxt21/bucket-inventory/config"
	"github.com/rvxt21/bucket-inventory/internal/database"
	"github.com/rvxt21/bucket-inventory/internal/service/files"
	"github.com/rvxt21/bucket-inventory/internal/storage"
	"github.com/rvxt21/bucket-inventory/pkg/logger"
	"github.com/rvxt21/bucket-inventory/server"
	filehandlers "github.com/rvxt21/bucket-inventory/server/handlers/files"
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
			fx.Annotate(storage.NewS3, fx.As(new(files.S3)), fx.As(fx.Self())),
			fx.Annotate(database.NewPostgres, fx.As(new(files.Database)), fx.As(fx.Self())),
			fx.Annotate(files.NewFileService, fx.As(new(files.Service))),
			filehandlers.NewHandler,
		),

		fx.Invoke(invokeHooks),
	)
}

type hooks struct {
	fx.In

	Server   *server.Server
	S3       *storage.S3
	Postgres *database.Postgres
}

func invokeHooks(lc fx.Lifecycle, h hooks) {
	lc.Append(fx.Hook{OnStart: h.Postgres.Start, OnStop: h.Postgres.Stop})
	lc.Append(fx.Hook{OnStart: h.S3.Start, OnStop: h.S3.Stop})
	lc.Append(fx.Hook{OnStart: h.Server.Start, OnStop: h.Server.Stop})
}
