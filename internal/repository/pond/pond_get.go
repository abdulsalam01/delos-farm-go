package pond

import (
	"context"

	"github.com/abdulsalam/delos/internal/entity"
	"github.com/google/uuid"
)

func (q *Queries) GetPondAll(ctx context.Context, args entity.PondRequestWithPagination) ([]entity.Pond, error) {
	var (
		pond []entity.Pond
		err  error
	)

	row, err := q.db.QueryContext(ctx, querySelectByIDWithPagination, args.Limit, args.Offset)
	if err != nil {
		return pond, err
	}
	defer row.Close()

	for row.Next() {
		var result pondTable

		if err = row.Scan(
			&result.ID,
			&result.FarmID,
			&result.Name,
			&result.Slug,
			&result.Size,
			&result.WaterSource,
			&result.CreatedBy,
			&result.UpdatedBy,
			&result.CreatedAt,
			&result.UpdatedAt,
		); err != nil {
			return pond, err
		}

		data, err := result.ToEntity()
		if err != nil {
			continue
		}
		if data.ID == uuid.Nil {
			continue
		}

		pond = append(pond, data)
	}

	// Close rows scanner.
	if err := row.Close(); err != nil {
		return nil, err
	}
	if err := row.Err(); err != nil {
		return nil, err
	}

	return pond, nil
}

func (q *Queries) GetPondByID(ctx context.Context, id uuid.UUID) (entity.Pond, error) {
	var (
		pond   entity.Pond
		result pondTable
		err    error
	)

	row := q.db.QueryRowContext(ctx, querySelectByID, id)
	if err = row.Scan(
		&result.ID,
		&result.FarmID,
		&result.Name,
		&result.Slug,
		&result.Size,
		&result.WaterSource,
		&result.CreatedBy,
		&result.UpdatedBy,
		&result.CreatedAt,
		&result.UpdatedAt,
	); err != nil {
		return pond, err
	}

	pond, err = result.ToEntity()
	if err != nil {
		return pond, err
	}
	if pond.ID == uuid.Nil {
		return pond, err
	}

	return pond, nil
}

func (q *Queries) GetPondBySlug(ctx context.Context, slug string) (entity.Pond, error) {
	var (
		pond   entity.Pond
		result pondTable
		err    error
	)

	row := q.db.QueryRowContext(ctx, querySelectBySlug, slug)
	if err = row.Scan(&result.ID,
		&result.FarmID,
		&result.Name,
		&result.Slug,
		&result.Size,
		&result.WaterSource,
		&result.CreatedBy,
		&result.UpdatedBy,
		&result.CreatedAt,
		&result.UpdatedAt,
	); err != nil {
		return pond, err
	}

	pond, err = result.ToEntity()
	if err != nil {
		return pond, err
	}
	if pond.ID == uuid.Nil {
		return pond, err
	}

	return pond, nil
}

func (q *Queries) GetPondByFarmID(ctx context.Context, farmId uuid.UUID) ([]entity.Pond, error) {
	var (
		pond []entity.Pond
		err  error
	)

	row, err := q.db.QueryContext(ctx, querySelectByFarmID, farmId)
	if err != nil {
		return pond, err
	}
	defer row.Close()

	for row.Next() {
		var result pondTable

		if err = row.Scan(
			&result.ID,
			&result.FarmID,
			&result.Name,
			&result.Slug,
			&result.Size,
			&result.WaterSource,
			&result.CreatedBy,
			&result.UpdatedBy,
			&result.CreatedAt,
			&result.UpdatedAt,
		); err != nil {
			return pond, err
		}

		data, err := result.ToEntity()
		if err != nil {
			continue
		}
		if data.ID == uuid.Nil {
			continue
		}

		pond = append(pond, data)
	}

	// Close rows scanner.
	if err := row.Close(); err != nil {
		return nil, err
	}
	if err := row.Err(); err != nil {
		return nil, err
	}

	return pond, nil
}

func (q *Queries) GetPondCount(ctx context.Context) (int64, error) {
	var (
		pondCount int64
		pondId    string
		err       error
	)

	row := q.db.QueryRowContext(ctx, queryPartialSelectCountActiveOnly)
	if err = row.Scan(&pondId, &pondCount); err != nil {
		return pondCount, err
	}

	return pondCount, nil
}

func (q *Queries) GetPondCountBySlug(ctx context.Context, slug string) (int, string, error) {
	var (
		pondCount int
		pondId    string
		err       error
	)

	row := q.db.QueryRowContext(ctx, querySelectCountBySlug, slug)
	if err = row.Scan(&pondId, &pondCount); err != nil {
		return pondCount, pondId, err
	}

	return pondCount, pondId, nil
}
