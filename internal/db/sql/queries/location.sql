-- name: GetAllLocations :many
SELECT *
FROM location;

-- name: GetLocationByID :one
SELECT *
FROM location
WHERE id = $1;