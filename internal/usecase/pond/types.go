package pond

import (
	"context"

	"github.com/abdulsalam/delos/internal/entity"
	"github.com/google/uuid"
)

type farmRepoResource interface {
	// Get.
	GetFarmByID(ctx context.Context, id uuid.UUID) (entity.Farm, error)
}

type pondRepoResource interface {
	// Get.
	GetPondAll(ctx context.Context, args entity.PondRequestWithPagination) ([]entity.Pond, error)
	GetPondByID(ctx context.Context, id uuid.UUID) (entity.Pond, error)
	GetPondBySlug(ctx context.Context, slug string) (entity.Pond, error)
	GetPondCount(ctx context.Context) (int64, error)
	GetPondCountBySlug(ctx context.Context, slug string) (int, string, error)
	// Insert.
	InsertPond(ctx context.Context, args entity.Pond) (uuid.UUID, error)
	// Update.
	UpdatePondByID(ctx context.Context, args entity.Pond) (int64, error)
	// Delete.
	DeletePondByID(ctx context.Context, id uuid.UUID) (int64, error)
}

type Usecase struct {
	farmRepo farmRepoResource
	pondRepo pondRepoResource
}
