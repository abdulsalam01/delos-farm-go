package pond

import (
	"context"

	"github.com/abdulsalam/delos/internal/entity"
	"github.com/google/uuid"
)

func (q *Queries) InsertPond(ctx context.Context, args entity.Pond) (uuid.UUID, error) {
	var (
		id    uuid.UUID
		idStr string
		err   error
	)

	data := q.toBase(args)
	data.ID = uuid.NewString()

	row := q.db.QueryRowContext(ctx, queryInsertFromBaseRepo,
		data.ID,
		data.FarmID,
		data.Name,
		data.Slug,
		data.Size,
		data.WaterSource,
		uuid.Nil,
		uuid.Nil,
		data.CreatedAt,
		data.UpdatedAt,
	)
	if err = row.Scan(&idStr); err != nil {
		return id, err
	}

	id, err = uuid.Parse(idStr)
	if err != nil {
		return id, err
	}

	return id, nil
}
