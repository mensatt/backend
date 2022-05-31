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