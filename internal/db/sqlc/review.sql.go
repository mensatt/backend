// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0
// source: review.sql

package sqlc

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/jackc/pgtype"
)

const createReview = `-- name: CreateReview :one
INSERT INTO review (occurrence, display_name, stars, text)
VALUES ($1, $2, $3, $4)
RETURNING id, occurrence, display_name, stars, text, created_at, updated_at, accepted_at
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
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.AcceptedAt,
	)
	return &i, err
}

const deleteReview = `-- name: DeleteReview :one
DELETE FROM review
WHERE id = $1
RETURNING id, occurrence, display_name, stars, text, created_at, updated_at, accepted_at
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
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.AcceptedAt,
	)
	return &i, err
}

const getAllReviews = `-- name: GetAllReviews :many
SELECT id, occurrence, display_name, stars, text, created_at, updated_at, accepted_at
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

const getDishReviewMetadata = `-- name: GetDishReviewMetadata :one
SELECT AVG(review.stars) AS average_stars, COUNT(*) AS review_count
FROM review
JOIN occurrence ON (review.occurrence = occurrence.id)
JOIN dish ON (occurrence.dish = dish.id)
WHERE dish.id = $1
    AND (CASE 
            WHEN $2::bool = TRUE THEN review.accepted_at IS NOT NULL
            WHEN $2::bool = FALSE THEN review.accepted_at IS NULL
            ELSE TRUE 
        END)
`

type GetDishReviewMetadataParams struct {
	ID       uuid.UUID    `json:"id"`
	Approved sql.NullBool `json:"approved"`
}

type GetDishReviewMetadataRow struct {
	AverageStars pgtype.Numeric `json:"average_stars"`
	ReviewCount  int64          `json:"review_count"`
}

func (q *Queries) GetDishReviewMetadata(ctx context.Context, arg *GetDishReviewMetadataParams) (*GetDishReviewMetadataRow, error) {
	row := q.db.QueryRow(ctx, getDishReviewMetadata, arg.ID, arg.Approved)
	var i GetDishReviewMetadataRow
	err := row.Scan(&i.AverageStars, &i.ReviewCount)
	return &i, err
}

const getDishReviews = `-- name: GetDishReviews :many
SELECT review.id, review.occurrence, review.display_name, review.stars, review.text, review.created_at, review.updated_at, review.accepted_at
FROM review
JOIN occurrence ON (review.occurrence = occurrence.id)
JOIN dish ON (occurrence.dish = dish.id)
WHERE dish.id = $1
    AND (CASE 
            WHEN $2::bool = TRUE THEN review.accepted_at IS NOT NULL
            WHEN $2::bool = FALSE THEN review.accepted_at IS NULL
            ELSE TRUE 
        END)
`

type GetDishReviewsParams struct {
	ID       uuid.UUID    `json:"id"`
	Approved sql.NullBool `json:"approved"`
}

func (q *Queries) GetDishReviews(ctx context.Context, arg *GetDishReviewsParams) ([]*Review, error) {
	rows, err := q.db.Query(ctx, getDishReviews, arg.ID, arg.Approved)
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

const getOccurrenceReviewMetadata = `-- name: GetOccurrenceReviewMetadata :one
SELECT AVG(review.stars) AS average_stars, COUNT(*) AS review_count
FROM review
WHERE review.occurrence = $1
    AND (CASE 
            WHEN $2::bool = TRUE THEN review.accepted_at IS NOT NULL
            WHEN $2::bool = FALSE THEN review.accepted_at IS NULL
            ELSE TRUE 
        END)
`

type GetOccurrenceReviewMetadataParams struct {
	Occurrence uuid.UUID    `json:"occurrence"`
	Approved   sql.NullBool `json:"approved"`
}

type GetOccurrenceReviewMetadataRow struct {
	AverageStars pgtype.Numeric `json:"average_stars"`
	ReviewCount  int64          `json:"review_count"`
}

func (q *Queries) GetOccurrenceReviewMetadata(ctx context.Context, arg *GetOccurrenceReviewMetadataParams) (*GetOccurrenceReviewMetadataRow, error) {
	row := q.db.QueryRow(ctx, getOccurrenceReviewMetadata, arg.Occurrence, arg.Approved)
	var i GetOccurrenceReviewMetadataRow
	err := row.Scan(&i.AverageStars, &i.ReviewCount)
	return &i, err
}

const getOccurrenceReviews = `-- name: GetOccurrenceReviews :many
SELECT review.id, review.occurrence, review.display_name, review.stars, review.text, review.created_at, review.updated_at, review.accepted_at
FROM occurrence JOIN review ON (occurrence.id = review.occurrence)
WHERE occurrence.id = $1
    AND (CASE 
            WHEN $2::bool = TRUE THEN review.accepted_at IS NOT NULL
            WHEN $2::bool = FALSE THEN review.accepted_at IS NULL
            ELSE TRUE 
        END)
