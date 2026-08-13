package database

import (
	"context"
	"errors"

	"github.com/go-jet/jet/v2/qrm"
	"github.com/rvxt21/bucket-inventory/internal/database/gen/files/files/model"
	"github.com/rvxt21/bucket-inventory/internal/database/gen/files/files/table"
	"github.com/rvxt21/bucket-inventory/pkg/dto"
	pkgerrors "github.com/rvxt21/bucket-inventory/pkg/errors"
)

var ErrListFiles = &pkgerrors.Error{Message: "error listing files"}

func (p *Postgres) List(ctx context.Context) ([]dto.File, error) {
	stmt := table.File.
		SELECT(
			table.File.ID,
			table.File.Name,
			table.File.StorageKey,
			table.File.ContentType,
			table.File.Size,
			table.File.CreatedAt,
			table.File.UpdatedAt,
		).
		ORDER_BY(table.File.CreatedAt.DESC())

	var res []model.File

	err := stmt.QueryContext(ctx, p.db, &res)
	if err != nil && !errors.Is(err, qrm.ErrNoRows) {
		return nil, ErrListFiles.Wrap(err)
	}

	return makeDTOs(res), nil
}
