// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.13.0
// source: location.sql

package sqlc

import (
	"context"

	"github.com/google/uuid"
)

const getAllLocations = `-- name: GetAllLocations :many
SELECT id, location_id, name
FROM location
`

func (q *Queries) GetAllLocations(ctx context.Context) ([]*Location, error) {
	rows, err := q.db.Query(ctx, getAllLocations)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []*Location
	for rows.Next() {
		var i Location
		if err := rows.Scan(&i.ID, &i.LocationID, &i.Name); err != nil {
			return nil, err
		}
		items = append(items, &i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getLocationByID = `-- name: GetLocationByID :one
SELECT id, location_id, name
FROM location
WHERE id = $1
`

func (q *Queries) GetLocationByID(ctx context.Context, id uuid.UUID) (*Location, error) {
	row := q.db.QueryRow(ctx, getLocationByID, id)
	var i Location
	err := row.Scan(&i.ID, &i.LocationID, &i.Name)
	return &i, err
}
