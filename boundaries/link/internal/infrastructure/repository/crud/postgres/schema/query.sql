-- name: CreateLink :exec
INSERT INTO link.links (url, hash, describe, json)
VALUES ($1, $2, $3, $4);

-- name: CreateLinks :copyfrom
INSERT INTO link.links (url, hash, describe, json)
VALUES ($1, $2, $3, $4);

-- name: GetLinkByHash :one
SELECT * FROM link.links
WHERE hash = $1;

-- name: GetLinks :many
SELECT * FROM link.links
LIMIT $1 OFFSET $2;

-- name: UpdateLink :execresult
UPDATE link.links
SET url = $1, hash = $2, describe = $3, json = $4::jsonb
WHERE id = $5;

-- name: DeleteLink :exec
DELETE FROM link.links
WHERE hash = $1;

