// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.14.0
// source: occurrence.sql

package sqlc

import (
	"context"
	"database/sql"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgtype"
)

const createOccurrence = `-- name: CreateOccurrence :one
INSERT INTO occurrence (location, dish, date, review_status, kj, kcal, fat, saturated_fat, carbohydrates, sugar, fiber, protein, salt, price_student, price_staff, price_guest)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16)
RETURNING id, location, dish, date, review_status, kj, kcal, fat, saturated_fat, carbohydrates, sugar, fiber, protein, salt, price_student, price_staff, price_guest
`

type CreateOccurrenceParams struct {
	Location      uuid.UUID     `json:"location"`
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
		arg.Location,
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
		&i.Location,
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
RETURNING id, location, dish, date, review_status, kj, kcal, fat, saturated_fat, carbohydrates, sugar, fiber, protein, salt, price_student, price_staff, price_guest
`

func (q *Queries) DeleteOccurrence(ctx context.Context, id uuid.UUID) (*Occurrence, error) {
	row := q.db.QueryRow(ctx, deleteOccurrence, id)
	var i Occurrence
	err := row.Scan(
		&i.ID,
		&i.Location,
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
SELECT id, location, dish, date, review_status, kj, kcal, fat, saturated_fat, carbohydrates, sugar, fiber, protein, salt, price_student, price_staff, price_guest
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
			&i.Location,
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
SELECT image.id, image.image_store_id, image.review
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
		if err := rows.Scan(&i.ID, &i.ImageStoreID, &i.Review); err != nil {
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
SELECT id, location, dish, date, review_status, kj, kcal, fat, saturated_fat, carbohydrates, sugar, fiber, protein, salt, price_student, price_staff, price_guest 
FROM occurrence
WHERE id = $1
`

func (q *Queries) GetOccurrenceByID(ctx context.Context, id uuid.UUID) (*Occurrence, error) {
	row := q.db.QueryRow(ctx, getOccurrenceByID, id)
	var i Occurrence
	err := row.Scan(
		&i.ID,
		&i.Location,
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

const getOccurrenceReviewMetadata = `-- name: GetOccurrenceReviewMetadata :one
SELECT AVG(review.stars) AS average_stars, COUNT(*) AS review_count
FROM review
WHERE review.occurrence = $1
`

type GetOccurrenceReviewMetadataRow struct {
	AverageStars pgtype.Numeric `json:"average_stars"`
	ReviewCount  int64          `json:"review_count"`
}

func (q *Queries) GetOccurrenceReviewMetadata(ctx context.Context, occurrence uuid.UUID) (*GetOccurrenceReviewMetadataRow, error) {
	row := q.db.QueryRow(ctx, getOccurrenceReviewMetadata, occurrence)
	var i GetOccurrenceReviewMetadataRow
	err := row.Scan(&i.AverageStars, &i.ReviewCount)
	return &i, err
}

const getOccurrencesAfterInclusiveDate = `-- name: GetOccurrencesAfterInclusiveDate :many
SELECT id, location, dish, date, review_status, kj, kcal, fat, saturated_fat, carbohydrates, sugar, fiber, protein, salt, price_student, price_staff, price_guest
FROM occurrence
WHERE date >= $1
`

func (q *Queries) GetOccurrencesAfterInclusiveDate(ctx context.Context, date time.Time) ([]*Occurrence, error) {
	rows, err := q.db.Query(ctx, getOccurrencesAfterInclusiveDate, date)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []*Occurrence
	for rows.Next() {
		var i Occurrence
		if err := rows.Scan(
			&i.ID,
			&i.Location,
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

const getOccurrencesByDate = `-- name: GetOccurrencesByDate :many
SELECT id, location, dish, date, review_status, kj, kcal, fat, saturated_fat, carbohydrates, sugar, fiber, protein, salt, price_student, price_staff, price_guest
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
			&i.Location,
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
SELECT dish.id, dish.name_de, dish.name_en
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
		if err := rows.Scan(&i.ID, &i.NameDe, &i.NameEn); err != nil {
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

const updateOccurrence = `-- name: UpdateOccurrence :one
UPDATE occurrence
SET 
    dish = COALESCE($2, dish),
    date = COALESCE($3, date),
    review_status = COALESCE($4, review_status),
    kj = COALESCE($5, kj),
    kcal = COALESCE($6, kcal),
    fat = COALESCE($7, fat),
    saturated_fat = COALESCE($8, saturated_fat),
    carbohydrates = COALESCE($9, carbohydrates),
    sugar = COALESCE($10, sugar),
    fiber = COALESCE($11, fiber),
    protein = COALESCE($12, protein),
    salt = COALESCE($13, salt),
    price_student = COALESCE($14, price_student),
    price_staff = COALESCE($15, price_staff),
    price_guest = COALESCE($16, price_guest)
WHERE id = $1
RETURNING id, location, dish, date, review_status, kj, kcal, fat, saturated_fat, carbohydrates, sugar, fiber, protein, salt, price_student, price_staff, price_guest
`

type UpdateOccurrenceParams struct {
	ID            uuid.UUID     `json:"id"`
	Dish          uuid.NullUUID `json:"dish"`
	Date          sql.NullTime  `json:"date"`
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

func (q *Queries) UpdateOccurrence(ctx context.Context, arg *UpdateOccurrenceParams) (*Occurrence, error) {
	row := q.db.QueryRow(ctx, updateOccurrence,
		arg.ID,
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
		&i.Location,
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
