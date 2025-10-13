-- name: CreateFeed :one
INSERT INTO feeds (id, created_at, updated_at, name, url, user_id)
VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6
)
RETURNING *;

-- name: GetFeed :one
SELECT f.*, u.name as user_name
FROM feeds as f
LEFT JOIN users as u
  ON (f.user_id = u.id)
WHERE url = $1;

-- name: GetFeeds :many
SELECT f.*, u.name as user_name
FROM feeds as f
LEFT JOIN users as u
  ON (f.user_id = u.id);
