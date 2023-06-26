package farm

import (
	"context"
	"database/sql"
	"strings"
	"time"

	"github.com/abdulsalam/delos/helper"
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

type farmTable struct {
	ID           string
	Name         string
	Slug         string // Should be unique.
	Location     string
	Size         float64
	Established  int64 // Timestamps.
	Technologies string
	Employees    int

	// Extra.
	CreatedBy string
	UpdatedBy string

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

// Implement the Scanner interface for NullTime
func (nt *farmTable) Scan(value interface{}) error {
	if value == nil {
		nt.DeletedAt = time.Time{}
		return nil
	}

	nt.DeletedAt = value.(time.Time)
	return nil
}

func (f *farmTable) ToEntity() (entity.Farm, error) {
	var (
		farm entity.Farm
		err  error
	)

	// Parsing.
	id, err := uuid.Parse(f.ID)
	if err != nil {
		return farm, err
	}
	createdBy, err := uuid.Parse(f.CreatedBy)
	if err != nil {
		return farm, err
	}
	updatedBy, err := uuid.Parse(f.UpdatedBy)
	if err != nil {
		return farm, err
	}

	farm = entity.Farm{
		ID:           id,
		Name:         f.Name,
		Slug:         f.Slug,
		Location:     f.Location,
		Size:         f.Size,
		Established:  helper.ConvertInt64ToTime(f.Established),
		Technologies: strings.Split(f.Technologies, ","),
		Employees:    f.Employees,
		MetaInfo: generic.MetaInfo{
			CreatedBy: createdBy,
			UpdatedBy: updatedBy,
			CreatedAt: f.CreatedAt,
			UpdatedAt: f.UpdatedAt,
		},
	}

	return farm, err
}
