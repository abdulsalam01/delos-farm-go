package pond

import (
	"context"

	"github.com/abdulsalam/delos/internal/entity"
)

func (q *Queries) UpdatePondByID(ctx context.Context, args entity.Pond) (int64, error) {
	var (
		rowsAffected int64
		err          error
	)

	data := q.toBase(args)
	row, err := q.db.ExecContext(ctx, queryUpdateFromBaseRepo,
		data.FarmID,
		data.Name,
		data.Size,
		data.WaterSource,
		data.CreatedBy,
		data.UpdatedBy,
		data.CreatedAt,
		data.UpdatedAt,
		data.ID,
	)

	if err != nil {
		return rowsAffected, err
	}

	return row.RowsAffected()
}
