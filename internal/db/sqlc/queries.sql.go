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

const createDishAlias = `-- name: CreateDishAlias :one
INSERT INTO dish_alias (alias_name, dish)
VALUES ($1, $2)
RETURNING alias_name, dish
`

type CreateDishAliasParams struct {
	AliasName string    `json:"alias_name"`
	Dish      uuid.UUID `json:"dish"`
}

func (q *Queries) CreateDishAlias(ctx context.Context, arg *CreateDishAliasParams) (*DishAlias, error) {
	row := q.db.QueryRow(ctx, createDishAlias, arg.AliasName, arg.Dish)
	var i DishAlias
	err := row.Scan(&i.AliasName, &i.Dish)
	return &i, err
}

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

const getAliasesForDish = `-- name: GetAliasesForDish :many
SELECT alias_name
FROM dish_alias
WHERE dish = $1
`

func (q *Queries) GetAliasesForDish(ctx context.Context, dish uuid.UUID) ([]string, error) {
	rows, err := q.db.Query(ctx, getAliasesForDish, dish)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []string
	for rows.Next() {
		var alias_name string
		if err := rows.Scan(&alias_name); err != nil {
			return nil, err
		}
		items = append(items, alias_name)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getAllAliases = `-- name: GetAllAliases :many
SELECT alias_name, dish
FROM dish_alias
`

func (q *Queries) GetAllAliases(ctx context.Context) ([]*DishAlias, error) {
	rows, err := q.db.Query(ctx, getAllAliases)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []*DishAlias
	for rows.Next() {
		var i DishAlias
		if err := rows.Scan(&i.AliasName, &i.Dish); err != nil {
			return nil, err
		}
		items = append(items, &i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
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

const removeOccurrenceSideDish = `-- name: RemoveOccurrenceSideDish :one
DELETE FROM occurrence_side_dishes
WHERE occurrence = $1 AND dish = $2
RETURNING occurrence, dish
`

type RemoveOccurrenceSideDishParams struct {
	Occurrence uuid.UUID `json:"occurrence"`
	Dish       uuid.UUID `json:"dish"`
}

func (q *Queries) RemoveOccurrenceSideDish(ctx context.Context, arg *RemoveOccurrenceSideDishParams) (*OccurrenceSideDish, error) {
	row := q.db.QueryRow(ctx, removeOccurrenceSideDish, arg.Occurrence, arg.Dish)
	var i OccurrenceSideDish
	err := row.Scan(&i.Occurrence, &i.Dish)
	return &i, err
}

const removeOccurrenceTag = `-- name: RemoveOccurrenceTag :one
DELETE FROM occurrence_tag
WHERE occurrence = $1 AND tag = $2
RETURNING occurrence, tag
`

type RemoveOccurrenceTagParams struct {
	Occurrence uuid.UUID `json:"occurrence"`
	Tag        string    `json:"tag"`
}

func (q *Queries) RemoveOccurrenceTag(ctx context.Context, arg *RemoveOccurrenceTagParams) (*OccurrenceTag, error) {
	row := q.db.QueryRow(ctx, removeOccurrenceTag, arg.Occurrence, arg.Tag)
	var i OccurrenceTag
	err := row.Scan(&i.Occurrence, &i.Tag)
	return &i, err
}
