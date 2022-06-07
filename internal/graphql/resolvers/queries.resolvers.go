package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/mensatt/backend/internal/db/sqlc"
	"github.com/mensatt/backend/internal/graphql/gqlserver"
	"github.com/mensatt/backend/internal/middleware"
	"github.com/mensatt/backend/pkg/utils"
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

func (r *queryResolver) GetOccurrencesByDate(ctx context.Context, date time.Time) ([]*sqlc.Occurrence, error) {
	return r.Database.GetOccurrencesByDate(ctx, date)
}

func (r *queryResolver) GetOccurrencesAfterInclusiveDate(ctx context.Context, start time.Time) ([]*sqlc.Occurrence, error) {
	return r.Database.GetOccurrencesAfterInclusiveDate(ctx, start)
}

func (r *queryResolver) GetAllReviews(ctx context.Context) ([]*sqlc.Review, error) {
	return r.Database.GetAllReviews(ctx)
}

func (r *queryResolver) GetAllImages(ctx context.Context) ([]*sqlc.Image, error) {
	return r.Database.GetAllImages(ctx)
}

func (r *queryResolver) GetAllLocations(ctx context.Context) ([]*sqlc.Location, error) {
	return r.Database.GetAllLocations(ctx)
}

func (r *queryResolver) GetLocationByID(ctx context.Context, id uuid.UUID) (*sqlc.Location, error) {
	return r.Database.GetLocationByID(ctx, id)
}

func (r *queryResolver) GetVcsBuildInfo(ctx context.Context) (*utils.VCSBuildInfo, error) {
	if r.VCSBuildInfo == nil {
		return nil, errors.New("VCSBuildInfo not enabled/found")
	}
	return r.VCSBuildInfo, nil
}

// Query returns gqlserver.QueryResolver implementation.
func (r *Resolver) Query() gqlserver.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
