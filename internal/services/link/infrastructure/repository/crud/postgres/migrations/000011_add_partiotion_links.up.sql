CREATE TABLE link.links_partitioned_by_created_at
(
    id         uuid      DEFAULT link.uuid_generate_v4() NOT NULL,
    url        text                                      NOT NULL,
    hash       link.hash                                 NOT NULL,
    describe   text,
    json       link.link                                 NOT NULL,
    created_at timestamp DEFAULT CURRENT_TIMESTAMP       NOT NULL,
    updated_at timestamp DEFAULT CURRENT_TIMESTAMP       NOT NULL,
    PRIMARY KEY (created_at)
)
    PARTITION BY RANGE (hash, created_at);

CREATE SCHEMA IF NOT EXISTS partman;
CREATE EXTENSION IF NOT EXISTS pg_partman SCHEMA partman;

SELECT partman.create_parent('link.links_partitioned_by_created_at', 'created_at', 'native', 'daily');
UPDATE partman.part_config SET retention = '10 days' WHERE parent_table = 'link.links_partitioned_by_created_at';

SELECT partman.run_maintenance();
