-- name: CreateLink :execresult
INSERT INTO links (id, url, hash, `describe`, json)
  VALUES (?, ?, ?, ?, ?);

-- name: GetLinkByHash :one
SELECT * FROM links
  WHERE hash = ?;

-- name: GetLinks :many
SELECT * FROM links;

-- name: UpdateLink :execresult
UPDATE links
    SET url = ?, hash = ?, `describe` = ?, json = ?
WHERE id = ?;

-- name: DeleteLink :exec
DELETE FROM links
  WHERE hash = ?;
