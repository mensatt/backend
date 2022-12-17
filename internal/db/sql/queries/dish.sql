-- name: GetAllDishes :many
SELECT *
FROM dish;

-- name: GetDishByID :one
SELECT *
FROM dish
WHERE id = $1;

-- name: GetSideDishesForOccurrence :many
SELECT dish.*
FROM occurrence_side_dishes JOIN dish ON (occurrence_side_dishes.dish = dish.id)
WHERE occurrence_side_dishes.occurrence = $1;

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

-- name: DeleteDish :one
DELETE FROM dish
WHERE id = $1
RETURNING *;