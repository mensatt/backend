-- name: GetAllReviews :many
SELECT *
FROM review;

-- name: GetReviewByID :one
SELECT *
FROM review
WHERE id = $1;

-- name: CreateReview :one
INSERT INTO review (occurrence, display_name, stars, text)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: EditReview :one
UPDATE review
SET occurrence = $1, display_name = $2, stars = $3, text = $4, updated_at = NOW(), accepted_at = $5
WHERE id = $6
RETURNING *;

-- name: DeleteReview :one
DELETE FROM review
WHERE id = $1
RETURNING *;