// Code generated by sqlc. DO NOT EDIT.

package db

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type Querier interface {
	GetAllDishes(ctx context.Context) ([]*Dish, error)
	GetAllImages(ctx context.Context) ([]*Image, error)
	GetAllReviews(ctx context.Context) ([]*Review, error)
	GetAllTags(ctx context.Context) ([]*Tag, error)
	GetDishByID(ctx context.Context, id uuid.UUID) (*Dish, error)
	GetImagesForOccurrence(ctx context.Context, occurrence uuid.UUID) ([]*Image, error)
	GetOccurrenceByID(ctx context.Context, id uuid.UUID) (*Occurrence, error)
	GetOccurrencesByDate(ctx context.Context, date time.Time) ([]*Occurrence, error)
	GetSideDishesForOccurrence(ctx context.Context, occurrenceID uuid.UUID) ([]*Dish, error)
	GetTagsForOccurrence(ctx context.Context, occurrenceID uuid.UUID) ([]*Tag, error)
}

var _ Querier = (*Queries)(nil)
