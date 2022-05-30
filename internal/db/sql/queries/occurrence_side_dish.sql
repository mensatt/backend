-- name: AddOccurrenceSideDish :one
INSERT INTO occurrence_side_dishes (occurrence, dish)
VALUES ($1, $2)
RETURNING *;

-- name: AddMultipleOccurrenceSideDishes :copyfrom
INSERT INTO occurrence_side_dishes (occurrence, dish)
VALUES ($1, $2);

-- name: RemoveOccurrenceSideDish :one
DELETE FROM occurrence_side_dishes
WHERE occurrence = $1 AND dish = $2
RETURNING *;