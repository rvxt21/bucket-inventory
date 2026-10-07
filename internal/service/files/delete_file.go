package files

import (
	"context"
	"errors"
	"log/slog"

	"github.com/rvxt21/bucket-inventory/internal/database"
	"github.com/rvxt21/bucket-inventory/pkg/dto"
)

func (s *FileService) DeleteFile(ctx context.Context, req *dto.DeleteFileRequest) error {
	file, err := s.db.GetFileByID(ctx, req.ID)
	if err != nil {
		if errors.Is(err, database.ErrNotFound) {
			return ErrFileNotFound
		}

		return ErrGetFile.Wrap(err)
	}

	del, err := s.db.DeleteFile(ctx, req.ID)
	if err != nil {
		return ErrDeleteFile.Wrap(err)
	}

	if !del {
		return ErrFileNotFound
	}

	err = s.storage.DeleteObject(ctx, file.StorageKey)
	if err != nil {
		s.log.ErrorContext(ctx, "failed to delete file from s3 bucket",
			slog.String("file_id", file.ID),
			slog.String("storage_key", file.StorageKey),
			slog.Any("error", err),
		)
	}

	return nil
}
