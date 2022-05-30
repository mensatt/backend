-- name: GetAllDishes :many
SELECT *
FROM dish;

-- name: GetDishByID :one
SELECT *
FROM dish
WHERE id = $1;

-- name: CreateDish :one
INSERT INTO dish (name)
VALUES ($1)
RETURNING *;

-- name: RenameDish :one
UPDATE dish
SET name = $1
WHERE id = $2
RETURNING *;
