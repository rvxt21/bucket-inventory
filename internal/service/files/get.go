package files

import (
	"context"
	"errors"

	"github.com/rvxt21/bucket-inventory/internal/database"
	"github.com/rvxt21/bucket-inventory/pkg/dto"
)

func (s *FileService) GetFileByID(ctx context.Context, req *dto.GetFileRequest) (*dto.FileResponse, error) {
	file, err := s.db.GetFileByID(ctx, req.ID)
	if err != nil {
		if errors.Is(err, database.ErrNotFound) {
			return nil, ErrFileNotFound
		}

		return nil, ErrGetFile.Wrap(err)
	}

	link, err := s.storage.PresignGetObject(ctx, file.StorageKey, s.linkTTL)
	if err != nil {
		return nil, ErrGetFile.Wrap(err)
	}

	return new(makeFileResponse(file, link)), nil
}
