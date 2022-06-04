// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package models

import (
	"github.com/google/uuid"
)

type DeleteDishAliasInput struct {
	Alias string `json:"alias"`
}

type DeleteImageInput struct {
	ID uuid.UUID `json:"id"`
}

type DeleteOccurrenceInput struct {
	ID uuid.UUID `json:"id"`
}

type DeleteReviewInput struct {
	ID uuid.UUID `json:"id"`
}

type DishCreateInput struct {
	Name string `json:"name"`
}
