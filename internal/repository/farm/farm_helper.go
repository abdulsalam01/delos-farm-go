package farm

import (
	"strings"

	"github.com/abdulsalam/delos/helper"
	"github.com/abdulsalam/delos/internal/entity"
)

func (q *Queries) toBase(args entity.Farm) farmTable {
	args.CreatedAt = helper.GetTimeNow()
	args.UpdatedAt = helper.GetTimeNow()

	return farmTable{
		ID:           args.ID.String(),
		Name:         args.Name,
		Slug:         args.Slug,
		Location:     args.Location,
		Size:         args.Size,
		Established:  helper.ConvertTimeToInt64(args.Established),
		Technologies: strings.Join(args.Technologies, ","),
		Employees:    args.Employees,

		// Extras.
		CreatedBy: args.MetaInfo.CreatedBy.String(),
		UpdatedBy: args.MetaInfo.UpdatedBy.String(),
		CreatedAt: args.CreatedAt,
		UpdatedAt: args.UpdatedAt,
	}
}
