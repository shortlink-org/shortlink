-- Create the partitioned table
CREATE TABLE link.links_partitioned_by_created_at
(
    id         uuid      DEFAULT gen_random_uuid() NOT NULL,
    url        text      NOT NULL,
    hash       hash      NOT NULL,
    describe   text,
    json       link      NOT NULL,
    created_at timestamp DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at timestamp DEFAULT CURRENT_TIMESTAMP NOT NULL,
    PRIMARY KEY (hash, created_at)
)
    PARTITION BY RANGE (created_at);

-- Create schema and extension for pg_partman
CREATE SCHEMA IF NOT EXISTS partman;
CREATE EXTENSION IF NOT EXISTS pg_partman SCHEMA partman;

-- Create parent table with pg_partman using 'native' partitioning
SELECT partman.create_parent(
    p_parent_table => 'link.links_partitioned_by_created_at',
    p_control => 'created_at'::text,
    p_type => 'native', -- 'native' or 'range'
    p_interval => '1 day',
    p_template_table => 'link.links_partitioned_by_created_at'
);

-- Update retention policy
UPDATE partman.part_config
SET retention = '10 days'
WHERE parent_table = 'link.links_partitioned_by_created_at';

-- Run maintenance
SELECT partman.run_maintenance();

-- Migrate existing data into the partitioned table
INSERT INTO link.links_partitioned_by_created_at
SELECT * FROM link.links;
DROP TABLE link.links;
ALTER TABLE link.links_partitioned_by_created_at RENAME TO links;
