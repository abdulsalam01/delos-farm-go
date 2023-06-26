package farm

const (
	tableDef = `farms`

	// Queries.
	querySelectCountFromBaseRepo = `SELECT id, count(*) cnt`
	querySelectFromBaseRepo      = `SELECT 
									id, 
									name,
									slug,
									location,
									size,
									established,
									technologies,
									employees,
									created_by,
									updated_by,
									created_at,
									updated_at`

	queryPartialSelectActiveOnly      = querySelectFromBaseRepo + ` FROM ` + tableDef + ` WHERE deleted_at IS NULL `
	queryPartialSelectCountActiveOnly = querySelectCountFromBaseRepo + ` FROM ` + tableDef + ` WHERE deleted_at IS NULL `

	querySelectByID               = queryPartialSelectActiveOnly + ` AND id = $1`
	querySelectBySlug             = queryPartialSelectActiveOnly + ` AND slug = $1`
	querySelectCountByID          = queryPartialSelectCountActiveOnly + ` AND id = $1`
	querySelectCountBySlug        = queryPartialSelectCountActiveOnly + ` AND slug = $1`
	querySelectByIDWithPagination = queryPartialSelectActiveOnly + ` LIMIT $1 OFFSET $2`

	// Insert.
	queryInsertFromBaseRepo = `INSERT INTO ` + tableDef + `
							(
								id, 
								name, 
								slug, 
								location, 
								size, 
								established, 
								technologies, 
								employees,
								created_by, 
								updated_by, 
								created_at, 
								updated_at
							) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12) RETURNING id`
	// Update.
	queryUpdateFromBaseRepo = `	UPDATE ` + tableDef + `
								SET name = $1, 
									location = $2, 
									size = $3, 
									established = $4, 
									technologies = $5, 
									employees = $6,
									created_by = $7, 
									updated_by = $8, 
									created_at = $9, 
									updated_at = $10
								WHERE id = $11
								`
	// Soft delete.
	queryDeleteFromBaseRepo = `	UPDATE ` + tableDef + ` SET deleted_at = $1 WHERE id = $2`
)
