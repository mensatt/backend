-- name: GetAllImages :many
SELECT *
FROM image;

-- name: GetImageByID :one
SELECT *
FROM image
WHERE id = $1;

-- name: GetImagesByDish :many
SELECT image.*
FROM image
JOIN occurrence ON (image.occurrence = occurrence.id)
JOIN dish ON (occurrence.dish = dish.id)
WHERE dish.id = $1;