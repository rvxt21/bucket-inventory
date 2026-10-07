package files

import (
	"time"

	"github.com/rvxt21/bucket-inventory/config"
)

type FileService struct {
	db      Database
	storage S3
	linkTTL time.Duration
}

func NewFileService(cfg *config.Config, storage S3, db Database) *FileService {
	return &FileService{
		storage: storage,
		db:      db,
		linkTTL: cfg.LinkTTL,
	}
}
