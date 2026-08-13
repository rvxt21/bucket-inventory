package file

import (
	"context"
	"io"
	"time"

	"github.com/rvxt21/bucket-inventory/config"
	"github.com/rvxt21/bucket-inventory/pkg/dto"
)

type Service interface {
	UploadFile(ctx context.Context, req dto.UploadFile) (*dto.File, error)
	GetFiles(ctx context.Context) ([]dto.FileResponse, error)
}

type Database interface {
	Create(ctx context.Context, file *dto.CreateFile) (*dto.File, error)
	List(ctx context.Context) ([]dto.File, error)
}

type S3 interface {
	UploadObject(ctx context.Context, key string, body io.Reader, contentType string) error
	GetObject(ctx context.Context, key string) (io.ReadCloser, *string, error)
	PresignGetObject(ctx context.Context, key string, ttl time.Duration) (string, error)
}

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
