package resolvers

import (
	"context"
	ent "github.com/mensatt/backend/internal/database/ent"
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

//func approvedBoolToNullTime(approved bool) sql.NullTime {
//	if approved {
//		return sql.NullTime{Valid: true, Time: time.Now()}
//	}
//	return sql.NullTime{Valid: false}
//}
//
//func pgtypeNumericToFloat(numeric pgtype.Numeric) (*float64, error) {
//	if numeric.Status == pgtype.Null {
//		return nil, nil
//	}
//
//	var f float64
//	err := numeric.AssignTo(&f)
//	return &f, err
//}
//
//func boolPointerToNullBool(b *bool) sql.NullBool {
//	if b == nil {
//		return sql.NullBool{Valid: false}
//	}
//	return sql.NullBool{Valid: true, Bool: *b}
//}

func (r *queryResolver) filteredOccurrences(ctx context.Context, filter *models.OccurrenceFilter) ([]*ent.Occurrence, error) {
	queryBuilder := r.Database.Occurrence.Query()

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
		// todo: implement location filter
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
