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

-- name: UpdateDish :one
UPDATE dish
SET
    name = COALESCE(sqlc.narg('name'), name)
WHERE id = $1
RETURNING *;