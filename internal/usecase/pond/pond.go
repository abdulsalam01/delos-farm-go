package pond

import (
	"context"
	"fmt"

	"github.com/abdulsalam/delos/helper"
	"github.com/abdulsalam/delos/internal/entity"
	"github.com/abdulsalam/delos/internal/entity/generic"
	"github.com/google/uuid"
)

func (u *Usecase) GetAll(ctx context.Context, args entity.PondRequestWithPagination) (entity.PondResponseWithPagination, error) {
	var (
		pondResp entity.PondResponseWithPagination
		pond     []entity.Pond
		err      error
	)

	pond, err = u.pondRepo.GetPondAll(ctx, args)
	if err != nil {
		return pondResp, err
	}

	cnt, err := u.pondRepo.GetPondCount(ctx)
	if err != nil {
		cnt = 0
	}

	farm, err := u.getFarmByFarmID(ctx, pond)
	if err != nil {
		return pondResp, err
	}

	var pondWithFarm []entity.Pond
	for i := 0; i < len(pond); i++ {
		if farm[i].ID == uuid.Nil {
			continue
		}

		pond[i].Farm = &farm[i]
		pondWithFarm = append(pondWithFarm, pond[i])
	}

	return entity.PondResponseWithPagination{
		Pond: pondWithFarm,
		PaginationResponse: generic.PaginationResponse{
			Limit:  args.Limit,
			Offset: args.Offset,
			Total:  cnt,
		},
	}, nil
}

func (u *Usecase) GetByID(ctx context.Context, id uuid.UUID) (entity.Pond, error) {
	var (
		pond entity.Pond
		err  error
	)

	pond, err = u.pondRepo.GetPondByID(ctx, id)
	if err != nil {
		return pond, err
	}

	pondWithFarm, err := u.getFarmByFarmID(ctx, []entity.Pond{pond})
	if err != nil {
		return pond, err
	}
	if len(pondWithFarm) > 0 {
		pond.Farm = &pondWithFarm[0]
	}

	return pond, nil
}

func (u *Usecase) GetBySlug(ctx context.Context, slug string) (entity.Pond, error) {
	var (
		pond entity.Pond
		err  error
	)

	pond, err = u.pondRepo.GetPondBySlug(ctx, slug)
	if err != nil {
		return pond, err
	}

	pondWithFarm, err := u.getFarmByFarmID(ctx, []entity.Pond{pond})
	if err != nil {
		return pond, err
	}
	if len(pondWithFarm) > 0 {
		pond.Farm = &pondWithFarm[0]
	}

	return pond, nil
}

func (u *Usecase) Create(ctx context.Context, args entity.PondRequest) (uuid.UUID, error) {
	// Mutex Lock.
	helper.MutexLockUnLock()

	var (
		pondId uuid.UUID
		err    error
	)

	slug, err := helper.StringToSlug(args.Name)
	if err != nil {
		return pondId, err
	}

	args.Slug = slug
	pondId, err = u.pondRepo.InsertPond(ctx, args.ToBaseEntity())
	if err != nil {
		return pondId, err
	}

	return pondId, nil
}

func (u *Usecase) UpdateByID(ctx context.Context, args entity.PondRequest) (int64, error) {
	// Mutex Lock.
	helper.MutexLockUnLock()

	var (
		rows int64
		err  error
	)

	rows, err = u.pondRepo.UpdatePondByID(ctx, args.ToBaseEntity())
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

	rows, err = u.pondRepo.DeletePondByID(ctx, id)
	if err != nil {
		return rows, err
	}

	return rows, nil
}

func (u *Usecase) Upsert(ctx context.Context, args entity.PondRequest) (string, error) {
	var (
		result string
		err    error
	)

	slug, err := helper.StringToSlug(args.Name)
	if err != nil {
		return result, err
	}

	rows, idPond, err := u.pondRepo.GetPondCountBySlug(ctx, slug)
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
		idUuid, err := uuid.Parse(idPond)
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
