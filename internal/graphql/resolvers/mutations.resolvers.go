package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/google/uuid"
	"github.com/mensatt/backend/internal/db/sqlc"
	"github.com/mensatt/backend/internal/graphql/gqlserver"
	"github.com/mensatt/backend/internal/graphql/models"
)

func (r *mutationResolver) CreateTag(ctx context.Context, tag sqlc.CreateTagParams) (*sqlc.Tag, error) {
	return r.Database.CreateTag(ctx, &tag)
}

func (r *mutationResolver) CreateDish(ctx context.Context, name string) (*sqlc.Dish, error) {
	return r.Database.CreateDish(ctx, name)
}

func (r *mutationResolver) RenameDish(ctx context.Context, id uuid.UUID, name string) (*sqlc.Dish, error) {
	return r.Database.RenameDish(ctx, &sqlc.RenameDishParams{
		ID:   id,
		Name: name,
	})
}

func (r *mutationResolver) CreateAlias(ctx context.Context, alias string, dish uuid.UUID) (*sqlc.DishAlias, error) {
	return r.Database.CreateDishAlias(ctx, &sqlc.CreateDishAliasParams{
		AliasName: alias,
		Dish:      dish,
	})
}

func (r *mutationResolver) UpdateAlias(ctx context.Context, oldAlias string, alias string, dish uuid.UUID) (*sqlc.DishAlias, error) {
	return r.Database.UpdateDishAlias(ctx, &sqlc.UpdateDishAliasParams{
		OldAliasName: oldAlias,
		AliasName:    alias,
		Dish:         dish,
	})
}

func (r *mutationResolver) DeleteAlias(ctx context.Context, alias string, dish uuid.UUID) (*sqlc.DishAlias, error) {
	return r.Database.DeleteDishAlias(ctx, &sqlc.DeleteDishAliasParams{
		AliasName: alias,
		Dish:      dish,
	})
}

func (r *mutationResolver) CreateOccurrence(ctx context.Context, input models.OccurrenceInputHelper) (*sqlc.Occurrence, error) {
	return r.Database.CreateOccurrenceWithSideDishesAndTags(ctx, &input.CreateOccurrenceParams, input.SideDishes, input.Tags)
}

func (r *mutationResolver) DeleteOccurrence(ctx context.Context, id uuid.UUID) (*sqlc.Occurrence, error) {
	return r.Database.DeleteOccurrence(ctx, id)
}

func (r *mutationResolver) EditOccurrence(ctx context.Context, id uuid.UUID, input sqlc.EditOccurrenceParams) (*sqlc.Occurrence, error) {
	input.ID = id
	return r.Database.EditOccurrence(ctx, &input)
}

func (r *mutationResolver) AddTagToOccurrence(ctx context.Context, occurrenceID uuid.UUID, tag string) (*sqlc.OccurrenceTag, error) {
	return r.Database.AddOccurrenceTag(ctx, &sqlc.AddOccurrenceTagParams{
		Occurrence: occurrenceID,
		Tag:        tag,
	})
}

func (r *mutationResolver) AddSideDishToOccurrence(ctx context.Context, occurrenceID uuid.UUID, sideDish uuid.UUID) (*sqlc.OccurrenceSideDish, error) {
	return r.Database.AddOccurrenceSideDish(ctx, &sqlc.AddOccurrenceSideDishParams{
		Occurrence: occurrenceID,
		Dish:       sideDish,
	})
}

func (r *mutationResolver) RemoveTagFromOccurrence(ctx context.Context, occurrenceID uuid.UUID, tag string) (*sqlc.OccurrenceTag, error) {
	return r.Database.RemoveOccurrenceTag(ctx, &sqlc.RemoveOccurrenceTagParams{
		Occurrence: occurrenceID,
		Tag:        tag,
	})
}

func (r *mutationResolver) RemoveSideDishFromOccurrence(ctx context.Context, occurrenceID uuid.UUID, sideDish uuid.UUID) (*sqlc.OccurrenceSideDish, error) {
	return r.Database.RemoveOccurrenceSideDish(ctx, &sqlc.RemoveOccurrenceSideDishParams{
		Occurrence: occurrenceID,
		Dish:       sideDish,
	})
}

func (r *mutationResolver) CreateReview(ctx context.Context, review sqlc.CreateReviewParams) (*sqlc.Review, error) {
	return r.Database.CreateReview(ctx, &review)
}

func (r *mutationResolver) DeleteReview(ctx context.Context, id uuid.UUID) (*sqlc.Review, error) {
	return r.Database.DeleteReview(ctx, id)
}

func (r *mutationResolver) EditReview(ctx context.Context, id uuid.UUID, review sqlc.EditReviewParams) (*sqlc.Review, error) {
	return r.Database.EditReview(ctx, &review)
}

// Mutation returns gqlserver.MutationResolver implementation.
func (r *Resolver) Mutation() gqlserver.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
