package pond

import (
	"context"

	"github.com/abdulsalam/delos/helper"
	"github.com/google/uuid"
)

func (q *Queries) DeletePondByID(ctx context.Context, id uuid.UUID) (int64, error) {
	var (
		rowsAffected int64
		err          error
	)

	row, err := q.db.ExecContext(ctx, queryDeleteFromBaseRepo,
		helper.GetTimeNow(),
		id.String(),
	)
	if err != nil {
		return rowsAffected, err
	}

	return row.RowsAffected()
}
