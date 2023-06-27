package farm

import (
	"context"

	"github.com/abdulsalam/delos/internal/entity"
	"github.com/google/uuid"
)

func (q *Queries) GetFarmAll(ctx context.Context, args entity.FarmRequestWithPagination) ([]entity.Farm, error) {
	var (
		farm []entity.Farm
		err  error
	)

	row, err := q.db.QueryContext(ctx, querySelectByIDWithPagination, args.Limit, args.Offset)
	if err != nil {
		return farm, err
	}
	defer row.Close()

	for row.Next() {
		var result farmTable
		if err = row.Scan(
			&result.ID,
			&result.Name,
			&result.Slug,
			&result.Location,
			&result.Size,
			&result.Established,
			&result.Technologies,
			&result.Employees,
			&result.CreatedBy,
			&result.UpdatedBy,
			&result.CreatedAt,
			&result.UpdatedAt,
		); err != nil {
			return farm, err
		}

		data, err := result.ToEntity()
		if err != nil {
			continue
		}
		if data.ID == uuid.Nil {
			continue
		}

		farm = append(farm, data)
	}

	// Close rows scanner.
	if err := row.Close(); err != nil {
		return nil, err
	}
	if err := row.Err(); err != nil {
		return nil, err
	}

	return farm, nil
}

func (q *Queries) GetFarmByID(ctx context.Context, id uuid.UUID) (entity.Farm, error) {
	var (
		farm   entity.Farm
		result farmTable
		err    error
	)

	row := q.db.QueryRowContext(ctx, querySelectByID, id)
	if err = row.Scan(
		&result.ID,
		&result.Name,
		&result.Slug,
		&result.Location,
		&result.Size,
		&result.Established,
		&result.Technologies,
		&result.Employees,
		&result.CreatedBy,
		&result.UpdatedBy,
		&result.CreatedAt,
		&result.UpdatedAt,
	); err != nil {
		return farm, err
	}

	farm, err = result.ToEntity()
	if err != nil {
		return farm, err
	}
	if farm.ID == uuid.Nil {
		return farm, err
	}

	return farm, nil
}

func (q *Queries) GetFarmBySlug(ctx context.Context, slug string) (entity.Farm, error) {
	var (
		farm   entity.Farm
		result farmTable
		err    error
	)

	row := q.db.QueryRowContext(ctx, querySelectBySlug, slug)
	if err = row.Scan(&result.ID,
		&result.Name,
		&result.Slug,
		&result.Location,
		&result.Size,
		&result.Established,
		&result.Technologies,
		&result.Employees,
		&result.CreatedBy,
		&result.UpdatedBy,
		&result.CreatedAt,
		&result.UpdatedAt,
	); err != nil {
		return farm, err
	}

	farm, err = result.ToEntity()
	if err != nil {
		return farm, err
	}
	if farm.ID == uuid.Nil {
		return farm, err
	}

	return farm, nil
}

func (q *Queries) GetFarmCount(ctx context.Context) (int64, error) {
	var (
		farmCount int64
		farmId    string
		err       error
	)

	row := q.db.QueryRowContext(ctx, queryPartialSelectCountActiveOnly)
	if err = row.Scan(&farmId, &farmCount); err != nil {
		return farmCount, err
	}

	return farmCount, nil
}

func (q *Queries) GetFarmCountBySlug(ctx context.Context, slug string) (int, string, error) {
	var (
		farmCount int
		farmId    string
		err       error
	)

	row := q.db.QueryRowContext(ctx, querySelectCountBySlug, slug)
	if err = row.Scan(&farmId, &farmCount); err != nil {
		return farmCount, farmId, err
	}

	return farmCount, farmId, nil
}