`

type GetOccurrenceReviewsParams struct {
	ID       uuid.UUID    `json:"id"`
	Approved sql.NullBool `json:"approved"`
}

func (q *Queries) GetOccurrenceReviews(ctx context.Context, arg *GetOccurrenceReviewsParams) ([]*Review, error) {
	rows, err := q.db.Query(ctx, getOccurrenceReviews, arg.ID, arg.Approved)
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
SELECT id, occurrence, display_name, stars, text, created_at, updated_at, accepted_at
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
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.AcceptedAt,
	)
	return &i, err
}

const getReviewByImage = `-- name: GetReviewByImage :one
SELECT review.id, review.occurrence, review.display_name, review.stars, review.text, review.created_at, review.updated_at, review.accepted_at
FROM review
JOIN image ON (image.review = review.id)
WHERE image.id = $1
`

func (q *Queries) GetReviewByImage(ctx context.Context, id uuid.UUID) (*Review, error) {
	row := q.db.QueryRow(ctx, getReviewByImage, id)
	var i Review
	err := row.Scan(
		&i.ID,
		&i.Occurrence,
		&i.DisplayName,
		&i.Stars,
		&i.Text,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.AcceptedAt,
	)
	return &i, err
}

const setReviewApproval = `-- name: SetReviewApproval :one
UPDATE review
SET accepted_at = $1
WHERE id = $2
RETURNING id, occurrence, display_name, stars, text, created_at, updated_at, accepted_at
`

type SetReviewApprovalParams struct {
	AcceptedAt sql.NullTime `json:"accepted_at"`
	ID         uuid.UUID    `json:"id"`
}

func (q *Queries) SetReviewApproval(ctx context.Context, arg *SetReviewApprovalParams) (*Review, error) {
	row := q.db.QueryRow(ctx, setReviewApproval, arg.AcceptedAt, arg.ID)
	var i Review
	err := row.Scan(
		&i.ID,
		&i.Occurrence,
		&i.DisplayName,
		&i.Stars,
		&i.Text,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.AcceptedAt,
	)
	return &i, err
}

const updateReview = `-- name: UpdateReview :one
UPDATE review
SET 
    occurrence = COALESCE($2, occurrence),
    display_name = COALESCE($3, display_name),
    stars = COALESCE($4, stars),
    text = COALESCE($5, text),
    updated_at = NOW(),
    accepted_at = CASE 
        WHEN $6::bool = TRUE THEN  NOW()
        WHEN $6::bool = FALSE THEN NULL
        ELSE accepted_at
    END
WHERE id = $1
RETURNING id, occurrence, display_name, stars, text, created_at, updated_at, accepted_at
`

type UpdateReviewParams struct {
	ID          uuid.UUID      `json:"id"`
	Occurrence  uuid.NullUUID  `json:"occurrence"`
	DisplayName sql.NullString `json:"display_name"`
	Stars       sql.NullInt32  `json:"stars"`
	Text        sql.NullString `json:"text"`
	Approved    sql.NullBool   `json:"approved"`
}

func (q *Queries) UpdateReview(ctx context.Context, arg *UpdateReviewParams) (*Review, error) {
	row := q.db.QueryRow(ctx, updateReview,
		arg.ID,
		arg.Occurrence,
		arg.DisplayName,
		arg.Stars,
		arg.Text,
		arg.Approved,
	)
	var i Review
	err := row.Scan(
		&i.ID,
		&i.Occurrence,
		&i.DisplayName,
		&i.Stars,
		&i.Text,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.AcceptedAt,
	)
	return &i, err
}
