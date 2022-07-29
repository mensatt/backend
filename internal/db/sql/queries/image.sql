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
JOIN review ON (image.review = review.id)
JOIN occurrence ON (review.occurrence = occurrence.id)
JOIN dish ON (occurrence.dish = dish.id)
WHERE dish.id = $1;

-- name: GetImagesByOccurrence :many
SELECT image.*
FROM image
JOIN review ON (image.review = review.id)
JOIN occurrence ON (review.occurrence = occurrence.id)
WHERE occurrence.id = $1;

-- name: GetImagesByReview :many
SELECT image.*
FROM image JOIN review ON (image.review = review.id)
WHERE review.id = $1;

-- name: GetImageStoreIDByID :one
SELECT image_store_id
FROM image
WHERE id = $1;

-- name: AddImageToReview :one
INSERT INTO image (review, image_store_id)
VALUES ($1, $2)
RETURNING *;

-- name: AddMultipleImagesToReview :copyfrom
INSERT INTO image (review, image_store_id)
VALUES ($1, $2);

-- name: DeleteImage :one
DELETE FROM image
WHERE id = $1
RETURNING *;
