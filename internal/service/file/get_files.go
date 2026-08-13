package file

import (
	"context"

	"github.com/rvxt21/bucket-inventory/pkg/dto"
)

func (s *FileService) GetFiles(ctx context.Context) ([]dto.FileResponse, error) {
	files, err := s.db.List(ctx)
	if err != nil {
		return nil, ErrGetFiles.Wrap(err)
	}

	res := make([]dto.FileResponse, 0, len(files))

	for _, file := range files {
		link, err := s.storage.PresignGetObject(ctx, file.StorageKey, s.linkTTL)
		if err != nil {
			return nil, ErrGetFiles.Wrap(err)
		}

		res = append(res, makeFileResponse(file, link))
	}

	return res, nil
}
