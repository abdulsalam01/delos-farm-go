package farm

import (
	"context"

	"github.com/abdulsalam/delos/internal/entity"
	"github.com/google/uuid"
)

type farmRepoResource interface {
	// Get.
	GetFarmAll(ctx context.Context, args entity.FarmRequestWithPagination) ([]entity.Farm, error)
	GetFarmByID(ctx context.Context, id uuid.UUID) (entity.Farm, error)
	GetFarmBySlug(ctx context.Context, slug string) (entity.Farm, error)
	GetFarmCount(ctx context.Context) (int64, error)
	GetFarmCountBySlug(ctx context.Context, slug string) (int, string, error)
	// Insert.
	InsertFarm(ctx context.Context, args entity.Farm) (uuid.UUID, error)
	// Update.
	UpdateFarmByID(ctx context.Context, args entity.Farm) (int64, error)
	// Delete.
	DeleteFarmByID(ctx context.Context, id uuid.UUID) (int64, error)
}

type pondRepoResouce interface {
	GetPondByFarmID(ctx context.Context, farmId uuid.UUID) ([]entity.Pond, error)
}

type Usecase struct {
	farmRepo farmRepoResource
	pondRepo pondRepoResouce
}
