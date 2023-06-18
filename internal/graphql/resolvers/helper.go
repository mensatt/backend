package resolvers

import (
	"context"
	ent "github.com/mensatt/backend/internal/database/ent"
	"github.com/mensatt/backend/internal/database/ent/dish"
	"github.com/mensatt/backend/internal/database/ent/location"
	"github.com/mensatt/backend/internal/database/ent/occurrence"
	"github.com/mensatt/backend/internal/database/ent/review"
	"github.com/mensatt/backend/internal/graphql/models"
	"io"
)

func (r *mutationResolver) storeImages(ctx context.Context, review *ent.Review, images []*models.ImageInput) ([]*ent.Image, error) {
	var imageEntities []*ent.Image
	for _, image := range images {
		if image.Image.Size > int64(r.ImageUploader.GetMaxImageSize()) {
			continue // todo: perhaps error?
		}

		imageBytes, err := io.ReadAll(image.Image.File)
		if err != nil {
			return nil, err
		}

		imageUUID, imageHash, err := r.ImageUploader.ValidateAndStoreImage(imageBytes)
		if err != nil {
			return nil, err
		}

		imageEntity, err := r.Database.Image.Create().
			SetID(imageUUID).
			SetImageHash(imageHash).
			SetReview(review).
			Save(ctx)
		if err != nil {
			return nil, err
		}

		imageEntities = append(imageEntities, imageEntity)
	}

	return imageEntities, nil
}

func (r *mutationResolver) deleteImages(ctx context.Context, images []*ent.Image) error {
	for _, image := range images {
		err := r.ImageUploader.RemoveImage(image.ID)
		if err != nil {
			return err
		}
	}

	return nil
}

func (r *queryResolver) filteredOccurrences(ctx context.Context, filter *models.OccurrenceFilter) ([]*ent.Occurrence, error) {
	queryBuilder := r.Database.Occurrence.Query()

	if filter.Occurrences != nil {
		queryBuilder = queryBuilder.Where(occurrence.IDIn(filter.Occurrences...))
	}

	if filter.Dishes != nil {
		queryBuilder = queryBuilder.Where(occurrence.HasDishWith(dish.IDIn(filter.Dishes...)))
	}

	if filter.Status != nil {
		queryBuilder = queryBuilder.Where(occurrence.StatusEQ(*filter.Status))
	}

	if filter.StartDate != nil {
		queryBuilder = queryBuilder.Where(occurrence.DateGTE(*filter.StartDate))
	}

	if filter.EndDate != nil {
		queryBuilder = queryBuilder.Where(occurrence.DateLTE(*filter.EndDate))
	}

	if filter.Location != nil {
		queryBuilder = queryBuilder.Where(occurrence.HasLocationWith(location.IDEQ(*filter.Location)))
	}

	return queryBuilder.All(ctx)
}

func (r *queryResolver) filteredReviews(ctx context.Context, filter *models.ReviewFilter) ([]*ent.Review, error) {
	queryBuilder := r.Database.Review.Query()

	if filter.Approved != nil {
		if *filter.Approved {
			queryBuilder = queryBuilder.Where(review.AcceptedAtNotNil())
		} else {
			queryBuilder = queryBuilder.Where(review.AcceptedAtIsNil())
		}
	}

	return queryBuilder.All(ctx)
}
