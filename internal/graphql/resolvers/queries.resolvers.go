package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/mensatt/backend/internal/db/sqlc"
	"github.com/mensatt/backend/internal/graphql/gqlserver"
	"github.com/mensatt/backend/internal/middleware"
	"github.com/mensatt/backend/pkg/utils"
)

// CurrentUser is the resolver for the currentUser field.
func (r *queryResolver) CurrentUser(ctx context.Context) (*sqlc.User, error) {
	return middleware.GetUserIDFromCtx(ctx), nil
}

// Tags is the resolver for the tags field.
func (r *queryResolver) Tags(ctx context.Context) ([]*sqlc.Tag, error) {
	return r.Database.GetAllTags(ctx)
}

// Dishes is the resolver for the dishes field.
func (r *queryResolver) Dishes(ctx context.Context) ([]*sqlc.Dish, error) {
	return r.Database.GetAllDishes(ctx)
}

// Occurrences is the resolver for the occurrences field.
func (r *queryResolver) Occurrences(ctx context.Context, filter *sqlc.GetFilteredOccurrencesParams) ([]*sqlc.Occurrence, error) {
	return r.Database.GetFilteredOccurrences(ctx, filter)
}

// Reviews is the resolver for the reviews field.
func (r *queryResolver) Reviews(ctx context.Context) ([]*sqlc.Review, error) {
	return r.Database.GetAllReviews(ctx)
}

// Images is the resolver for the images field.
func (r *queryResolver) Images(ctx context.Context) ([]*sqlc.Image, error) {
	return r.Database.GetAllImages(ctx)
}

// Locations is the resolver for the locations field.
func (r *queryResolver) Locations(ctx context.Context) ([]*sqlc.Location, error) {
	return r.Database.GetAllLocations(ctx)
}

// LocationByID is the resolver for the locationById field.
func (r *queryResolver) LocationByID(ctx context.Context, id uuid.UUID) (*sqlc.Location, error) {
	return r.Database.GetLocationByID(ctx, id)
}

// VcsBuildInfo is the resolver for the vcsBuildInfo field.
func (r *queryResolver) VcsBuildInfo(ctx context.Context) (*utils.VCSBuildInfo, error) {
	if r.VCSBuildInfo == nil {
		return nil, errors.New("VCSBuildInfo not enabled/found")
	}
	return r.VCSBuildInfo, nil
}

// Query returns gqlserver.QueryResolver implementation.
func (r *Resolver) Query() gqlserver.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
