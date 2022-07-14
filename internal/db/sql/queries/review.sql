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

-- name: UpdateReview :one
UPDATE review
SET 
    occurrence = COALESCE(sqlc.narg('occurrence'), occurrence),
    display_name = COALESCE(sqlc.narg('display_name'), display_name),
    stars = COALESCE(sqlc.narg('stars'), stars),
    text = COALESCE(sqlc.narg('text'), text),
    updated_at = NOW(),
    accepted_at = COALESCE(sqlc.narg('accepted_at'), accepted_at)
WHERE id = $1
RETURNING *;

-- name: DeleteReview :one
DELETE FROM review
WHERE id = $1
RETURNING *;

-- name: GetReviewsByDish :many
SELECT review.*
FROM review
JOIN occurrence ON (review.occurrence = occurrence.id)
JOIN dish ON (occurrence.dish = dish.id)
WHERE dish.id = $1;