package pond

import (
	"context"
	"database/sql"
	"time"

	"github.com/abdulsalam/delos/internal/entity"
	"github.com/abdulsalam/delos/internal/entity/generic"
	"github.com/google/uuid"
)

type dbTX interface {
	ExecContext(context.Context, string, ...interface{}) (sql.Result, error)
	PrepareContext(context.Context, string) (*sql.Stmt, error)
	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
}

type Queries struct {
	db dbTX
}

type pondTable struct {
	ID          string
	FarmID      string
	Name        string
	Slug        string // Should be unique.
	Size        float64
	WaterSource string

	// Extra.
	CreatedBy string
	UpdatedBy string

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

// Implement the Scanner interface for NullTime
func (nt *pondTable) Scan(value interface{}) error {
	if value == nil {
		nt.DeletedAt = time.Time{}
		return nil
	}

	nt.DeletedAt = value.(time.Time)
	return nil
}

func (f *pondTable) ToEntity() (entity.Pond, error) {
	var (
		pond entity.Pond
		err  error
	)

	// Parsing.
	id, err := uuid.Parse(f.ID)
	if err != nil {
		return pond, err
	}
	farmId, err := uuid.Parse(f.FarmID)
	if err != nil {
		return pond, err
	}
	createdBy, err := uuid.Parse(f.CreatedBy)
	if err != nil {
		return pond, err
	}
	updatedBy, err := uuid.Parse(f.UpdatedBy)
	if err != nil {
		return pond, err
	}

	pond = entity.Pond{
		ID:          id,
		FarmID:      farmId,
		Name:        f.Name,
		Slug:        f.Slug,
		Size:        f.Size,
		WaterSource: f.WaterSource,
		MetaInfo: generic.MetaInfo{
			CreatedBy: createdBy,
			UpdatedBy: updatedBy,
			CreatedAt: f.CreatedAt,
			UpdatedAt: f.UpdatedAt,
		},
	}

	return pond, err
}
