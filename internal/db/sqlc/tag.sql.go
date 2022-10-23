// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0
// source: tag.sql

package sqlc

import (
	"context"
	"database/sql"
)

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
