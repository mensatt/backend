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
    accepted_at = CASE 
        WHEN sqlc.narg('approved')::bool = TRUE THEN  NOW()
        WHEN sqlc.narg('approved')::bool = FALSE THEN NULL
        ELSE accepted_at
    END
WHERE id = $1
RETURNING *;

-- name: SetReviewApproval :one
UPDATE review
SET accepted_at = $1
WHERE id = $2
RETURNING *;

-- name: DeleteReview :one
DELETE FROM review
WHERE id = $1
RETURNING *;


-- name: GetReviewByImage :one
SELECT review.*
FROM review
JOIN image ON (image.review = review.id)
WHERE image.id = $1;


-- name: GetDishReviews :many
SELECT review.*
FROM review
JOIN occurrence ON (review.occurrence = occurrence.id)
JOIN dish ON (occurrence.dish = dish.id)
WHERE dish.id = $1
    AND (CASE 
            WHEN sqlc.narg('approved')::bool = TRUE THEN review.accepted_at IS NOT NULL
            WHEN sqlc.narg('approved')::bool = FALSE THEN review.accepted_at IS NULL
            ELSE TRUE 
        END);

-- name: GetDishReviewMetadata :one
SELECT AVG(review.stars) AS average_stars, COUNT(*) AS review_count
FROM review
JOIN occurrence ON (review.occurrence = occurrence.id)
JOIN dish ON (occurrence.dish = dish.id)
WHERE dish.id = $1
    AND (CASE 
            WHEN sqlc.narg('approved')::bool = TRUE THEN review.accepted_at IS NOT NULL
            WHEN sqlc.narg('approved')::bool = FALSE THEN review.accepted_at IS NULL
            ELSE TRUE 
        END);

-- name: GetOccurrenceReviews :many
SELECT review.*
FROM occurrence JOIN review ON (occurrence.id = review.occurrence)
WHERE occurrence.id = $1
    AND (CASE 
            WHEN sqlc.narg('approved')::bool = TRUE THEN review.accepted_at IS NOT NULL
            WHEN sqlc.narg('approved')::bool = FALSE THEN review.accepted_at IS NULL
            ELSE TRUE 
        END);

-- name: GetOccurrenceReviewMetadata :one
SELECT AVG(review.stars) AS average_stars, COUNT(*) AS review_count
FROM review
WHERE review.occurrence = $1
    AND (CASE 
            WHEN sqlc.narg('approved')::bool = TRUE THEN review.accepted_at IS NOT NULL
            WHEN sqlc.narg('approved')::bool = FALSE THEN review.accepted_at IS NULL
            ELSE TRUE 
        END);