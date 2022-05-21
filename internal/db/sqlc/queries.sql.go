// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.13.0
// source: queries.sql

package sqlc

import (
	"context"
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type AddMultipleOccurrenceSideDishesParams struct {
	Occurrence uuid.UUID `json:"occurrence"`
	Dish       uuid.UUID `json:"dish"`
}

type AddMultipleOccurrenceTagsParams struct {
	Occurrence uuid.UUID `json:"occurrence"`
	Tag        string    `json:"tag"`
}

const addOccurrenceSideDish = `-- name: AddOccurrenceSideDish :one
INSERT INTO occurrence_side_dishes (occurrence, dish)
VALUES ($1, $2)
RETURNING occurrence, dish
`

type AddOccurrenceSideDishParams struct {
	Occurrence uuid.UUID `json:"occurrence"`
	Dish       uuid.UUID `json:"dish"`
}

func (q *Queries) AddOccurrenceSideDish(ctx context.Context, arg *AddOccurrenceSideDishParams) (*OccurrenceSideDish, error) {
	row := q.db.QueryRow(ctx, addOccurrenceSideDish, arg.Occurrence, arg.Dish)
	var i OccurrenceSideDish
	err := row.Scan(&i.Occurrence, &i.Dish)
	return &i, err
}

const addOccurrenceTag = `-- name: AddOccurrenceTag :one
INSERT INTO occurrence_tag (occurrence, tag)
VALUES ($1, $2)
RETURNING occurrence, tag
`

type AddOccurrenceTagParams struct {
	Occurrence uuid.UUID `json:"occurrence"`
	Tag        string    `json:"tag"`
}

func (q *Queries) AddOccurrenceTag(ctx context.Context, arg *AddOccurrenceTagParams) (*OccurrenceTag, error) {
	row := q.db.QueryRow(ctx, addOccurrenceTag, arg.Occurrence, arg.Tag)
	var i OccurrenceTag
	err := row.Scan(&i.Occurrence, &i.Tag)
	return &i, err
}

const createDish = `-- name: CreateDish :one
INSERT INTO dish (name)
VALUES ($1)
RETURNING id, name
`

func (q *Queries) CreateDish(ctx context.Context, name string) (*Dish, error) {
	row := q.db.QueryRow(ctx, createDish, name)
	var i Dish
	err := row.Scan(&i.ID, &i.Name)
	return &i, err
}

const createOccurrence = `-- name: CreateOccurrence :one
INSERT INTO occurrence (dish, date, kj, kcal, fat, saturated_fat, carbohydrates, sugar, fiber, protein, salt, price_student, price_staff, price_guest)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14)
RETURNING id, dish, date, kj, kcal, fat, saturated_fat, carbohydrates, sugar, fiber, protein, salt, price_student, price_staff, price_guest
`

type CreateOccurrenceParams struct {
	Dish          uuid.UUID       `json:"dish"`
	Date          time.Time       `json:"date"`
	Kj            sql.NullFloat64 `json:"kj"`
	Kcal          sql.NullFloat64 `json:"kcal"`
	Fat           sql.NullFloat64 `json:"fat"`
	SaturatedFat  sql.NullFloat64 `json:"saturated_fat"`
	Carbohydrates sql.NullFloat64 `json:"carbohydrates"`
	Sugar         sql.NullFloat64 `json:"sugar"`
	Fiber         sql.NullFloat64 `json:"fiber"`
	Protein       sql.NullFloat64 `json:"protein"`
	Salt          sql.NullFloat64 `json:"salt"`
	PriceStudent  int32           `json:"price_student"`
	PriceStaff    int32           `json:"price_staff"`
	PriceGuest    int32           `json:"price_guest"`
}

func (q *Queries) CreateOccurrence(ctx context.Context, arg *CreateOccurrenceParams) (*Occurrence, error) {
	row := q.db.QueryRow(ctx, createOccurrence,
		arg.Dish,
		arg.Date,
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

const createTag = `-- name: CreateTag :one
INSERT INTO tag (key, name, description, short_name, priority, is_allergy)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING key, name, description, short_name, priority, is_allergy
`

type CreateTagParams struct {
	Key         string         `json:"key"`
	Name        string         `json:"name"`
	Description string         `json:"description"`
	ShortName   sql.NullString `json:"short_name"`
	Priority    Priority       `json:"priority"`
	IsAllergy   bool           `json:"is_allergy"`
}

func (q *Queries) CreateTag(ctx context.Context, arg *CreateTagParams) (*Tag, error) {
	row := q.db.QueryRow(ctx, createTag,
		arg.Key,
		arg.Name,
		arg.Description,
		arg.ShortName,
		arg.Priority,
		arg.IsAllergy,
	)
	var i Tag
	err := row.Scan(
		&i.Key,
		&i.Name,
		&i.Description,
		&i.ShortName,
		&i.Priority,
		&i.IsAllergy,
	)
	return &i, err
}

const getAllDishes = `-- name: GetAllDishes :many
SELECT id, name
FROM dish
`

func (q *Queries) GetAllDishes(ctx context.Context) ([]*Dish, error) {
	rows, err := q.db.Query(ctx, getAllDishes)
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

const getAllOccurrences = `-- name: GetAllOccurrences :many
SELECT id, dish, date, kj, kcal, fat, saturated_fat, carbohydrates, sugar, fiber, protein, salt, price_student, price_staff, price_guest
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

const getAllTags = `-- name: GetAllTags :many
SELECT key, name, description, short_name, priority, is_allergy
FROM tag
`

func (q *Queries) GetAllTags(ctx context.Context) ([]*Tag, error) {
	rows, err := q.db.Query(ctx, getAllTags)
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

const getDishByID = `-- name: GetDishByID :one
SELECT id, name
FROM dish
WHERE id = $1
`

func (q *Queries) GetDishByID(ctx context.Context, id uuid.UUID) (*Dish, error) {
	row := q.db.QueryRow(ctx, getDishByID, id)
	var i Dish
	err := row.Scan(&i.ID, &i.Name)
	return &i, err
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
SELECT id, dish, date, kj, kcal, fat, saturated_fat, carbohydrates, sugar, fiber, protein, salt, price_student, price_staff, price_guest 
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
SELECT id, dish, date, kj, kcal, fat, saturated_fat, carbohydrates, sugar, fiber, protein, salt, price_student, price_staff, price_guest
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

const getTagByKey = `-- name: GetTagByKey :one
SELECT key, name, description, short_name, priority, is_allergy
FROM tag
WHERE key = $1
`

func (q *Queries) GetTagByKey(ctx context.Context, key string) (*Tag, error) {
	row := q.db.QueryRow(ctx, getTagByKey, key)
	var i Tag
	err := row.Scan(
		&i.Key,
		&i.Name,
		&i.Description,
		&i.ShortName,
		&i.Priority,
		&i.IsAllergy,
	)
	return &i, err
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
