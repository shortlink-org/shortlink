-- CreateTable
CREATE TABLE "Stats" (
    "hash" TEXT NOT NULL,
    "count_redirect" INTEGER NOT NULL,
    "createdAt" TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP(3) NOT NULL
);

-- CreateIndex
CREATE UNIQUE INDEX "Stats.hash_unique" ON "Stats"("hash");
