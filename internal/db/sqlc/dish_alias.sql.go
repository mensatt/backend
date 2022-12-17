// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0
// source: dish_alias.sql

package sqlc

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
)

const createDishAlias = `-- name: CreateDishAlias :one
INSERT INTO dish_alias (alias_name, normalized_alias_name, dish)
VALUES ($1, $2, $3)
RETURNING alias_name, normalized_alias_name, dish
`

type CreateDishAliasParams struct {
	AliasName           string    `json:"alias_name"`
	NormalizedAliasName string    `json:"normalized_alias_name"`
	Dish                uuid.UUID `json:"dish"`
}

func (q *Queries) CreateDishAlias(ctx context.Context, arg *CreateDishAliasParams) (*DishAlias, error) {
	row := q.db.QueryRow(ctx, createDishAlias, arg.AliasName, arg.NormalizedAliasName, arg.Dish)
	var i DishAlias
	err := row.Scan(&i.AliasName, &i.NormalizedAliasName, &i.Dish)
	return &i, err
}

const deleteDishAlias = `-- name: DeleteDishAlias :one
DELETE FROM dish_alias
WHERE alias_name = $1
RETURNING alias_name, normalized_alias_name, dish
`

func (q *Queries) DeleteDishAlias(ctx context.Context, aliasName string) (*DishAlias, error) {
	row := q.db.QueryRow(ctx, deleteDishAlias, aliasName)
	var i DishAlias
	err := row.Scan(&i.AliasName, &i.NormalizedAliasName, &i.Dish)
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
SELECT alias_name, normalized_alias_name, dish
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
		if err := rows.Scan(&i.AliasName, &i.NormalizedAliasName, &i.Dish); err != nil {
			return nil, err
		}
		items = append(items, &i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const mergeAliases = `-- name: MergeAliases :exec
UPDATE dish_alias
SET dish = $1
WHERE dish = $2
`

type MergeAliasesParams struct {
	Keep  uuid.UUID `json:"keep"`
	Merge uuid.UUID `json:"merge"`
}

func (q *Queries) MergeAliases(ctx context.Context, arg *MergeAliasesParams) error {
	_, err := q.db.Exec(ctx, mergeAliases, arg.Keep, arg.Merge)
	return err
}

const updateDishAlias = `-- name: UpdateDishAlias :one
UPDATE dish_alias
SET 
    alias_name = COALESCE($2, alias_name),
    normalized_alias_name = COALESCE($3, normalized_alias_name),
    dish = COALESCE($4, dish)
WHERE alias_name = $1
RETURNING alias_name, normalized_alias_name, dish
`

type UpdateDishAliasParams struct {
	AliasName           string         `json:"alias_name"`
	NewAliasName        sql.NullString `json:"new_alias_name"`
	NormalizedAliasName sql.NullString `json:"normalized_alias_name"`
	Dish                uuid.NullUUID  `json:"dish"`
}

func (q *Queries) UpdateDishAlias(ctx context.Context, arg *UpdateDishAliasParams) (*DishAlias, error) {
	row := q.db.QueryRow(ctx, updateDishAlias,
		arg.AliasName,
		arg.NewAliasName,
		arg.NormalizedAliasName,
		arg.Dish,
	)
	var i DishAlias
	err := row.Scan(&i.AliasName, &i.NormalizedAliasName, &i.Dish)
	return &i, err
}
