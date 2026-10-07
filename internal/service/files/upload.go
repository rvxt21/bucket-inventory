package files

import (
	"context"

	"github.com/google/uuid"
	"github.com/rvxt21/bucket-inventory/pkg/dto"
)

func (s *FileService) UploadFile(ctx context.Context, req dto.UploadFile) (*dto.File, error) {
	key := uuid.New().String() + "." + req.Filename

	err := s.storage.UploadObject(ctx, key, req.File, req.ContentType)
	if err != nil {
		return nil, ErrUploadFile.Wrap(err)
	}

	file, err := s.db.Create(ctx, &dto.CreateFile{
		Name:        req.Filename,
		StorageKey:  key,
		ContentType: req.ContentType,
		Size:        req.Size,
	})
	if err != nil {
		return nil, ErrUploadFile.Wrap(err)
	}

	return file, nil
}
