-- BILLING SCHEMA ======================================================================================================

-- for local development
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
-- CREATE EXTENSION IF NOT EXISTS pg_prewarm;
-- ALTER SYSTEM SET shared_preload_libraries = 'pg_prewarm';
CREATE SCHEMA IF NOT EXISTS billing;

-- TARIFF TABLE ========================================================================================================
CREATE TABLE tariff(
    "id" UUID NOT NULL,
    "name" VARCHAR(255) NOT NULL,
    "payload" jsonb NOT NULL,
    "created_at" TIMESTAMP(0) WITHOUT TIME ZONE NOT NULL,
    "updated_at" TIMESTAMP(0) WITHOUT TIME ZONE NOT NULL
) WITH (fillfactor = 100);
ALTER TABLE
    tariff ADD PRIMARY KEY("id");
COMMENT
ON COLUMN
  tariff."payload" IS '{
  "productNameA": "price",
  "productNameB": "price",
}';

-- AGGREGATE TABLE =====================================================================================================
-- We cann't use the same table for aggregate because of the different structure, it's need made from code side
-- ALTER TABLE
--    account ADD CONSTRAINT "account_tariff_id_foreign" FOREIGN KEY("tariff_id") REFERENCES tariff("id");
