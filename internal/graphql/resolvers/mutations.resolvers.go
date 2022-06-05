package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/mensatt/backend/internal/db/sqlc"
	"github.com/mensatt/backend/internal/graphql/gqlserver"
	"github.com/mensatt/backend/internal/graphql/models"
)

func (r *mutationResolver) CreateTag(ctx context.Context, input sqlc.CreateTagParams) (*sqlc.Tag, error) {
	return r.Database.CreateTag(ctx, &input)
}

func (r *mutationResolver) CreateDish(ctx context.Context, input models.DishCreateInput) (*sqlc.Dish, error) {
	return r.Database.CreateDish(ctx, input.Name)
}

func (r *mutationResolver) UpdateDish(ctx context.Context, input sqlc.UpdateDishParams) (*sqlc.Dish, error) {
	return r.Database.UpdateDish(ctx, &input)
}

func (r *mutationResolver) CreateDishAlias(ctx context.Context, input sqlc.CreateDishAliasParams) (*sqlc.DishAlias, error) {
	return r.Database.CreateDishAlias(ctx, &input)
}

func (r *mutationResolver) UpdateDishAlias(ctx context.Context, input sqlc.UpdateDishAliasParams) (*sqlc.DishAlias, error) {
	return r.Database.UpdateDishAlias(ctx, &input)
}

func (r *mutationResolver) DeleteDishAlias(ctx context.Context, input models.DeleteDishAliasInput) (*sqlc.DishAlias, error) {
	return r.Database.DeleteDishAlias(ctx, input.Alias)
}

func (r *mutationResolver) CreateOccurrence(ctx context.Context, input models.OccurrenceInputHelper) (*sqlc.Occurrence, error) {
	return r.Database.CreateOccurrenceWithSideDishesAndTags(ctx, &input.CreateOccurrenceParams, input.SideDishes, input.Tags)
}

func (r *mutationResolver) UpdateOccurrence(ctx context.Context, input sqlc.UpdateOccurrenceParams) (*sqlc.Occurrence, error) {
	return r.Database.UpdateOccurrence(ctx, &input)
}

func (r *mutationResolver) DeleteOccurrence(ctx context.Context, input models.DeleteOccurrenceInput) (*sqlc.Occurrence, error) {
	return r.Database.DeleteOccurrence(ctx, input.ID)
}

func (r *mutationResolver) AddTagToOccurrence(ctx context.Context, input sqlc.AddOccurrenceTagParams) (*sqlc.OccurrenceTag, error) {
	return r.Database.AddOccurrenceTag(ctx, &input)
}

func (r *mutationResolver) RemoveTagFromOccurrence(ctx context.Context, input sqlc.RemoveOccurrenceTagParams) (*sqlc.OccurrenceTag, error) {
	return r.Database.RemoveOccurrenceTag(ctx, &input)
}

func (r *mutationResolver) AddSideDishToOccurrence(ctx context.Context, input sqlc.AddOccurrenceSideDishParams) (*sqlc.OccurrenceSideDish, error) {
	return r.Database.AddOccurrenceSideDish(ctx, &input)
}

func (r *mutationResolver) RemoveSideDishFromOccurrence(ctx context.Context, input sqlc.RemoveOccurrenceSideDishParams) (*sqlc.OccurrenceSideDish, error) {
	return r.Database.RemoveOccurrenceSideDish(ctx, &input)
}

func (r *mutationResolver) CreateReview(ctx context.Context, input sqlc.CreateReviewParams) (*sqlc.Review, error) {
	return r.Database.CreateReview(ctx, &input)
}

func (r *mutationResolver) UpdateReview(ctx context.Context, input sqlc.UpdateReviewParams) (*sqlc.Review, error) {
	return r.Database.UpdateReview(ctx, &input)
}

func (r *mutationResolver) DeleteReview(ctx context.Context, input models.DeleteReviewInput) (*sqlc.Review, error) {
	return r.Database.DeleteReview(ctx, input.ID)
}

// Mutation returns gqlserver.MutationResolver implementation.
func (r *Resolver) Mutation() gqlserver.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
