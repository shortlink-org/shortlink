-- INITIALIZATION ======================================================================================================

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- EVENTS TABLE ========================================================================================================
CREATE TABLE events(
    "aggregate_id" UUID NOT NULL,
    "aggregate_type" TEXT NOT NULL,
    "id" UUID NOT NULL,
    "payload" jsonb NOT NULL,
    "version" INTEGER NOT NULL,
    "created_at" TIMESTAMP(0) WITHOUT TIME ZONE NOT NULL
) WITH (fillfactor = 100);
ALTER TABLE
    events ADD PRIMARY KEY("id");

COMMENT ON COLUMN
    events."aggregate_id" IS 'is used to detect recurring events/messages. It stores the ID of the message/event whose processing generated this event.';

COMMENT ON COLUMN
    events."payload" IS '{
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
CREATE TABLE snapshots(
    "aggregate_id" UUID NOT NULL,
    "aggregate_type" TEXT NOT NULL,
    "aggregate_version" INTEGER NOT NULL,
    "payload" jsonb NOT NULL,
    "created_at" TIMESTAMP(0) WITHOUT TIME ZONE NOT NULL,
    "updated_at" TIMESTAMP(0) WITHOUT TIME ZONE NOT NULL
) WITH (fillfactor = 100);
CREATE UNIQUE INDEX "snapshots_aggregate_id_uindex" ON
    snapshots("aggregate_id");

-- AGGREGATE TABLE =====================================================================================================
CREATE TABLE aggregates(
    "id" UUID NOT NULL,
    "type" TEXT NOT NULL,
    "version" INTEGER NOT NULL,
    "created_at" TIMESTAMP(0) WITHOUT TIME ZONE NOT NULL,
    "updated_at" TIMESTAMP(0) WITHOUT TIME ZONE NOT NULL
) WITH (fillfactor = 100);
ALTER TABLE
    aggregates ADD PRIMARY KEY("id");
COMMENT ON COLUMN
    aggregates."version" IS 'Each time the entity is changed, the column is updated.';
ALTER TABLE
    events ADD CONSTRAINT "events_aggregate_id_foreign" FOREIGN KEY("aggregate_id") REFERENCES aggregates("id");
ALTER TABLE
    snapshots ADD CONSTRAINT "snapshots_aggregate_id_foreign" FOREIGN KEY("aggregate_id") REFERENCES aggregates("id");
