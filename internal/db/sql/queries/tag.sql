-- name: GetAllTags :many
SELECT *
FROM tag;

-- name: GetTagByKey :one
SELECT *
FROM tag
WHERE key = $1;

-- name: CreateTag :one
INSERT INTO tag (key, name, description, short_name, priority, is_allergy)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;