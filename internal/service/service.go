package service

import (
	"context"

	"github.com/google/uuid"

	"github.com/rvxt21/bucket-inventory/internal/storage"
	"github.com/rvxt21/bucket-inventory/pkg/dto"
)

type Service interface {
	UploadFile(ctx context.Context, upload dto.UploadFile) error
}

type S3Service struct {
	storage *storage.S3
}

func (s *S3Service) UploadFile(ctx context.Context, upload dto.UploadFile) error {
	key := uuid.New().String() + "." + upload.Filename
	err := s.storage.UploadObject(ctx, key, upload.File, upload.ContentType)
	if err != nil {
		return ErrUploadFile
	}

	return nil
}

func NewS3Service(storage *storage.S3) *S3Service {
	return &S3Service{
		storage: storage,
	}
}
