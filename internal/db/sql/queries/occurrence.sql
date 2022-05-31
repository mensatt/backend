-- name: GetAllOccurrences :many
SELECT *
FROM occurrence;

-- name: GetOccurrenceByID :one
SELECT * 
FROM occurrence
WHERE id = $1;

-- name: GetOccurrencesByDate :many
SELECT *
FROM occurrence
WHERE date = $1;

-- name: GetSideDishesForOccurrence :many
SELECT dish.*
FROM occurrence_side_dishes JOIN dish ON occurrence_side_dishes.dish = dish.id
WHERE occurrence_side_dishes.occurrence = $1;

-- name: GetTagsForOccurrence :many
SELECT tag.*
FROM occurrence_tag JOIN tag ON occurrence_tag.tag = tag.key
WHERE occurrence_tag.occurrence = $1;

-- name: GetReviewsForOccurrence :many
SELECT review.*
FROM occurrence JOIN review ON occurrence.id = review.occurrence
WHERE occurrence.id = $1;

-- name: GetImagesForOccurrence :many
SELECT image.*
FROM occurrence JOIN image ON occurrence.id = image.occurrence
WHERE occurrence.id = $1;

-- name: CreateOccurrence :one
INSERT INTO occurrence (dish, date, review_status, kj, kcal, fat, saturated_fat, carbohydrates, sugar, fiber, protein, salt, price_student, price_staff, price_guest)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15)
RETURNING *;

-- name: EditOccurrence :one
UPDATE occurrence
SET dish = $1, date = $2, review_status = $3, kj = $4, kcal = $5, fat = $6, saturated_fat = $7, carbohydrates = $8, sugar = $9, fiber = $10, protein = $11, salt = $12, price_student = $13, price_staff = $14, price_guest = $15
WHERE id = $16
RETURNING *;

-- name: DeleteOccurrence :one
DELETE FROM occurrence
WHERE id = $1
RETURNING *;