-- DROP FOREIGN KEY CONSTRAINTS ========================================================================================
ALTER TABLE events DROP CONSTRAINT IF EXISTS "events_aggregate_id_foreign";
ALTER TABLE snapshots DROP CONSTRAINT IF EXISTS "snapshots_aggregate_id_foreign";

-- DROP TABLES =========================================================================================================
DROP TABLE IF EXISTS events;
DROP TABLE IF EXISTS snapshots;
DROP TABLE IF EXISTS aggregates;
