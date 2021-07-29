-- BILLING SCHEMA ======================================================================================================

CREATE SCHEMA billing;

COMMENT ON SCHEMA billing IS 'Billing schema';

-- SNAPSHOTS TARIFF ====================================================================================================
CREATE TABLE billing.tariff(
    "id" UUID NOT NULL,
    "name" VARCHAR(255) NOT NULL,
    "payload" jsonb NOT NULL,
    "created_at" TIMESTAMP(0) WITHOUT TIME ZONE NOT NULL,
    "updated_at" TIMESTAMP(0) WITHOUT TIME ZONE NOT NULL
);
ALTER TABLE
    billing.tariff ADD PRIMARY KEY("id");
COMMENT
ON COLUMN
    billing.tariff."payload" IS '{
  "productNameA": "price",
  "productNameB": "price",
}';

-- ACCOUNT TABLE =======================================================================================================
CREATE TABLE billing.account(
    "id" UUID NOT NULL,
    "user_id" UUID NOT NULL,
    "tariff_id" UUID NOT NULL
);

ALTER TABLE
    billing.account ADD PRIMARY KEY("id");

COMMENT ON COLUMN
    billing.account."user_id" IS 'Alias for GDPR';

-- EVENTS TABLE ========================================================================================================
CREATE TABLE billing.events(
    "aggregate_id" UUID NOT NULL,
    "aggregate_type" TEXT NOT NULL,
    "id" UUID NOT NULL,
    "payload" jsonb NOT NULL,
    "version" INTEGER NOT NULL,
    "created_at" TIMESTAMP(0) WITHOUT TIME ZONE NOT NULL
);
ALTER TABLE
    billing.events ADD PRIMARY KEY("id");

COMMENT ON COLUMN
    billing.events."aggregate_id" IS 'используется для обнаружения повторяющихся со-
бытий/сообщений. Он хранит ID сообщения/события, обработка которого сгене­
рировала это событие.';

COMMENT ON COLUMN
    billing.events."payload" IS '{
  event: "BALANCE_ADD"
  payload: {amount: 5000}
},
{
  event: "NEW_ORDER",
  payload: {list: [
    name: productNameA,
    price: 1500
  ]}
}';

-- SNAPSHOTS TABLE =====================================================================================================
CREATE TABLE billing.snapshots(
    "aggregate_id" UUID NOT NULL,
    "aggregate_type" TEXT NOT NULL,
    "agregate_version" INTEGER NOT NULL,
    "payload" jsonb NOT NULL,
    "created_at" TIMESTAMP(0) WITHOUT TIME ZONE NOT NULL,
    "updated_at" TIMESTAMP(0) WITHOUT TIME ZONE NOT NULL
);
CREATE INDEX "snapshots_aggregate_id_index" ON
    billing.snapshots("aggregate_id");
CREATE INDEX "snapshots_aggregate_type_index" ON
    billing.snapshots("aggregate_type");
CREATE INDEX "snapshots_agregate_version_index" ON
    billing.snapshots("agregate_version");


-- AGGREGATE TABLE =====================================================================================================
CREATE TABLE billing.aggregates(
    "id" UUID NOT NULL,
    "type" TEXT NOT NULL,
    "version" INTEGER NOT NULL,
    "created_at" TIMESTAMP(0) WITHOUT TIME ZONE NOT NULL,
    "updated_at" TIMESTAMP(0) WITHOUT TIME ZONE NOT NULL
);
ALTER TABLE
    billing.aggregates ADD PRIMARY KEY("id");
CREATE INDEX "aggregate_type_index" ON
    billing.aggregates("type");
COMMENT ON COLUMN
    billing.aggregates."version" IS 'При каждом изменении сущности обновляется столбец';
ALTER TABLE
    billing.account ADD CONSTRAINT "account_tariff_id_foreign" FOREIGN KEY("tariff_id") REFERENCES billing.tariff("id");
ALTER TABLE
    billing.events ADD CONSTRAINT "events_aggregate_id_foreign" FOREIGN KEY("aggregate_id") REFERENCES billing.aggregates("id");
ALTER TABLE
    billing.snapshots ADD CONSTRAINT "snapshots_aggregate_id_foreign" FOREIGN KEY("aggregate_id") REFERENCES billing.aggregates("id");
