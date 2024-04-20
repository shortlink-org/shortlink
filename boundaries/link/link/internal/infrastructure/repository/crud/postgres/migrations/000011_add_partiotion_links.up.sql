CREATE TABLE link.links_partitioned_by_created_at
(
    id         uuid      DEFAULT gen_random_uuid() NOT NULL,
    url        text                                 NOT NULL,
    hash       hash                                 NOT NULL,
    describe   text,
    json       link                                 NOT NULL,
    created_at timestamp DEFAULT CURRENT_TIMESTAMP  NOT NULL,
    updated_at timestamp DEFAULT CURRENT_TIMESTAMP  NOT NULL,
    PRIMARY KEY (hash, created_at)
)
    PARTITION BY RANGE (created_at);

CREATE SCHEMA IF NOT EXISTS partman;
CREATE EXTENSION IF NOT EXISTS pg_partman SCHEMA partman;

--SELECT partman.create_parent('link.links_partitioned_by_created_at', 'created_at', '1 days', 'range');
-- TODO: wait version 5.1 https://github.com/pgpartman/pg_partman/issues/265
SELECT partman.create_parent(
        p_parent_table => 'link.links_partitioned_by_created_at',
        p_control => 'created_at'::text,
        p_type =>'native',
        p_interval => '1 day',
        p_template_table := 'link.links_partitioned_by_created_at'
      );
UPDATE partman.part_config SET retention = '10 days' WHERE parent_table = 'link.links_partitioned_by_created_at';

SELECT partman.run_maintenance();

INSERT INTO link.links_partitioned_by_created_at SELECT * FROM link.links;
DROP TABLE link.links;
ALTER TABLE link.links_partitioned_by_created_at RENAME TO links;
