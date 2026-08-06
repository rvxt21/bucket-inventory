package service

import (
	"context"
	"io"

	"github.com/google/uuid"

	"github.com/rvxt21/bucket-inventory/internal/storage"
	"github.com/rvxt21/bucket-inventory/pkg/dto"
)

type Service interface {
	UploadFile(ctx context.Context, upload dto.UploadFile) error
	GetFile(ctx context.Context, req dto.GetFileRequest) (*dto.GetFileResponse, error)
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

func (s *S3Service) GetFile(ctx context.Context, req dto.GetFileRequest) (*dto.GetFileResponse, error) {
	body, contentType, err := s.storage.GetObject(ctx, req.Filename)
	if err != nil {
		return nil, ErrGetFile
	}
	defer body.Close()

	data, err := io.ReadAll(body)
	if err != nil {
		return nil, err
	}

	resp := &dto.GetFileResponse{
		File:        data,
		ContentType: contentType,
	}

	return resp, nil
}

func NewS3Service(storage *storage.S3) *S3Service {
	return &S3Service{
		storage: storage,
	}
}
