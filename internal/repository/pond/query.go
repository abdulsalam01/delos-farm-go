package pond

const (
	tableDef = `ponds`

	// Queries.
	querySelectCountFromBaseRepo = `SELECT id, count(*) cnt`
	querySelectFromBaseRepo      = `SELECT 
									id, 
									farm_id,
									name,
									slug,
									size,
									water_source,
									created_by,
									updated_by,
									created_at,
									updated_at`

	queryPartialSelectActiveOnly      = querySelectFromBaseRepo + ` FROM ` + tableDef + ` WHERE deleted_at IS NULL `
	queryPartialSelectCountActiveOnly = querySelectCountFromBaseRepo + ` FROM ` + tableDef + ` WHERE deleted_at IS NULL `
	queryPartialOrderByCreatedAt      = ` ORDER BY created_at DESC `

	querySelectByID               = queryPartialSelectActiveOnly + ` AND id = $1`
	querySelectByFarmID           = queryPartialSelectActiveOnly + ` AND farm_id = $1`
	querySelectBySlug             = queryPartialSelectActiveOnly + ` AND slug = $1`
	querySelectCountByID          = queryPartialSelectCountActiveOnly + ` AND id = $1`
	querySelectCountBySlug        = queryPartialSelectCountActiveOnly + ` AND slug = $1`
	querySelectByIDWithPagination = queryPartialSelectActiveOnly + queryPartialOrderByCreatedAt + ` LIMIT $1 OFFSET $2`

	// Insert.
	queryInsertFromBaseRepo = `INSERT INTO ` + tableDef + `
							(
								id, 
								farm_id,
								name, 
								slug, 
								size, 
								water_source,
								created_by, 
								updated_by, 
								created_at, 
								updated_at
							) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) RETURNING id`
	// Update.
	queryUpdateFromBaseRepo = `	UPDATE ` + tableDef + `
								SET farm_id = $1,
									name = $2,
									size = $3,
									water_source = $4, 
									created_by = $5, 
									updated_by = $6, 
									created_at = $7, 
									updated_at = $8
								WHERE id = $9
								`
	// Soft delete.
	queryDeleteFromBaseRepo = `	UPDATE ` + tableDef + ` SET deleted_at = $1 WHERE id = $2`
)
