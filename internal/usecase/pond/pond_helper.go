package pond

import (
	"context"
	"log"

	"github.com/abdulsalam/delos/internal/entity"
)

func (u *Usecase) getFarmByFarmID(ctx context.Context, pond []entity.Pond) ([]entity.Farm, error) {
	var (
		pondWithFarm []entity.Farm
		err          error
	)

	for _, v := range pond {
		farm, err := u.farmRepo.GetFarmByID(ctx, v.FarmID)
		if err != nil {
			log.Printf("error when retrieve farm data %v", err)
			continue
		}

		pondWithFarm = append(pondWithFarm, farm)
	}
	return pondWithFarm, err
}
