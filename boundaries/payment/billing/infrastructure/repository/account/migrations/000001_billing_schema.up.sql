-- BILLING SCHEMA ======================================================================================================

-- for local development
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
-- CREATE EXTENSION IF NOT EXISTS pg_prewarm;
-- ALTER SYSTEM SET shared_preload_libraries = 'pg_prewarm';
CREATE SCHEMA IF NOT EXISTS billing;

-- ACCOUNT TABLE =======================================================================================================
CREATE TABLE billing.account(
    "id" UUID NOT NULL,
    "user_id" UUID NOT NULL,
    "tariff_id" UUID NOT NULL
) WITH (fillfactor = 100);

ALTER TABLE
    billing.account ADD PRIMARY KEY("id");

COMMENT ON COLUMN
    billing.account."user_id" IS 'Alias for GDPR';
