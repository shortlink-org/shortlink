-- Drop the maintenance function
SELECT partman.undo_partition('link.links_partitioned_by_created_at');

-- Drop the pg_partman extension; note that this will remove pg_partman and all related objects
-- Make sure this is what you want, especially if you have other partitioned tables managed by pg_partman
DROP EXTENSION IF EXISTS pg_partman CASCADE;

-- Drop the partman schema if it's empty (no other objects besides pg_partman)
-- Be cautious with this step if you have other objects in the partman schema
DROP SCHEMA IF EXISTS partman CASCADE;

-- Drop the partitioned table
DROP TABLE IF EXISTS link.links_partitioned_by_created_at;
