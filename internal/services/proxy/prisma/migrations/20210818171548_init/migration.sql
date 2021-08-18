-- CreateTable
CREATE TABLE "stats" (
    "hash" TEXT NOT NULL,
    "count_redirect" INTEGER NOT NULL DEFAULT 0,
    "createdAt" TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP(3) NOT NULL
);

-- CreateIndex
CREATE UNIQUE INDEX "stats.hash_unique" ON "stats"("hash");

-- Move this table
ALTER TABLE stats
    SET SCHEMA shortlink;
