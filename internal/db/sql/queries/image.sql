-- name: GetAllImages :many
SELECT *
FROM image;

-- name: GetImageByID :one
SELECT *
FROM image
WHERE id = $1;