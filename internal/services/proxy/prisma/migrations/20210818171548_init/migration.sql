-- CreateTable
CREATE TABLE "stats" (
    "hash" varchar(9) NOT NULL,
    "count_redirect" INTEGER NOT NULL DEFAULT 0,
    "createdAt" TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP(3) NOT NULL
) WITH (fillfactor = 100);

-- CreateIndex
CREATE UNIQUE INDEX "stats.hash_unique" ON "stats"("hash");

-- Move this table
ALTER TABLE stats
    SET SCHEMA shortlink;

-- Add foreign key
alter table shortlink.stats
	add constraint stats_links_hash_fk
		foreign key (hash) references shortlink.links (hash)
			on delete cascade;
