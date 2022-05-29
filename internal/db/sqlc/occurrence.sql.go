// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.13.0
// source: occurrence.sql

package sqlc

import (
	"context"
	"database/sql"
	"time"

	"github.com/google/uuid"
)

const createOccurrence = `-- name: CreateOccurrence :one
INSERT INTO occurrence (dish, date, review_status, kj, kcal, fat, saturated_fat, carbohydrates, sugar, fiber, protein, salt, price_student, price_staff, price_guest)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15)
RETURNING id, dish, date, review_status, kj, kcal, fat, saturated_fat, carbohydrates, sugar, fiber, protein, salt, price_student, price_staff, price_guest
`

type CreateOccurrenceParams struct {
	Dish          uuid.UUID     `json:"dish"`
	Date          time.Time     `json:"date"`
	ReviewStatus  ReviewStatus  `json:"review_status"`
	Kj            sql.NullInt32 `json:"kj"`
	Kcal          sql.NullInt32 `json:"kcal"`
	Fat           sql.NullInt32 `json:"fat"`
	SaturatedFat  sql.NullInt32 `json:"saturated_fat"`
	Carbohydrates sql.NullInt32 `json:"carbohydrates"`
	Sugar         sql.NullInt32 `json:"sugar"`
	Fiber         sql.NullInt32 `json:"fiber"`
	Protein       sql.NullInt32 `json:"protein"`
	Salt          sql.NullInt32 `json:"salt"`
	PriceStudent  sql.NullInt32 `json:"price_student"`
	PriceStaff    sql.NullInt32 `json:"price_staff"`
	PriceGuest    sql.NullInt32 `json:"price_guest"`
}

func (q *Queries) CreateOccurrence(ctx context.Context, arg *CreateOccurrenceParams) (*Occurrence, error) {
	row := q.db.QueryRow(ctx, createOccurrence,
		arg.Dish,
		arg.Date,
		arg.ReviewStatus,
		arg.Kj,
		arg.Kcal,
		arg.Fat,
		arg.SaturatedFat,
		arg.Carbohydrates,
		arg.Sugar,
		arg.Fiber,
		arg.Protein,
		arg.Salt,
		arg.PriceStudent,
		arg.PriceStaff,
		arg.PriceGuest,
	)
	var i Occurrence
	err := row.Scan(
		&i.ID,
		&i.Dish,
		&i.Date,
		&i.ReviewStatus,
		&i.Kj,
		&i.Kcal,
		&i.Fat,
		&i.SaturatedFat,
		&i.Carbohydrates,
		&i.Sugar,
		&i.Fiber,
		&i.Protein,
		&i.Salt,
		&i.PriceStudent,
		&i.PriceStaff,
		&i.PriceGuest,
	)
	return &i, err
}

const deleteOccurrence = `-- name: DeleteOccurrence :one
DELETE FROM occurrence
WHERE id = $1
RETURNING id, dish, date, review_status, kj, kcal, fat, saturated_fat, carbohydrates, sugar, fiber, protein, salt, price_student, price_staff, price_guest
`

func (q *Queries) DeleteOccurrence(ctx context.Context, id uuid.UUID) (*Occurrence, error) {
	row := q.db.QueryRow(ctx, deleteOccurrence, id)
	var i Occurrence
	err := row.Scan(
		&i.ID,
		&i.Dish,
		&i.Date,
		&i.ReviewStatus,
		&i.Kj,
		&i.Kcal,
		&i.Fat,
		&i.SaturatedFat,
		&i.Carbohydrates,
		&i.Sugar,
		&i.Fiber,
		&i.Protein,
		&i.Salt,
		&i.PriceStudent,
		&i.PriceStaff,
		&i.PriceGuest,
	)
	return &i, err
}

