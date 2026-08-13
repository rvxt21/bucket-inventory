package database

import (
	"context"

	"github.com/rvxt21/bucket-inventory/internal/database/gen/files/files/model"
	"github.com/rvxt21/bucket-inventory/internal/database/gen/files/files/table"
	"github.com/rvxt21/bucket-inventory/pkg/dto"
	"github.com/rvxt21/bucket-inventory/pkg/errors"
)

var ErrCreateFile = &errors.Error{Message: "error creating file"}

func (p *Postgres) Create(ctx context.Context, req *dto.CreateFile) (*dto.File, error) {
	q := table.File.INSERT(
		table.File.Name,
		table.File.StorageKey,
		table.File.ContentType,
		table.File.Size,
	).VALUES(
		req.Name,
		req.StorageKey,
		req.ContentType,
		req.Size,
	).RETURNING(table.File.AllColumns)

	var res model.File

	err := q.QueryContext(ctx, p.db, &res)
	if err != nil {
		return nil, ErrCreateFile.Wrap(err)
	}

	file := makeDTO(res)

	return &file, nil
}
