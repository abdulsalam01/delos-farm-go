-- farms
CREATE TABLE farms (
    id VARCHAR PRIMARY KEY,
    name VARCHAR,
    slug VARCHAR UNIQUE,
    location VARCHAR,
    size REAL,
    established NUMERIC,
    technologies VARCHAR,
    employees INTEGER,

    -- extras info
    created_by VARCHAR DEFAULT '00000000-0000-0000-0000-000000000000',
    updated_by VARCHAR DEFAULT '00000000-0000-0000-0000-000000000000',

    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,

    INDEX farm_idx_created_at (created_at),
    INDEX farm_idx_deleted_at (deleted_at)
);

-- ponds
CREATE TABLE ponds (
    id VARCHAR PRIMARY KEY,
    farm_id VARCHAR,
    name VARCHAR,
    slug VARCHAR UNIQUE,
    size REAL,
    water_source INTEGER,

    -- extras info
    created_by VARCHAR DEFAULT '00000000-0000-0000-0000-000000000000',
    updated_by VARCHAR DEFAULT '00000000-0000-0000-0000-000000000000',

    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,

    INDEX pond_idx_created_at (created_at),
    INDEX pond_idx_deleted_at (deleted_at)
);