package database

import (
	"context"
	"errors"

	jet "github.com/go-jet/jet/v2/postgres"
	"github.com/go-jet/jet/v2/qrm"
	"github.com/rvxt21/bucket-inventory/internal/database/gen/files/files/model"
	"github.com/rvxt21/bucket-inventory/internal/database/gen/files/files/table"
)

func (p *Postgres) DeleteFile(ctx context.Context, id string) (string, error) {
	stmt := table.File.
		DELETE().
		WHERE(table.File.ID.EQ(jet.String(id))).
		RETURNING(table.File.StorageKey)

	var res model.File

	err := stmt.QueryContext(ctx, p.db, &res)
	if err != nil {
		if errors.Is(err, qrm.ErrNoRows) {
			return "", ErrNotFound
		}

		return "", ErrDeleteFile.Wrap(err)
	}

	return res.StorageKey, nil
}
