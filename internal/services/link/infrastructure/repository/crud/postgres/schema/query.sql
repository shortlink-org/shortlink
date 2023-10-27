-- name: CreateLink :execresult
INSERT INTO link.links (id, url, hash, describe, json)
VALUES ($1, $2, $3, $4, $5);

-- name: GetLinkByHash :one
SELECT * FROM link.links
WHERE hash = $1;

-- name: GetLinks :many
SELECT * FROM link.links;

-- name: UpdateLink :execresult
UPDATE link.links
SET url = $1, hash = $2, describe = $3, json = $4
WHERE id = $5;

-- name: DeleteLink :exec
DELETE FROM link.links
WHERE hash = $1;

