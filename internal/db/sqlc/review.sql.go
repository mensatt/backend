// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.13.0
// source: review.sql

package sqlc

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
)

const createReview = `-- name: CreateReview :one
INSERT INTO review (occurrence, display_name, stars, text)
VALUES ($1, $2, $3, $4)
RETURNING id, occurrence, display_name, stars, text, up_votes, down_votes, created_at, updated_at, accepted_at
`

type CreateReviewParams struct {
	Occurrence  uuid.UUID      `json:"occurrence"`
	DisplayName sql.NullString `json:"display_name"`
	Stars       int32          `json:"stars"`
	Text        sql.NullString `json:"text"`
}

func (q *Queries) CreateReview(ctx context.Context, arg *CreateReviewParams) (*Review, error) {
	row := q.db.QueryRow(ctx, createReview,
		arg.Occurrence,
		arg.DisplayName,
		arg.Stars,
		arg.Text,
	)
	var i Review
	err := row.Scan(
		&i.ID,
		&i.Occurrence,
		&i.DisplayName,
		&i.Stars,
		&i.Text,
		&i.UpVotes,
		&i.DownVotes,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.AcceptedAt,
	)
	return &i, err
}

const deleteReview = `-- name: DeleteReview :one
DELETE FROM review
WHERE id = $1
RETURNING id, occurrence, display_name, stars, text, up_votes, down_votes, created_at, updated_at, accepted_at
`

func (q *Queries) DeleteReview(ctx context.Context, id uuid.UUID) (*Review, error) {
	row := q.db.QueryRow(ctx, deleteReview, id)
	var i Review
	err := row.Scan(
		&i.ID,
		&i.Occurrence,
		&i.DisplayName,
		&i.Stars,
		&i.Text,
		&i.UpVotes,
		&i.DownVotes,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.AcceptedAt,
	)
	return &i, err
}

const editReview = `-- name: EditReview :one
UPDATE review
SET occurrence = $1, display_name = $2, stars = $3, text = $4, updated_at = NOW(), accepted_at = $5
WHERE id = $6
RETURNING id, occurrence, display_name, stars, text, up_votes, down_votes, created_at, updated_at, accepted_at
`

type EditReviewParams struct {
	Occurrence  uuid.UUID      `json:"occurrence"`
	DisplayName sql.NullString `json:"display_name"`
	Stars       int32          `json:"stars"`
	Text        sql.NullString `json:"text"`
	AcceptedAt  sql.NullTime   `json:"accepted_at"`
	ID          uuid.UUID      `json:"id"`
}

func (q *Queries) EditReview(ctx context.Context, arg *EditReviewParams) (*Review, error) {
	row := q.db.QueryRow(ctx, editReview,
		arg.Occurrence,
		arg.DisplayName,
		arg.Stars,
		arg.Text,
		arg.AcceptedAt,
		arg.ID,
	)
	var i Review
	err := row.Scan(
		&i.ID,
		&i.Occurrence,
		&i.DisplayName,
		&i.Stars,
		&i.Text,
		&i.UpVotes,
		&i.DownVotes,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.AcceptedAt,
	)
	return &i, err
}

const getAllReviews = `-- name: GetAllReviews :many
SELECT id, occurrence, display_name, stars, text, up_votes, down_votes, created_at, updated_at, accepted_at
FROM review
`

func (q *Queries) GetAllReviews(ctx context.Context) ([]*Review, error) {
	rows, err := q.db.Query(ctx, getAllReviews)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []*Review
	for rows.Next() {
		var i Review
		if err := rows.Scan(
			&i.ID,
			&i.Occurrence,
			&i.DisplayName,
			&i.Stars,
			&i.Text,
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

const getReviewByID = `-- name: GetReviewByID :one
SELECT id, occurrence, display_name, stars, text, up_votes, down_votes, created_at, updated_at, accepted_at
FROM review
WHERE id = $1
`

func (q *Queries) GetReviewByID(ctx context.Context, id uuid.UUID) (*Review, error) {
	row := q.db.QueryRow(ctx, getReviewByID, id)
	var i Review
	err := row.Scan(
		&i.ID,
		&i.Occurrence,
		&i.DisplayName,
		&i.Stars,
		&i.Text,
		&i.UpVotes,
		&i.DownVotes,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.AcceptedAt,
	)
	return &i, err
}
