package pond

import (
	"github.com/abdulsalam/delos/helper"
	"github.com/abdulsalam/delos/internal/entity"
)

func (q *Queries) toBase(args entity.Pond) pondTable {
	args.CreatedAt = helper.GetTimeNow()
	args.UpdatedAt = helper.GetTimeNow()

	return pondTable{
		ID:          args.ID.String(),
		FarmID:      args.FarmID.String(),
		Name:        args.Name,
		Slug:        args.Slug,
		WaterSource: args.WaterSource,
		Size:        args.Size,

		// Extras.
		CreatedBy: args.MetaInfo.CreatedBy.String(),
		UpdatedBy: args.MetaInfo.UpdatedBy.String(),
		CreatedAt: args.CreatedAt,
		UpdatedAt: args.UpdatedAt,
	}
}
