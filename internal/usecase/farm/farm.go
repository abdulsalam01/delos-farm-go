package farm

import (
	"context"
	"fmt"

	"github.com/abdulsalam/delos/helper"
	"github.com/abdulsalam/delos/internal/entity"
	"github.com/abdulsalam/delos/internal/entity/generic"
	"github.com/google/uuid"
)

func (u *Usecase) GetAll(ctx context.Context, args entity.FarmRequestWithPagination) (entity.FarmResponseWithPagination, error) {
	var (
		farmResp entity.FarmResponseWithPagination
		farm     []entity.Farm
		err      error
	)

	farm, err = u.farmRepo.GetFarmAll(ctx, args)
	if err != nil {
		return farmResp, err
	}

	cnt, err := u.farmRepo.GetFarmCount(ctx)
	if err != nil {
		return farmResp, err
	}

	// Accept as array.
	farmWithPond := u.getPondByFarmID(ctx, farm)
	return entity.FarmResponseWithPagination{
		Farm: farmWithPond,
		PaginationResponse: generic.PaginationResponse{
			Limit:  args.Limit,
			Offset: args.Offset,
			Total:  cnt,
		},
	}, nil
}

func (u *Usecase) GetByID(ctx context.Context, id uuid.UUID) (entity.Farm, error) {
	var (
		farm entity.Farm
		err  error
	)

	farm, err = u.farmRepo.GetFarmByID(ctx, id)
	if err != nil {
		return farm, err
	}

	// Accept as array.
	farmWithPond := u.getPondByFarmID(ctx, []entity.Farm{farm})
	if len(farmWithPond) > 0 {
		return farmWithPond[0], nil
	}

	return farm, nil
}

func (u *Usecase) GetBySlug(ctx context.Context, slug string) (entity.Farm, error) {
	var (
		farm entity.Farm
		err  error
	)

	farm, err = u.farmRepo.GetFarmBySlug(ctx, slug)
	if err != nil {
		return farm, err
	}

	// Accept as array.
	farmWithPond := u.getPondByFarmID(ctx, []entity.Farm{farm})
	if len(farmWithPond) > 0 {
		return farmWithPond[0], nil
	}

	return farm, nil
}

func (u *Usecase) Create(ctx context.Context, args entity.FarmRequest) (uuid.UUID, error) {
	// Mutex Lock.
	helper.MutexLockUnLock()

	var (
		farmId uuid.UUID
		err    error
	)

	slug, err := helper.StringToSlug(args.Name)
	if err != nil {
		return farmId, err
	}

	args.Slug = slug
	farmId, err = u.farmRepo.InsertFarm(ctx, args.ToBaseEntity())
	if err != nil {
		return farmId, err
	}

	return farmId, nil
}

func (u *Usecase) UpdateByID(ctx context.Context, args entity.FarmRequest) (int64, error) {
	// Mutex Lock.
	helper.MutexLockUnLock()

	var (
		rows int64
		err  error
	)

	rows, err = u.farmRepo.UpdateFarmByID(ctx, args.ToBaseEntity())
	if err != nil {
		return rows, err
	}

	return rows, nil
}

func (u *Usecase) DeleteByID(ctx context.Context, id uuid.UUID) (int64, error) {
	// Mutex Lock.
	helper.MutexLockUnLock()

	var (
		rows int64
		err  error
	)

	rows, err = u.farmRepo.DeleteFarmByID(ctx, id)
	if err != nil {
		return rows, err
	}

	return rows, nil
}

func (u *Usecase) Upsert(ctx context.Context, args entity.FarmRequest) (string, error) {
	var (
		result string
		err    error
	)

	slug, err := helper.StringToSlug(args.Name)
	if err != nil {
		return result, err
	}

	rows, idFarm, err := u.farmRepo.GetFarmCountBySlug(ctx, slug)
	if err != nil {
		return result, err
	}

	// Do Upsert.
	var (
		rowsAffected int64
		idRaw        uuid.UUID
	)

	if rows > 0 { // Do update if rows > 0
		// Convert to uuid.
		idUuid, err := uuid.Parse(idFarm)
		if err != nil {
			return result, err
		}

		// Attach.
		args.ID = idUuid
		rowsAffected, err = u.UpdateByID(ctx, args)
		result = fmt.Sprintf("%d", rowsAffected)
	} else { // Do insert.
		idRaw, err = u.Create(ctx, args)
		result = idRaw.String()
	}

	return result, err
}
