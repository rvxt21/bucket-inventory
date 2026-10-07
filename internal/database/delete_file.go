package database

import (
	"context"

	jet "github.com/go-jet/jet/v2/postgres"
	"github.com/rvxt21/bucket-inventory/internal/database/gen/files/files/table"
)

func (p *Postgres) DeleteFile(ctx context.Context, id string) (bool, error) {
	stmt := table.File.DELETE().WHERE(table.File.ID.EQ(jet.String(id)))

	res, err := stmt.ExecContext(ctx, p.db)
	if err != nil {
		return false, ErrDeleteFile.Wrap(err)
	}

	affected, err := res.RowsAffected()
	if err != nil {
		return false, ErrCheckAffectedRows.Wrap(err)
	}

	return affected > 0, nil
}
