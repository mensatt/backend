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

// Aliases is the resolver for the aliases field.
func (r *dishResolver) Aliases(ctx context.Context, obj *sqlc.Dish) ([]string, error) {
	return r.Database.GetAliasesForDish(ctx, obj.ID)
}

// ReviewData is the resolver for the reviewData field.
func (r *dishResolver) ReviewData(ctx context.Context, obj *sqlc.Dish, filter *models.ReviewFilter) (*models.ReviewDataDish, error) {
	return &models.ReviewDataDish{
		DishID: obj.ID,
		Filter: filter,
	}, nil
}

// Review is the resolver for the review field.
func (r *imageResolver) Review(ctx context.Context, obj *sqlc.Image) (*sqlc.Review, error) {
	return r.Database.GetReviewByImage(ctx, obj.ID)
}

// ImageURL is the resolver for the imageUrl field.
func (r *imageResolver) ImageURL(ctx context.Context, obj *sqlc.Image) (string, error) {
	return fmt.Sprintf("%s/%s", r.ImageBaseURL, obj.ImageHash), nil
}

// Location is the resolver for the location field.
func (r *occurrenceResolver) Location(ctx context.Context, obj *sqlc.Occurrence) (*sqlc.Location, error) {
	return r.Database.GetLocationByID(ctx, obj.Location)
}

// Dish is the resolver for the dish field.
func (r *occurrenceResolver) Dish(ctx context.Context, obj *sqlc.Occurrence) (*sqlc.Dish, error) {
	return r.Database.GetDishByID(ctx, obj.Dish)
}

// SideDishes is the resolver for the sideDishes field.
func (r *occurrenceResolver) SideDishes(ctx context.Context, obj *sqlc.Occurrence) ([]*sqlc.Dish, error) {
	return r.Database.GetSideDishesForOccurrence(ctx, obj.ID)
}

// Tags is the resolver for the tags field.
func (r *occurrenceResolver) Tags(ctx context.Context, obj *sqlc.Occurrence) ([]*sqlc.Tag, error) {
	return r.Database.GetTagsForOccurrence(ctx, obj.ID)
}

// ReviewData is the resolver for the reviewData field.
func (r *occurrenceResolver) ReviewData(ctx context.Context, obj *sqlc.Occurrence, filter *models.ReviewFilter) (*models.ReviewDataOccurrence, error) {
	return &models.ReviewDataOccurrence{
		OccurrenceID: obj.ID,
		Filter:       filter,
	}, nil
}

// Occurrence is the resolver for the occurrence field.
func (r *occurrenceSideDishResolver) Occurrence(ctx context.Context, obj *sqlc.OccurrenceSideDish) (*sqlc.Occurrence, error) {
	return r.Database.GetOccurrenceByID(ctx, obj.Occurrence)
}

// Dish is the resolver for the dish field.
func (r *occurrenceSideDishResolver) Dish(ctx context.Context, obj *sqlc.OccurrenceSideDish) (*sqlc.Dish, error) {
	return r.Database.GetDishByID(ctx, obj.Dish)
}

// Occurrence is the resolver for the occurrence field.
func (r *occurrenceTagResolver) Occurrence(ctx context.Context, obj *sqlc.OccurrenceTag) (*sqlc.Occurrence, error) {
	return r.Database.GetOccurrenceByID(ctx, obj.Occurrence)
}

// Tag is the resolver for the tag field.
func (r *occurrenceTagResolver) Tag(ctx context.Context, obj *sqlc.OccurrenceTag) (*sqlc.Tag, error) {
	return r.Database.GetTagByKey(ctx, obj.Tag)
}

// Occurrence is the resolver for the occurrence field.
func (r *reviewResolver) Occurrence(ctx context.Context, obj *sqlc.Review) (*sqlc.Occurrence, error) {
	return r.Database.GetOccurrenceByID(ctx, obj.Occurrence)
}

// Images is the resolver for the images field.
func (r *reviewResolver) Images(ctx context.Context, obj *sqlc.Review) ([]*sqlc.Image, error) {
	return r.Database.GetImagesByReview(ctx, obj.ID)
}

// Reviews is the resolver for the reviews field.
func (r *reviewDataDishResolver) Reviews(ctx context.Context, obj *models.ReviewDataDish) ([]*sqlc.Review, error) {
	param := sqlc.GetDishReviewsParams{
		ID: obj.DishID,
	}
	if obj.Filter != nil {
		param.Approved = boolPointerToNullBool(obj.Filter.Approved)
	}
	return r.Database.GetDishReviews(ctx, &param)
}

// Images is the resolver for the images field.
func (r *reviewDataDishResolver) Images(ctx context.Context, obj *models.ReviewDataDish) ([]*sqlc.Image, error) {
	return r.Database.GetImagesByDish(ctx, obj.DishID)
}

// Metadata is the resolver for the metadata field.
func (r *reviewDataDishResolver) Metadata(ctx context.Context, obj *models.ReviewDataDish) (*sqlc.GetDishReviewMetadataRow, error) {
	param := sqlc.GetDishReviewMetadataParams{
		ID: obj.DishID,
	}
	if obj.Filter != nil {
		param.Approved = boolPointerToNullBool(obj.Filter.Approved)
	}
	return r.Database.GetDishReviewMetadata(ctx, &param)
}

// Reviews is the resolver for the reviews field.
func (r *reviewDataOccurrenceResolver) Reviews(ctx context.Context, obj *models.ReviewDataOccurrence) ([]*sqlc.Review, error) {
	param := sqlc.GetOccurrenceReviewsParams{
		ID: obj.OccurrenceID,
	}
	if obj.Filter != nil {
		param.Approved = boolPointerToNullBool(obj.Filter.Approved)
	}
	return r.Database.GetOccurrenceReviews(ctx, &param)
}

// Images is the resolver for the images field.
func (r *reviewDataOccurrenceResolver) Images(ctx context.Context, obj *models.ReviewDataOccurrence) ([]*sqlc.Image, error) {
	return r.Database.GetImagesByOccurrence(ctx, obj.OccurrenceID)
}

// Metadata is the resolver for the metadata field.
func (r *reviewDataOccurrenceResolver) Metadata(ctx context.Context, obj *models.ReviewDataOccurrence) (*sqlc.GetOccurrenceReviewMetadataRow, error) {
	param := sqlc.GetOccurrenceReviewMetadataParams{
		Occurrence: obj.OccurrenceID,
	}
	if obj.Filter != nil {
		param.Approved = boolPointerToNullBool(obj.Filter.Approved)
	}
	return r.Database.GetOccurrenceReviewMetadata(ctx, &param)
}

// AverageStars is the resolver for the averageStars field.
func (r *reviewMetadataDishResolver) AverageStars(ctx context.Context, obj *sqlc.GetDishReviewMetadataRow) (*float64, error) {
	return pgtypeNumericToFloat(obj.AverageStars)
}

// AverageStars is the resolver for the averageStars field.
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
