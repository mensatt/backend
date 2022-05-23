package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"time"

	"github.com/mensatt/mensatt-backend/internal/db/sqlc"
	"github.com/mensatt/mensatt-backend/internal/graphql/gqlserver"
	"github.com/mensatt/mensatt-backend/internal/middleware"
)

func (r *queryResolver) GetCurrentUser(ctx context.Context) (*sqlc.User, error) {
	return middleware.GetUserIDFromCtx(ctx), nil
}

func (r *queryResolver) GetAllTags(ctx context.Context) ([]*sqlc.Tag, error) {
	return r.Database.GetAllTags(ctx)
}

func (r *queryResolver) GetAllDishes(ctx context.Context) ([]*sqlc.Dish, error) {
	return r.Database.GetAllDishes(ctx)
}

func (r *queryResolver) GetAllOccurrences(ctx context.Context) ([]*sqlc.Occurrence, error) {
	return r.Database.GetAllOccurrences(ctx)
}

func (r *queryResolver) GetAllReviews(ctx context.Context) ([]*sqlc.Review, error) {
	return r.Database.GetAllReviews(ctx)
}

func (r *queryResolver) GetAllImages(ctx context.Context) ([]*sqlc.Image, error) {
	return r.Database.GetAllImages(ctx)
}

func (r *queryResolver) GetOccurrencesByDate(ctx context.Context, date time.Time) ([]*sqlc.Occurrence, error) {
	return r.Database.GetOccurrencesByDate(ctx, date)
}

// Query returns gqlserver.QueryResolver implementation.
func (r *Resolver) Query() gqlserver.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
