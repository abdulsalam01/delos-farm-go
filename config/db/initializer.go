package db

import (
	"context"
)

func (q *Queries) InitFarmTable(ctx context.Context) error {
	_, err := q.db.ExecContext(ctx, queryCreateFarmTable)
	return err
}

func (q *Queries) InitPondTable(ctx context.Context) error {
	_, err := q.db.ExecContext(ctx, queryCreatePondTable)
	return err
}

func (q *Queries) Index(ctx context.Context) error {
	var (
		err error
	)

	if _, err = q.db.QueryContext(ctx, queryCreateIndexFarmCreatedAt); err != nil {
		return err
	}
	if _, err = q.db.QueryContext(ctx, queryCreateIndexFarmDeletedAt); err != nil {
		return err
	}
	if _, err = q.db.QueryContext(ctx, queryCreateIndexPondCreatedAt); err != nil {
		return err
	}
	if _, err = q.db.QueryContext(ctx, queryCreateIndexPondDeletedAt); err != nil {
		return err
	}

	return nil
}
