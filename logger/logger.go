package logger

import (
	"log/slog"
	"os"

	"github.com/rvxt21/bucket-inventory/config"
)

func NewLogger(cfg *config.Config) (*slog.Logger, error) {
	var logLvl slog.Level

	err := logLvl.UnmarshalText([]byte(cfg.LogLevel))
	if err != nil {
		return nil, ErrParsingLogLevel
	}

	handler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: logLvl,
	})

	log := slog.New(handler)

	return log, nil
}
