package app

import (
	"github.com/rvxt21/bucket-inventory/config"
	"github.com/rvxt21/bucket-inventory/logger"
	"github.com/rvxt21/bucket-inventory/server"
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
		),

		fx.Invoke(invokeHooks),
	)
}

type hooks struct {
	fx.In

	Server *server.Server
}

func invokeHooks(lc fx.Lifecycle, h hooks) {
	lc.Append(fx.Hook{OnStart: h.Server.Start, OnStop: h.Server.Stop})
}
