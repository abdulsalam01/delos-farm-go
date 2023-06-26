package db

const (
	queryCreateFarmTable = `
		CREATE TABLE IF NOT EXISTS farms (
			id VARCHAR PRIMARY KEY,
			name VARCHAR,
			slug VARCHAR UNIQUE,
			location VARCHAR,
			size REAL,
			established NUMERIC,
			technologies VARCHAR,
			employees INTEGER,
		
			created_by VARCHAR DEFAULT '00000000-0000-0000-0000-000000000000',
			updated_by VARCHAR DEFAULT '00000000-0000-0000-0000-000000000000',
		
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			deleted_at TIMESTAMP
		);
	`

	queryCreatePondTable = `
		CREATE TABLE IF NOT EXISTS ponds (
			id VARCHAR PRIMARY KEY,
			farm_id VARCHAR,
			name VARCHAR,
			slug VARCHAR UNIQUE,
			size REAL,
			water_source INTEGER,
		
			created_by VARCHAR DEFAULT '00000000-0000-0000-0000-000000000000',
			updated_by VARCHAR DEFAULT '00000000-0000-0000-0000-000000000000',
		
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			deleted_at TIMESTAMP
		);
	`

	queryCreateIndexFarmCreatedAt = `
		CREATE INDEX IF NOT EXISTS farms_idx_created_at ON farms (created_at);
	`
	queryCreateIndexFarmDeletedAt = `
		CREATE INDEX IF NOT EXISTS farms_idx_deleted_at ON farms (deleted_at);
	`
	queryCreateIndexPondCreatedAt = `
		CREATE INDEX IF NOT EXISTS ponds_idx_created_at ON ponds (created_at);
	`
	queryCreateIndexPondDeletedAt = `
		CREATE INDEX IF NOT EXISTS ponds_idx_deleted_at ON ponds (deleted_at);
	`
)
