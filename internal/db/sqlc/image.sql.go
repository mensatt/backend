// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.14.0
// source: image.sql

package sqlc

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
)

const createImage = `-- name: CreateImage :one
INSERT INTO image (occurrence, display_name, description)
VALUES ($1, $2, $3)
RETURNING id, occurrence, display_name, description, up_votes, down_votes, created_at, updated_at, accepted_at
`

type CreateImageParams struct {
	Occurrence  uuid.UUID      `json:"occurrence"`
	DisplayName string         `json:"display_name"`
	Description sql.NullString `json:"description"`
}

func (q *Queries) CreateImage(ctx context.Context, arg *CreateImageParams) (*Image, error) {
	row := q.db.QueryRow(ctx, createImage, arg.Occurrence, arg.DisplayName, arg.Description)
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

const deleteImage = `-- name: DeleteImage :one
DELETE FROM image
WHERE id = $1
RETURNING id, occurrence, display_name, description, up_votes, down_votes, created_at, updated_at, accepted_at
`

func (q *Queries) DeleteImage(ctx context.Context, id uuid.UUID) (*Image, error) {
	row := q.db.QueryRow(ctx, deleteImage, id)
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

const updateImage = `-- name: UpdateImage :one
UPDATE image
SET 
    occurrence = COALESCE($2, occurrence),
    display_name = COALESCE($3, display_name),
    description = COALESCE($4, description),
    updated_at = NOW(),
    accepted_at = COALESCE($5, accepted_at)
WHERE id = $1
RETURNING id, occurrence, display_name, description, up_votes, down_votes, created_at, updated_at, accepted_at
`

type UpdateImageParams struct {
	ID          uuid.UUID      `json:"id"`
	Occurrence  uuid.NullUUID  `json:"occurrence"`
	DisplayName sql.NullString `json:"display_name"`
	Description sql.NullString `json:"description"`
	AcceptedAt  sql.NullTime   `json:"accepted_at"`
}

func (q *Queries) UpdateImage(ctx context.Context, arg *UpdateImageParams) (*Image, error) {
	row := q.db.QueryRow(ctx, updateImage,
		arg.ID,
		arg.Occurrence,
		arg.DisplayName,
		arg.Description,
		arg.AcceptedAt,
	)
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
