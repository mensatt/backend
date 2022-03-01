-- name: GetAllTags :many
SELECT *
FROM tag;

-- name: GetAllDishes :many
SELECT *
FROM dish;

-- name: GetAllReviews :many
SELECT *
FROM review;

-- name: GetAllImages :many
SELECT *
FROM image;

-- name: GetOccurrenceByID :one
SELECT * 
FROM occurrence
WHERE id = $1;

-- name: GetDishByID :one
SELECT * 
FROM dish
WHERE id = $1;

-- name: GetSideDishesForOccurrence :many
SELECT dish.*
FROM occurrence_side_dishes JOIN dish ON occurrence_side_dishes.dish_id = dish.id
WHERE occurrence_side_dishes.occurrence_id = $1;

-- name: GetTagsForOccurrence :many
SELECT tag.*
FROM occurrence_tag JOIN tag ON occurrence_tag.tag_key = tag.key
WHERE occurrence_tag.occurrence_id = $1; 

-- name: GetOccurrencesByDate :many
SELECT *
FROM occurrence
WHERE date = $1;

-- name: GetImagesForOccurrence :many
SELECT *
FROM image
WHERE occurrence = $1;

-- name: CreateTag :one
INSERT INTO tag (key, name, description, short_name, priority, is_allergy)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;
