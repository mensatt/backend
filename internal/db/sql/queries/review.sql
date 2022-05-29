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
