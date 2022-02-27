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
FROM occurrence_tag JOIN tag ON occurrence_tag.tag_abbreviation = tag.abbreviation
WHERE occurrence_tag.occurrence_id = $1; 
