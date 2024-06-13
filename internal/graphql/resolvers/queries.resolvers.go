package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.49

import (
	"context"
	"fmt"

	"github.com/mensatt/backend/internal/database/ent"
	"github.com/mensatt/backend/internal/graphql/gqlserver"
	"github.com/mensatt/backend/internal/graphql/models"
	"github.com/mensatt/backend/internal/middleware"
	"github.com/mensatt/backend/pkg/utils"
)

// CurrentUser is the resolver for the currentUser field.
func (r *queryResolver) CurrentUser(ctx context.Context) (*ent.User, error) {
	return middleware.GetUserIDFromCtx(ctx), nil
}

// Tags is the resolver for the tags field.
func (r *queryResolver) Tags(ctx context.Context) ([]*ent.Tag, error) {
	return r.Database.Tag.Query().All(ctx)
}

// Dishes is the resolver for the dishes field.
func (r *queryResolver) Dishes(ctx context.Context, filter *models.DishFilter) ([]*ent.Dish, error) {
	if filter == nil {
		return r.Database.Dish.Query().All(ctx)
	}
	return r.filteredDishes(ctx, filter)
}

// Occurrences is the resolver for the occurrences field.
func (r *queryResolver) Occurrences(ctx context.Context, filter *models.OccurrenceFilter) ([]*ent.Occurrence, error) {
	if filter == nil {
		return r.Database.Occurrence.Query().All(ctx)
	}
	return r.filteredOccurrences(ctx, filter)
}

// Reviews is the resolver for the reviews field.
func (r *queryResolver) Reviews(ctx context.Context, filter *models.ReviewFilter) ([]*ent.Review, error) {
	if filter == nil {
		return r.Database.Review.Query().All(ctx)
	}
	return r.filteredReviews(ctx, filter)
}

// Locations is the resolver for the locations field.
func (r *queryResolver) Locations(ctx context.Context, filter *models.LocationFilter) ([]*ent.Location, error) {
	if filter == nil {
		return r.Database.Location.Query().All(ctx)
	}
	return r.filteredLocations(ctx, filter)
}

// VcsBuildInfo is the resolver for the vcsBuildInfo field.
func (r *queryResolver) VcsBuildInfo(ctx context.Context) (*utils.VCSBuildInfo, error) {
	if r.VCSBuildInfo == nil {
		return nil, fmt.Errorf("VCS build info is not available")
	}
	return r.VCSBuildInfo, nil
}

// Query returns gqlserver.QueryResolver implementation.
func (r *Resolver) Query() gqlserver.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
