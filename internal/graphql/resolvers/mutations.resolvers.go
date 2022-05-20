package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/mensatt/mensatt-backend/internal/db"
	"github.com/mensatt/mensatt-backend/internal/graphql/gqlserver"
	"github.com/mensatt/mensatt-backend/internal/graphql/models"
)

func (r *mutationResolver) CreateTag(ctx context.Context, tag db.CreateTagParams) (*db.Tag, error) {
	return r.Database.CreateTag(ctx, &tag)
}

func (r *mutationResolver) CreateDish(ctx context.Context, name string) (*db.Dish, error) {
	return r.Database.CreateDish(ctx, name)
}

func (r *mutationResolver) CreateOccurrence(ctx context.Context, occurrence models.OccurrenceInputHelper) (*db.Occurrence, error) {
	return r.Database.CreateOccurrenceWithSideDishesAndTags(ctx, &occurrence.CreateOccurrenceParams, occurrence.SideDishes, occurrence.Tags)
}

func (r *mutationResolver) CreateReview(ctx context.Context, review db.CreateReviewParams) (*db.Review, error) {
	return r.Database.CreateReview(ctx, &review)
}

// Mutation returns gqlserver.MutationResolver implementation.
func (r *Resolver) Mutation() gqlserver.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
