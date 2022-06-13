-- name: GetAllImages :many
SELECT *
FROM image;

-- name: GetImageByID :one
SELECT *
FROM image
WHERE id = $1;

-- name: GetImageStoreIDByID :one
SELECT image_store_id
FROM image
WHERE id = $1;

-- name: CreateImage :one
INSERT INTO image (image_store_id, occurrence, display_name, description)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: UpdateImage :one
UPDATE image
SET 
    occurrence = COALESCE(sqlc.narg('occurrence'), occurrence),
    display_name = COALESCE(sqlc.narg('display_name'), display_name),
    description = COALESCE(sqlc.narg('description'), description),
    updated_at = NOW(),
    accepted_at = COALESCE(sqlc.narg('accepted_at'), accepted_at)
WHERE id = $1
RETURNING *;

-- name: DeleteImage :one
DELETE FROM image
WHERE id = $1
RETURNING *;
