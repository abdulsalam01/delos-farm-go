package farm

import (
	"context"
	"log"

	"github.com/abdulsalam/delos/internal/entity"
)

func (u *Usecase) getPondByFarmID(ctx context.Context, farm []entity.Farm) []entity.Farm {
	var farmWithPond []entity.Farm

	for _, v := range farm {
		pond, err := u.pondRepo.GetPondByFarmID(ctx, v.ID)
		if err != nil {
			log.Printf("error when retrieve pond data %v", err)
			continue
		}

		// Clear farms inside pond.
		var pondWithouthFarm []entity.Pond
		for _, f := range pond {
			f.Farm = nil

			pondWithouthFarm = append(pondWithouthFarm, f)
		}

		v.Pond = pondWithouthFarm
		farmWithPond = append(farmWithPond, v)
	}

	return farmWithPond
}
