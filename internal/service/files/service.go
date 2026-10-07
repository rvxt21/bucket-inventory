package files

import (
	"log/slog"
	"time"

	"github.com/rvxt21/bucket-inventory/config"
)

type FileService struct {
	log     *slog.Logger
	db      Database
	storage S3
	linkTTL time.Duration
}

func NewFileService(cfg *config.Config, log *slog.Logger, storage S3, db Database) *FileService {
	return &FileService{
		log:     log,
		storage: storage,
		db:      db,
		linkTTL: cfg.LinkTTL,
	}
}
