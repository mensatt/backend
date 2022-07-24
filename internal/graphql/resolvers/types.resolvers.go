package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/mensatt/backend/internal/db/sqlc"
	"github.com/mensatt/backend/internal/graphql/gqlserver"
	"github.com/mensatt/backend/internal/graphql/models"
)

func (r *dishResolver) Aliases(ctx context.Context, obj *sqlc.Dish) ([]string, error) {
	return r.Database.GetAliasesForDish(ctx, obj.ID)
}

func (r *dishResolver) ReviewData(ctx context.Context, obj *sqlc.Dish) (*models.ReviewDataDish, error) {
	return &models.ReviewDataDish{
		DishID: obj.ID,
	}, nil
}

func (r *imageResolver) Review(ctx context.Context, obj *sqlc.Image) (*sqlc.Review, error) {
	return r.Database.GetReviewByImage(ctx, obj.ID)
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

func (r *occurrenceResolver) ReviewData(ctx context.Context, obj *sqlc.Occurrence) (*models.ReviewDataOccurrence, error) {
	return &models.ReviewDataOccurrence{
		OccurrenceID: obj.ID,
	}, nil
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

func (r *reviewResolver) Images(ctx context.Context, obj *sqlc.Review) ([]*sqlc.Image, error) {
	return r.Database.GetImagesByReview(ctx, obj.ID)
}

func (r *reviewDataDishResolver) Reviews(ctx context.Context, obj *models.ReviewDataDish) ([]*sqlc.Review, error) {
	return r.Database.GetReviewsForOccurrence(ctx, obj.DishID)
}

func (r *reviewDataDishResolver) Metadata(ctx context.Context, obj *models.ReviewDataDish) (*sqlc.GetDishReviewMetadataRow, error) {
	return r.Database.GetDishReviewMetadata(ctx, obj.DishID)
}

func (r *reviewDataOccurrenceResolver) Reviews(ctx context.Context, obj *models.ReviewDataOccurrence) ([]*sqlc.Review, error) {
	return r.Database.GetReviewsForOccurrence(ctx, obj.OccurrenceID)
}

func (r *reviewDataOccurrenceResolver) Metadata(ctx context.Context, obj *models.ReviewDataOccurrence) (*sqlc.GetOccurrenceReviewMetadataRow, error) {
	return r.Database.GetOccurrenceReviewMetadata(ctx, obj.OccurrenceID)
}

func (r *reviewMetadataDishResolver) AverageStars(ctx context.Context, obj *sqlc.GetDishReviewMetadataRow) (*float64, error) {
	return pgtypeNumericToFloat(obj.AverageStars)
}

func (r *reviewMetadataOccurrenceResolver) AverageStars(ctx context.Context, obj *sqlc.GetOccurrenceReviewMetadataRow) (*float64, error) {
	return pgtypeNumericToFloat(obj.AverageStars)
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

// ReviewDataDish returns gqlserver.ReviewDataDishResolver implementation.
func (r *Resolver) ReviewDataDish() gqlserver.ReviewDataDishResolver {
	return &reviewDataDishResolver{r}
}

// ReviewDataOccurrence returns gqlserver.ReviewDataOccurrenceResolver implementation.
func (r *Resolver) ReviewDataOccurrence() gqlserver.ReviewDataOccurrenceResolver {
	return &reviewDataOccurrenceResolver{r}
}

// ReviewMetadataDish returns gqlserver.ReviewMetadataDishResolver implementation.
func (r *Resolver) ReviewMetadataDish() gqlserver.ReviewMetadataDishResolver {
	return &reviewMetadataDishResolver{r}
}

// ReviewMetadataOccurrence returns gqlserver.ReviewMetadataOccurrenceResolver implementation.
func (r *Resolver) ReviewMetadataOccurrence() gqlserver.ReviewMetadataOccurrenceResolver {
	return &reviewMetadataOccurrenceResolver{r}
}

type dishResolver struct{ *Resolver }
type imageResolver struct{ *Resolver }
type occurrenceResolver struct{ *Resolver }
type occurrenceSideDishResolver struct{ *Resolver }
type occurrenceTagResolver struct{ *Resolver }
type reviewResolver struct{ *Resolver }
type reviewDataDishResolver struct{ *Resolver }
type reviewDataOccurrenceResolver struct{ *Resolver }
type reviewMetadataDishResolver struct{ *Resolver }
type reviewMetadataOccurrenceResolver struct{ *Resolver }
