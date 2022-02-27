-- name: GetAllAllergies :many
SELECT *
FROM allergy;

-- name: GetAllTags :many
SELECT *
FROM tag;

-- name: GetOccurenceByID :one
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

-- name: GetAllergiesForOccurrence :many
SELECT allergy.*
FROM occurrence_allergy JOIN allergy ON occurrence_allergy.allergy_abbreviation = allergy.abbreviation
WHERE occurrence_allergy.occurrence_id = $1; 

-- name: GetTagsForOccurrence :many
SELECT tag.*
FROM occurrence_tag JOIN tag ON occurrence_tag.tag_abbreviation = tag.abbreviation
WHERE occurrence_tag.occurrence_id = $1; 
