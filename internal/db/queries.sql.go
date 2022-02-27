// Code generated by sqlc. DO NOT EDIT.
// source: queries.sql

package db

import (
	"context"

	"github.com/google/uuid"
)

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

const getOccurrenceByID = `-- name: GetOccurrenceByID :one
SELECT id, dish, date, price_student, price_staff, price_guest 
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
		&i.PriceStudent,
		&i.PriceStaff,
		&i.PriceGuest,
	)
	return &i, err
}

const getSideDishesForOccurrence = `-- name: GetSideDishesForOccurrence :many
SELECT dish.id, dish.name
FROM occurrence_side_dishes JOIN dish ON occurrence_side_dishes.dish_id = dish.id
WHERE occurrence_side_dishes.occurrence_id = $1
`

func (q *Queries) GetSideDishesForOccurrence(ctx context.Context, occurrenceID uuid.UUID) ([]*Dish, error) {
	rows, err := q.db.Query(ctx, getSideDishesForOccurrence, occurrenceID)
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
FROM occurrence_tag JOIN tag ON occurrence_tag.tag_abbreviation = tag.abbreviation
WHERE occurrence_tag.occurrence_id = $1
`

func (q *Queries) GetTagsForOccurrence(ctx context.Context, occurrenceID uuid.UUID) ([]*Tag, error) {
	rows, err := q.db.Query(ctx, getTagsForOccurrence, occurrenceID)
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
