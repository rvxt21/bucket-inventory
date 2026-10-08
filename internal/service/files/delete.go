package files

import (
	"context"
	"errors"
	"log/slog"

	"github.com/rvxt21/bucket-inventory/internal/database"
	"github.com/rvxt21/bucket-inventory/pkg/dto"
)

func (s *FileService) DeleteFile(ctx context.Context, req *dto.DeleteFileRequest) error {
	key, err := s.db.DeleteFile(ctx, req.ID)
	if err != nil {
		if errors.Is(err, database.ErrNotFound) {
			return ErrFileNotFound
		}

		return ErrDeleteFile.Wrap(err)
	}

	err = s.storage.DeleteObject(ctx, key)
	if err != nil {
		s.log.ErrorContext(ctx, "failed to delete file from s3 bucket",
			slog.String("storage_key", key),
			slog.Any("error", err),
		)
	}

	return nil
}
