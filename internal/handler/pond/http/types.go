package http

import (
	"context"

	"github.com/abdulsalam/delos/internal/entity"
	"github.com/google/uuid"
)

type FarmUsecase interface {
	GetAll(ctx context.Context, args entity.FarmRequestWithPagination) (entity.FarmResponseWithPagination, error)
	GetByID(ctx context.Context, id uuid.UUID) (entity.Farm, error)
	GetBySlug(ctx context.Context, slug string) (entity.Farm, error)

	Create(ctx context.Context, args entity.FarmRequest) (uuid.UUID, error)
	UpdateByID(ctx context.Context, args entity.FarmRequest) (int64, error)
	DeleteByID(ctx context.Context, id uuid.UUID) (int64, error)
	Upsert(ctx context.Context, args entity.FarmRequest) (string, error)
}

type PondUsecase interface {
	GetAll(ctx context.Context, args entity.PondRequestWithPagination) (entity.PondResponseWithPagination, error)
	GetByID(ctx context.Context, id uuid.UUID) (entity.Pond, error)
	GetBySlug(ctx context.Context, slug string) (entity.Pond, error)

	Create(ctx context.Context, args entity.PondRequest) (uuid.UUID, error)
	UpdateByID(ctx context.Context, args entity.PondRequest) (int64, error)
	DeleteByID(ctx context.Context, id uuid.UUID) (int64, error)
	Upsert(ctx context.Context, args entity.PondRequest) (string, error)
}
type Handler struct {
	farmUsecase FarmUsecase
	pondUsecase PondUsecase
}
