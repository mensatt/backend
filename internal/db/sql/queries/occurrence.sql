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

-- name: GetOccurrencesAfterInclusiveDate :many
SELECT *
FROM occurrence
WHERE date >= $1;

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

-- name: GetOccurrenceReviewMetadata :one
SELECT AVG(CAST(review.stars as Float)) AS average_stars, COUNT(*) AS review_count
FROM review
WHERE review.occurrence = $1;

-- name: CreateOccurrence :one
INSERT INTO occurrence (location, dish, date, review_status, kj, kcal, fat, saturated_fat, carbohydrates, sugar, fiber, protein, salt, price_student, price_staff, price_guest)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16)
RETURNING *;

-- name: UpdateOccurrence :one
UPDATE occurrence
SET 
    dish = COALESCE(sqlc.narg('dish'), dish),
    date = COALESCE(sqlc.narg('date'), date),
    review_status = COALESCE(sqlc.narg('review_status'), review_status),
    kj = COALESCE(sqlc.narg('kj'), kj),
    kcal = COALESCE(sqlc.narg('kcal'), kcal),
    fat = COALESCE(sqlc.narg('fat'), fat),
    saturated_fat = COALESCE(sqlc.narg('saturated_fat'), saturated_fat),
    carbohydrates = COALESCE(sqlc.narg('carbohydrates'), carbohydrates),
    sugar = COALESCE(sqlc.narg('sugar'), sugar),
    fiber = COALESCE(sqlc.narg('fiber'), fiber),
    protein = COALESCE(sqlc.narg('protein'), protein),
    salt = COALESCE(sqlc.narg('salt'), salt),
    price_student = COALESCE(sqlc.narg('price_student'), price_student),
    price_staff = COALESCE(sqlc.narg('price_staff'), price_staff),
    price_guest = COALESCE(sqlc.narg('price_guest'), price_guest)
WHERE id = $1
RETURNING *;

-- name: DeleteOccurrence :one
DELETE FROM occurrence
WHERE id = $1
RETURNING *;