package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"

	"github.com/mensatt/backend/internal/db"
	"github.com/mensatt/backend/internal/db/sqlc"
	"github.com/mensatt/backend/internal/graphql/gqlserver"
	"github.com/mensatt/backend/internal/graphql/models"
	"github.com/mensatt/backend/pkg/utils"
)

// LoginUser is the resolver for the loginUser field.
func (r *mutationResolver) LoginUser(ctx context.Context, input models.LoginUserInput) (string, error) {
	user, err := r.Database.GetUserByEmail(ctx, input.Email)
	if err != nil {
		return "", err
	}

	if !utils.CheckPasswordHash(input.Password, user.PasswordHash) {
		return "", errors.New("wrong password")
	}

	return r.JWTKeyStore.GenerateJWT(user.ID)
}

// CreateTag is the resolver for the createTag field.
func (r *mutationResolver) CreateTag(ctx context.Context, input sqlc.CreateTagParams) (*sqlc.Tag, error) {
	return r.Database.CreateTag(ctx, &input)
}

// CreateDish is the resolver for the createDish field.
func (r *mutationResolver) CreateDish(ctx context.Context, input sqlc.CreateDishParams) (*sqlc.Dish, error) {
	return r.Database.CreateDish(ctx, &input)
}

// UpdateDish is the resolver for the updateDish field.
func (r *mutationResolver) UpdateDish(ctx context.Context, input sqlc.UpdateDishParams) (*sqlc.Dish, error) {
	return r.Database.UpdateDish(ctx, &input)
}

// CreateDishAlias is the resolver for the createDishAlias field.
func (r *mutationResolver) CreateDishAlias(ctx context.Context, input sqlc.CreateDishAliasParams) (*sqlc.DishAlias, error) {
	return r.Database.CreateDishAlias(ctx, &input)
}

// UpdateDishAlias is the resolver for the updateDishAlias field.
func (r *mutationResolver) UpdateDishAlias(ctx context.Context, input sqlc.UpdateDishAliasParams) (*sqlc.DishAlias, error) {
	return r.Database.UpdateDishAlias(ctx, &input)
}

// DeleteDishAlias is the resolver for the deleteDishAlias field.
func (r *mutationResolver) DeleteDishAlias(ctx context.Context, input models.DeleteDishAliasInput) (*sqlc.DishAlias, error) {
	return r.Database.DeleteDishAlias(ctx, input.Alias)
}

// CreateOccurrence is the resolver for the createOccurrence field.
func (r *mutationResolver) CreateOccurrence(ctx context.Context, input models.CreateOccurrenceInputHelper) (*sqlc.Occurrence, error) {
	return r.Database.CreateOccurrenceWithSideDishesAndTags(ctx, &input.CreateOccurrenceParams, input.SideDishes, input.Tags)
}

// UpdateOccurrence is the resolver for the updateOccurrence field.
func (r *mutationResolver) UpdateOccurrence(ctx context.Context, input sqlc.UpdateOccurrenceParams) (*sqlc.Occurrence, error) {
	return r.Database.UpdateOccurrence(ctx, &input)
}

// DeleteOccurrence is the resolver for the deleteOccurrence field.
func (r *mutationResolver) DeleteOccurrence(ctx context.Context, input models.DeleteOccurrenceInput) (*sqlc.Occurrence, error) {
	return r.Database.DeleteOccurrence(ctx, input.ID)
}

// AddTagToOccurrence is the resolver for the addTagToOccurrence field.
func (r *mutationResolver) AddTagToOccurrence(ctx context.Context, input sqlc.AddOccurrenceTagParams) (*sqlc.OccurrenceTag, error) {
	return r.Database.AddOccurrenceTag(ctx, &input)
}

// RemoveTagFromOccurrence is the resolver for the removeTagFromOccurrence field.
func (r *mutationResolver) RemoveTagFromOccurrence(ctx context.Context, input sqlc.RemoveOccurrenceTagParams) (*sqlc.OccurrenceTag, error) {
	return r.Database.RemoveOccurrenceTag(ctx, &input)
}

// AddSideDishToOccurrence is the resolver for the addSideDishToOccurrence field.
func (r *mutationResolver) AddSideDishToOccurrence(ctx context.Context, input sqlc.AddOccurrenceSideDishParams) (*sqlc.OccurrenceSideDish, error) {
	return r.Database.AddOccurrenceSideDish(ctx, &input)
}

// RemoveSideDishFromOccurrence is the resolver for the removeSideDishFromOccurrence field.
func (r *mutationResolver) RemoveSideDishFromOccurrence(ctx context.Context, input sqlc.RemoveOccurrenceSideDishParams) (*sqlc.OccurrenceSideDish, error) {
	return r.Database.RemoveOccurrenceSideDish(ctx, &input)
}

// CreateReview is the resolver for the createReview field.
func (r *mutationResolver) CreateReview(ctx context.Context, input db.CreateReviewWithImagesParams) (*sqlc.Review, error) {
	if len(input.Images) == 0 {
		return r.Database.CreateReview(ctx, &input.CreateReviewParams)
	}
	return r.Database.CreateReviewWithImages(ctx, &input)
}

// UpdateReview is the resolver for the updateReview field.
func (r *mutationResolver) UpdateReview(ctx context.Context, input sqlc.UpdateReviewParams) (*sqlc.Review, error) {
	return r.Database.UpdateReview(ctx, &input)
}

// DeleteReview is the resolver for the deleteReview field.
func (r *mutationResolver) DeleteReview(ctx context.Context, input models.DeleteReviewInput) (*sqlc.Review, error) {
	return r.Database.DeleteReview(ctx, input.ID)
}

// Mutation returns gqlserver.MutationResolver implementation.
func (r *Resolver) Mutation() gqlserver.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