const editOccurrence = `-- name: EditOccurrence :one
UPDATE occurrence
SET dish = $1, date = $2, review_status = $3, kj = $4, kcal = $5, fat = $6, saturated_fat = $7, carbohydrates = $8, sugar = $9, fiber = $10, protein = $11, salt = $12, price_student = $13, price_staff = $14, price_guest = $15
WHERE id = $16
RETURNING id, dish, date, review_status, kj, kcal, fat, saturated_fat, carbohydrates, sugar, fiber, protein, salt, price_student, price_staff, price_guest
`

type EditOccurrenceParams struct {
	Dish          uuid.UUID     `json:"dish"`
	Date          time.Time     `json:"date"`
	ReviewStatus  ReviewStatus  `json:"review_status"`
	Kj            sql.NullInt32 `json:"kj"`
	Kcal          sql.NullInt32 `json:"kcal"`
	Fat           sql.NullInt32 `json:"fat"`
	SaturatedFat  sql.NullInt32 `json:"saturated_fat"`
	Carbohydrates sql.NullInt32 `json:"carbohydrates"`
	Sugar         sql.NullInt32 `json:"sugar"`
	Fiber         sql.NullInt32 `json:"fiber"`
	Protein       sql.NullInt32 `json:"protein"`
	Salt          sql.NullInt32 `json:"salt"`
	PriceStudent  sql.NullInt32 `json:"price_student"`
	PriceStaff    sql.NullInt32 `json:"price_staff"`
	PriceGuest    sql.NullInt32 `json:"price_guest"`
	ID            uuid.UUID     `json:"id"`
}

func (q *Queries) EditOccurrence(ctx context.Context, arg *EditOccurrenceParams) (*Occurrence, error) {
	row := q.db.QueryRow(ctx, editOccurrence,
		arg.Dish,
		arg.Date,
		arg.ReviewStatus,
		arg.Kj,
		arg.Kcal,
		arg.Fat,
		arg.SaturatedFat,
		arg.Carbohydrates,
		arg.Sugar,
		arg.Fiber,
		arg.Protein,
		arg.Salt,
		arg.PriceStudent,
		arg.PriceStaff,
		arg.PriceGuest,
		arg.ID,
	)
	var i Occurrence
	err := row.Scan(
		&i.ID,
		&i.Dish,
		&i.Date,
		&i.ReviewStatus,
		&i.Kj,
		&i.Kcal,
		&i.Fat,
		&i.SaturatedFat,
		&i.Carbohydrates,
		&i.Sugar,
		&i.Fiber,
		&i.Protein,
		&i.Salt,
		&i.PriceStudent,
		&i.PriceStaff,
		&i.PriceGuest,
	)
	return &i, err
}

const getAllOccurrences = `-- name: GetAllOccurrences :many
SELECT id, dish, date, review_status, kj, kcal, fat, saturated_fat, carbohydrates, sugar, fiber, protein, salt, price_student, price_staff, price_guest
FROM occurrence
`

