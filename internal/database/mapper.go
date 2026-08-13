package database

import (
	"time"

	"github.com/rvxt21/bucket-inventory/internal/database/gen/files/files/model"
	"github.com/rvxt21/bucket-inventory/pkg/dto"
)

func makeDTO(m model.File) dto.File {
	return dto.File{
		ID:          m.ID,
		Name:        m.Name,
		StorageKey:  m.StorageKey,
		ContentType: derefOr(m.ContentType, ""),
		Size:        derefOr(m.Size, 0),
		CreatedAt:   derefOr(m.CreatedAt, time.Time{}),
		UpdatedAt:   derefOr(m.UpdatedAt, time.Time{}),
	}
}

func makeDTOs(ms []model.File) []dto.File {
	files := make([]dto.File, 0, len(ms))

	for _, m := range ms {
		files = append(files, makeDTO(m))
	}

	return files
}

func derefOr[T any](p *T, fallback T) T {
	if p == nil {
		return fallback
	}

	return *p
}
