-- name: GetAllAliases :many
SELECT *
FROM dish_alias;

-- name: GetAliasesForDish :many
SELECT alias_name
FROM dish_alias
WHERE dish = $1;

-- name: CreateDishAlias :one
INSERT INTO dish_alias (alias_name, dish)
VALUES ($1, $2)
RETURNING *;

-- name: UpdateDishAlias :one
UPDATE dish_alias
SET alias_name = sqlc.arg(old_alias_name)
WHERE alias_name = $1 AND dish = $2
RETURNING *;

-- name: DeleteDishAlias :one
DELETE FROM dish_alias
WHERE alias_name = $1 AND dish = $2
RETURNING *;