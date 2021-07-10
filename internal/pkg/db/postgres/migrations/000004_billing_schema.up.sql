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
    "user_id" UUID NOT NULL,
    "request_id" INTEGER NOT NULL,
    "event_id" UUID NOT NULL,
    "event_type" VARCHAR(255) NOT NULL,
    "event_payload" jsonb NOT NULL,
    "entity_id" INTEGER NOT NULL,
    "entity_type" INTEGER NOT NULL
);
ALTER TABLE
    billing.events ADD PRIMARY KEY("event_id");
COMMENT
ON COLUMN
    billing.events."request_id" IS 'используется для обнаружения повторяющихся со-
бытий/сообщений. Он хранит ID сообщения/события, обработка которого сгене­
рировала это событие.';
COMMENT
ON COLUMN
    billing.events."event_payload" IS '{
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

-- ENTITIES TABLE ======================================================================================================
CREATE TABLE billing.entities(
    "id" INTEGER NOT NULL,
    "type" INTEGER NOT NULL,
    "version" VARCHAR(255) NOT NULL
);
ALTER TABLE
    billing.entities ADD PRIMARY KEY("id");
CREATE INDEX "entities_type_index" ON
    billing.entities("type");
COMMENT
ON COLUMN
    billing.entities."version" IS 'При каждом изменении сущности обновляется столбец';
ALTER TABLE
    billing.events ADD CONSTRAINT "events_user_id_foreign" FOREIGN KEY("user_id") REFERENCES billing.account("id");
ALTER TABLE
    billing.account ADD CONSTRAINT "account_tariff_id_foreign" FOREIGN KEY("tariff_id") REFERENCES billing.tariff("id");
ALTER TABLE
    billing.events ADD CONSTRAINT "events_entity_id_foreign" FOREIGN KEY("entity_id") REFERENCES billing.entities("id");

-- SNAPSHOTS TABLE =====================================================================================================
CREATE TABLE billing.snapshots(
    "user_id" UUID NOT NULL,
    "entity_id" UUID NOT NULL,
    "entity_type" DECIMAL(8, 2) NOT NULL,
    "entity_version" INTEGER NOT NULL,
    "snapshot_type" INTEGER NOT NULL,
    "snaphot_payload" INTEGER NOT NULL,
    "created_at" TIMESTAMP(0) WITHOUT TIME ZONE NOT NULL,
    "updated_at" TIMESTAMP(0) WITHOUT TIME ZONE NOT NULL
);
CREATE INDEX "snapshots_entity_id_index" ON
    billing.snapshots("entity_id");
CREATE INDEX "snapshots_entity_type_index" ON
    billing.snapshots("entity_type");
CREATE INDEX "snapshots_entity_version_index" ON
    billing.snapshots("entity_version");
