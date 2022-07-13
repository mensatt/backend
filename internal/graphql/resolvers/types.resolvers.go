package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/mensatt/backend/internal/db/sqlc"
	"github.com/mensatt/backend/internal/graphql/gqlserver"
)

func (r *dishResolver) Aliases(ctx context.Context, obj *sqlc.Dish) ([]string, error) {
	return r.Database.GetAliasesForDish(ctx, obj.ID)
}

func (r *dishResolver) Images(ctx context.Context, obj *sqlc.Dish) ([]*sqlc.Image, error) {
	return r.Database.GetImagesByDish(ctx, obj.ID)
}

func (r *dishResolver) Reviews(ctx context.Context, obj *sqlc.Dish) ([]*sqlc.Review, error) {
	return r.Database.GetReviewsByDish(ctx, obj.ID)
}

func (r *imageResolver) Occurrence(ctx context.Context, obj *sqlc.Image) (*sqlc.Occurrence, error) {
	return r.Database.GetOccurrenceByID(ctx, obj.Occurrence)
}

func (r *imageResolver) ImageURL(ctx context.Context, obj *sqlc.Image) (string, error) {
	return fmt.Sprintf("%s/%s", r.ImageBaseURL, obj.ImageStoreID), nil
}

func (r *occurrenceResolver) Location(ctx context.Context, obj *sqlc.Occurrence) (*sqlc.Location, error) {
	return r.Database.GetLocationByID(ctx, obj.Location)
}

func (r *occurrenceResolver) Dish(ctx context.Context, obj *sqlc.Occurrence) (*sqlc.Dish, error) {
	return r.Database.GetDishByID(ctx, obj.Dish)
}

func (r *occurrenceResolver) SideDishes(ctx context.Context, obj *sqlc.Occurrence) ([]*sqlc.Dish, error) {
	return r.Database.GetSideDishesForOccurrence(ctx, obj.ID)
}

func (r *occurrenceResolver) Tags(ctx context.Context, obj *sqlc.Occurrence) ([]*sqlc.Tag, error) {
	return r.Database.GetTagsForOccurrence(ctx, obj.ID)
}

func (r *occurrenceResolver) Reviews(ctx context.Context, obj *sqlc.Occurrence) ([]*sqlc.Review, error) {
	return r.Database.GetReviewsForOccurrence(ctx, obj.ID)
}

func (r *occurrenceResolver) AverageReviewStars(ctx context.Context, obj *sqlc.Occurrence) (float64, error) {
	return r.Database.GetAverageReviewStars(ctx, obj.ID)
}

func (r *occurrenceResolver) Images(ctx context.Context, obj *sqlc.Occurrence) ([]*sqlc.Image, error) {
	return r.Database.GetImagesForOccurrence(ctx, obj.ID)
}

func (r *occurrenceSideDishResolver) Occurrence(ctx context.Context, obj *sqlc.OccurrenceSideDish) (*sqlc.Occurrence, error) {
	return r.Database.GetOccurrenceByID(ctx, obj.Occurrence)
}

func (r *occurrenceSideDishResolver) Dish(ctx context.Context, obj *sqlc.OccurrenceSideDish) (*sqlc.Dish, error) {
	return r.Database.GetDishByID(ctx, obj.Dish)
}

func (r *occurrenceTagResolver) Occurrence(ctx context.Context, obj *sqlc.OccurrenceTag) (*sqlc.Occurrence, error) {
	return r.Database.GetOccurrenceByID(ctx, obj.Occurrence)
}

func (r *occurrenceTagResolver) Tag(ctx context.Context, obj *sqlc.OccurrenceTag) (*sqlc.Tag, error) {
	return r.Database.GetTagByKey(ctx, obj.Tag)
}

func (r *reviewResolver) Occurrence(ctx context.Context, obj *sqlc.Review) (*sqlc.Occurrence, error) {
	return r.Database.GetOccurrenceByID(ctx, obj.Occurrence)
}

// Dish returns gqlserver.DishResolver implementation.
func (r *Resolver) Dish() gqlserver.DishResolver { return &dishResolver{r} }

// Image returns gqlserver.ImageResolver implementation.
func (r *Resolver) Image() gqlserver.ImageResolver { return &imageResolver{r} }

// Occurrence returns gqlserver.OccurrenceResolver implementation.
func (r *Resolver) Occurrence() gqlserver.OccurrenceResolver { return &occurrenceResolver{r} }

// OccurrenceSideDish returns gqlserver.OccurrenceSideDishResolver implementation.
func (r *Resolver) OccurrenceSideDish() gqlserver.OccurrenceSideDishResolver {
	return &occurrenceSideDishResolver{r}
}

// OccurrenceTag returns gqlserver.OccurrenceTagResolver implementation.
func (r *Resolver) OccurrenceTag() gqlserver.OccurrenceTagResolver { return &occurrenceTagResolver{r} }

// Review returns gqlserver.ReviewResolver implementation.
func (r *Resolver) Review() gqlserver.ReviewResolver { return &reviewResolver{r} }

type dishResolver struct{ *Resolver }
type imageResolver struct{ *Resolver }
type occurrenceResolver struct{ *Resolver }
type occurrenceSideDishResolver struct{ *Resolver }
type occurrenceTagResolver struct{ *Resolver }
type reviewResolver struct{ *Resolver }
