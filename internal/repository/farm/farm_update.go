package farm

import (
	"context"

	"github.com/abdulsalam/delos/internal/entity"
)

func (q *Queries) UpdateFarmByID(ctx context.Context, args entity.Farm) (int64, error) {
	var (
		rowsAffected int64
		err          error
	)

	data := q.toBase(args)
	row, err := q.db.ExecContext(ctx, queryUpdateFromBaseRepo,
		data.Name,
		data.Location,
		data.Size,
		data.Established,
		data.Technologies,
		data.Employees,
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
