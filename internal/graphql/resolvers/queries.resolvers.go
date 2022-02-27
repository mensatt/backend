package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"time"

	"github.com/mensatt/mensatt-backend/internal/db"
	"github.com/mensatt/mensatt-backend/internal/graphql/gqlserver"
)

func (r *queryResolver) GetAllTags(ctx context.Context) ([]*db.Tag, error) {
	return r.Database.GetAllTags(ctx)
}

func (r *queryResolver) GetAllDishes(ctx context.Context) ([]*db.Dish, error) {
	return r.Database.GetAllDishes(ctx)
}

func (r *queryResolver) GetAllImages(ctx context.Context) ([]*db.Image, error) {
	return r.Database.GetAllImages(ctx)
}

func (r *queryResolver) GetAllReviews(ctx context.Context) ([]*db.Review, error) {
	return r.Database.GetAllReviews(ctx)
}

func (r *queryResolver) GetOccurrencesByDate(ctx context.Context, date time.Time) ([]*db.Occurrence, error) {
	return r.Database.GetOccurrencesByDate(ctx, date)
}

// Query returns gqlserver.QueryResolver implementation.
func (r *Resolver) Query() gqlserver.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
