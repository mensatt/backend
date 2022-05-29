-- name: GetAllTags :many
SELECT *
FROM tag;

-- name: GetTagByKey :one
SELECT *
FROM tag
WHERE key = $1;

-- name: GetAllDishes :many
SELECT *
FROM dish;

-- name: GetDishByID :one
SELECT *
FROM dish
WHERE id = $1;

-- name: GetAllAliases :many
SELECT *
FROM dish_alias;

-- name: GetAliasesForDish :many
SELECT alias_name
FROM dish_alias
WHERE dish = $1;

-- name: GetAllReviews :many
SELECT *
FROM review;

-- name: GetReviewByID :one
SELECT *
FROM review
WHERE id = $1;

-- name: GetAllImages :many
SELECT *
FROM image;

-- name: GetImageByID :one
SELECT *
FROM image
WHERE id = $1;

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

-- name: CreateTag :one
INSERT INTO tag (key, name, description, short_name, priority, is_allergy)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;

-- name: CreateDish :one
INSERT INTO dish (name)
VALUES ($1)
RETURNING *;

-- name: CreateDishAlias :one
INSERT INTO dish_alias (alias_name, dish)
VALUES ($1, $2)
RETURNING *;

-- name: CreateOccurrence :one
INSERT INTO occurrence (dish, date, review_status, kj, kcal, fat, saturated_fat, carbohydrates, sugar, fiber, protein, salt, price_student, price_staff, price_guest)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15)
RETURNING *;

-- name: RemoveOccurrenceSideDish :one
DELETE FROM occurrence_side_dishes
WHERE occurrence = $1 AND dish = $2
RETURNING *;

-- name: AddOccurrenceSideDish :one
INSERT INTO occurrence_side_dishes (occurrence, dish)
VALUES ($1, $2)
RETURNING *;

-- name: AddMultipleOccurrenceSideDishes :copyfrom
INSERT INTO occurrence_side_dishes (occurrence, dish)
VALUES ($1, $2);

-- name: RemoveOccurrenceTag :one
DELETE FROM occurrence_tag
WHERE occurrence = $1 AND tag = $2
RETURNING *;

-- name: AddOccurrenceTag :one
INSERT INTO occurrence_tag (occurrence, tag)
VALUES ($1, $2)
RETURNING *;

-- name: AddMultipleOccurrenceTags :copyfrom
INSERT INTO occurrence_tag (occurrence, tag)
VALUES ($1, $2);

-- name: CreateReview :one
INSERT INTO review (occurrence, display_name, stars, text)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: DeleteOccurrence :one
DELETE FROM occurrence
WHERE id = $1
RETURNING *;

-- name: EditOccurrence :one
UPDATE occurrence
SET dish = $1, date = $2, review_status = $3, kj = $4, kcal = $5, fat = $6, saturated_fat = $7, carbohydrates = $8, sugar = $9, fiber = $10, protein = $11, salt = $12, price_student = $13, price_staff = $14, price_guest = $15
WHERE id = $16
RETURNING *;
