package database

import (
	"context"
	"errors"

	jet "github.com/go-jet/jet/v2/postgres"
	"github.com/go-jet/jet/v2/qrm"
	"github.com/rvxt21/bucket-inventory/internal/database/gen/files/files/model"
	"github.com/rvxt21/bucket-inventory/internal/database/gen/files/files/table"
	"github.com/rvxt21/bucket-inventory/pkg/dto"
	pkgerrors "github.com/rvxt21/bucket-inventory/pkg/errors"
)

var ErrGetFileByID = &pkgerrors.Error{Message: "error get file by id"}

func (p *Postgres) GetFileByID(ctx context.Context, id string) (*dto.File, error) {
	stmt := table.File.
		SELECT(
			table.File.ID,
			table.File.Name,
			table.File.StorageKey,
			table.File.ContentType,
			table.File.Size,
			table.File.CreatedAt,
			table.File.UpdatedAt,
		).WHERE(
		table.File.ID.EQ(jet.String(id)),
	)

	var res model.File

	err := stmt.QueryContext(ctx, p.db, &res)
	if err != nil {
		if errors.Is(err, qrm.ErrNoRows) {
			return nil, ErrNotFound
		}

		return nil, ErrGetFileByID.Wrap(err)
	}

	return new(makeDTO(res)), nil
}