func (q *Queries) GetAllOccurrences(ctx context.Context) ([]*Occurrence, error) {
	rows, err := q.db.Query(ctx, getAllOccurrences)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []*Occurrence
	for rows.Next() {
		var i Occurrence
		if err := rows.Scan(
			&i.ID,
			&i.Dish,
			&i.Date,
			&i.ReviewStatus,
			&i.Kj,
			&i.Kcal,
			&i.Fat,
			&i.SaturatedFat,
			&i.Carbohydrates,
			&i.Sugar,
			&i.Fiber,
			&i.Protein,
			&i.Salt,
			&i.PriceStudent,
			&i.PriceStaff,
			&i.PriceGuest,
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

const getImagesForOccurrence = `-- name: GetImagesForOccurrence :many
SELECT image.id, image.occurrence, image.display_name, image.description, image.up_votes, image.down_votes, image.created_at, image.updated_at, image.accepted_at
FROM occurrence JOIN image ON occurrence.id = image.occurrence
WHERE occurrence.id = $1
`

func (q *Queries) GetImagesForOccurrence(ctx context.Context, id uuid.UUID) ([]*Image, error) {
	rows, err := q.db.Query(ctx, getImagesForOccurrence, id)
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

const getOccurrenceByID = `-- name: GetOccurrenceByID :one
SELECT id, dish, date, review_status, kj, kcal, fat, saturated_fat, carbohydrates, sugar, fiber, protein, salt, price_student, price_staff, price_guest 
FROM occurrence
WHERE id = $1
`

func (q *Queries) GetOccurrenceByID(ctx context.Context, id uuid.UUID) (*Occurrence, error) {
	row := q.db.QueryRow(ctx, getOccurrenceByID, id)
	var i Occurrence
	err := row.Scan(
		&i.ID,
		&i.Dish,
		&i.Date,
		&i.ReviewStatus,
		&i.Kj,
		&i.Kcal,
		&i.Fat,
		&i.SaturatedFat,
		&i.Carbohydrates,
		&i.Sugar,
		&i.Fiber,
		&i.Protein,
		&i.Salt,
		&i.PriceStudent,
		&i.PriceStaff,
		&i.PriceGuest,
	)
	return &i, err
}

const getOccurrencesByDate = `-- name: GetOccurrencesByDate :many
SELECT id, dish, date, review_status, kj, kcal, fat, saturated_fat, carbohydrates, sugar, fiber, protein, salt, price_student, price_staff, price_guest
FROM occurrence
WHERE date = $1
`

func (q *Queries) GetOccurrencesByDate(ctx context.Context, date time.Time) ([]*Occurrence, error) {
	rows, err := q.db.Query(ctx, getOccurrencesByDate, date)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []*Occurrence
	for rows.Next() {
		var i Occurrence
		if err := rows.Scan(
			&i.ID,
			&i.Dish,
			&i.Date,
			&i.ReviewStatus,
			&i.Kj,
			&i.Kcal,
			&i.Fat,
			&i.SaturatedFat,
			&i.Carbohydrates,
			&i.Sugar,
			&i.Fiber,
			&i.Protein,
			&i.Salt,
			&i.PriceStudent,
			&i.PriceStaff,
			&i.PriceGuest,
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

const getReviewsForOccurrence = `-- name: GetReviewsForOccurrence :many
SELECT review.id, review.occurrence, review.display_name, review.stars, review.text, review.up_votes, review.down_votes, review.created_at, review.updated_at, review.accepted_at
FROM occurrence JOIN review ON occurrence.id = review.occurrence
WHERE occurrence.id = $1
`

func (q *Queries) GetReviewsForOccurrence(ctx context.Context, id uuid.UUID) ([]*Review, error) {
	rows, err := q.db.Query(ctx, getReviewsForOccurrence, id)
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

const getSideDishesForOccurrence = `-- name: GetSideDishesForOccurrence :many
SELECT dish.id, dish.name
FROM occurrence_side_dishes JOIN dish ON occurrence_side_dishes.dish = dish.id
WHERE occurrence_side_dishes.occurrence = $1
`

func (q *Queries) GetSideDishesForOccurrence(ctx context.Context, occurrence uuid.UUID) ([]*Dish, error) {
	rows, err := q.db.Query(ctx, getSideDishesForOccurrence, occurrence)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []*Dish
	for rows.Next() {
		var i Dish
		if err := rows.Scan(&i.ID, &i.Name); err != nil {
			return nil, err
		}
		items = append(items, &i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getTagsForOccurrence = `-- name: GetTagsForOccurrence :many
SELECT tag.key, tag.name, tag.description, tag.short_name, tag.priority, tag.is_allergy
FROM occurrence_tag JOIN tag ON occurrence_tag.tag = tag.key
WHERE occurrence_tag.occurrence = $1
`

func (q *Queries) GetTagsForOccurrence(ctx context.Context, occurrence uuid.UUID) ([]*Tag, error) {
	rows, err := q.db.Query(ctx, getTagsForOccurrence, occurrence)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []*Tag
	for rows.Next() {
		var i Tag
		if err := rows.Scan(
			&i.Key,
			&i.Name,
			&i.Description,
			&i.ShortName,
			&i.Priority,
			&i.IsAllergy,
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
