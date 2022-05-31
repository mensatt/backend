-- name: GetAllAliases :many
SELECT *
FROM dish_alias;

-- name: GetAliasesForDish :many
SELECT alias_name
FROM dish_alias
WHERE dish = $1;

-- name: CreateDishAlias :one
INSERT INTO dish_alias (alias_name, normalized_alias_name, dish)
VALUES ($1, $2, $3)
RETURNING *;

-- name: UpdateDishAlias :one
UPDATE dish_alias
SET alias_name = sqlc.arg(new_alias_name), normalized_alias_name = sqlc.arg(new_normalized_alias_name)
WHERE alias_name = $1
RETURNING *;

-- name: DeleteDishAlias :one
DELETE FROM dish_alias
WHERE alias_name = $1
RETURNING *;