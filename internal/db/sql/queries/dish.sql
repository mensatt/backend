-- name: GetAllDishes :many
SELECT *
FROM dish;

-- name: GetDishByID :one
SELECT *
FROM dish
WHERE id = $1;

-- name: GetDishReviewMetadata :one
SELECT AVG(review.stars) AS average_stars, COUNT(*) AS review_count
FROM review
JOIN occurrence ON (review.occurrence = occurrence.id)
JOIN dish ON (occurrence.dish = dish.id)
WHERE dish.id = $1 AND review.accepted_at IS NOT NULL;

-- name: CreateDish :one
INSERT INTO dish (name_de, name_en)
VALUES ($1, $2)
RETURNING *;

-- name: UpdateDish :one
UPDATE dish
SET
    name_de = COALESCE(sqlc.narg('name_de'), name_de),
    name_en = COALESCE(sqlc.narg('name_en'), name_en)
WHERE id = $1
RETURNING *;