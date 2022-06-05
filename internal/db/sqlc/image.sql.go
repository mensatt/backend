// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.13.0
// source: image.sql

package sqlc

import (
	"context"

	"github.com/google/uuid"
)

const getAllImages = `-- name: GetAllImages :many
SELECT id, occurrence, display_name, description, up_votes, down_votes, created_at, updated_at, accepted_at
FROM image
`

func (q *Queries) GetAllImages(ctx context.Context) ([]*Image, error) {
	rows, err := q.db.Query(ctx, getAllImages)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []*Image
	for rows.Next() {
		var i Image
		if err := rows.Scan(
			&i.ID,
			&i.Occurrence,
			&i.DisplayName,
			&i.Description,
			&i.UpVotes,
			&i.DownVotes,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.AcceptedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, &i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getImageByID = `-- name: GetImageByID :one
SELECT id, occurrence, display_name, description, up_votes, down_votes, created_at, updated_at, accepted_at
FROM image
WHERE id = $1
`

func (q *Queries) GetImageByID(ctx context.Context, id uuid.UUID) (*Image, error) {
	row := q.db.QueryRow(ctx, getImageByID, id)
	var i Image
	err := row.Scan(
		&i.ID,
		&i.Occurrence,
		&i.DisplayName,
		&i.Description,
		&i.UpVotes,
		&i.DownVotes,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.AcceptedAt,
	)
	return &i, err
}

const getImagesByDish = `-- name: GetImagesByDish :many
SELECT image.id, image.occurrence, image.display_name, image.description, image.up_votes, image.down_votes, image.created_at, image.updated_at, image.accepted_at
FROM image
JOIN occurrence ON (image.occurrence = occurrence.id)
JOIN dish ON (occurrence.dish = dish.id)
WHERE dish.id = $1
`

func (q *Queries) GetImagesByDish(ctx context.Context, id uuid.UUID) ([]*Image, error) {
	rows, err := q.db.Query(ctx, getImagesByDish, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []*Image
	for rows.Next() {
		var i Image
		if err := rows.Scan(
			&i.ID,
			&i.Occurrence,
			&i.DisplayName,
			&i.Description,
			&i.UpVotes,
			&i.DownVotes,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.AcceptedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, &i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
