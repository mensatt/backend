package resolvers

import (
	"context"
	"github.com/google/uuid"
	ent "github.com/mensatt/backend/internal/database/ent"
	"github.com/mensatt/backend/internal/database/ent/dish"
	"github.com/mensatt/backend/internal/database/ent/location"
	"github.com/mensatt/backend/internal/database/ent/occurrence"
	"github.com/mensatt/backend/internal/database/ent/review"
	"github.com/mensatt/backend/internal/graphql/models"
	"io"
	"net/http"
)

//func (r *mutationResolver) storeImages(tx *ent.Tx, ctx context.Context, review *ent.Review, images []*models.ImageInput) ([]*ent.Image, error) {
//	var imageEntities []*ent.Image
//	for _, image := range images {
//		if image.Image.Size > int64(r.ImageUploader.GetMaxImageSize()) {
//			return nil, fmt.Errorf("image size is too large") // let the caller handle the transaction rollback
//		}
//
//		imageBytes, err := io.ReadAll(image.Image.File)
//		if err != nil {
//			return nil, err // let the caller handle the transaction rollback
//		}
//
//		imageUUID, imageHash, err := r.ImageUploader.ValidateAndStoreImage(imageBytes, image.Rotation)
//		if err != nil {
//			return nil, err // let the caller handle the transaction rollback
//		}
//
//		// on tx rollback remove image
//		tx.OnRollback(func(next ent.Rollbacker) ent.Rollbacker {
//			return ent.RollbackFunc(func(ctx context.Context, tx *ent.Tx) error {
//				err := r.ImageUploader.RemoveImage(imageUUID)
//				if err != nil {
//					return err
//				}
//
//				return next.Rollback(ctx, tx)
//			})
//		})
//
//		imageEntity, err := tx.Image.Create().
//			SetID(imageUUID).
//			SetImageHash(imageHash).
//			SetReview(review).
//			Save(ctx)
//		if err != nil {
//			return nil, err // let the caller handle the transaction rollback
//		}
//
//		imageEntities = append(imageEntities, imageEntity)
//	}
//
//	return imageEntities, nil
//}

func (r *mutationResolver) submitImages(uuids []uuid.UUID) []uuid.UUID {
	var submittedImages []uuid.UUID

	for _, imageUUID := range uuids {
		url := "http://localhost:3000/image/" + imageUUID.String()
		request, err := http.NewRequest("POST", url, nil)
		if err != nil {
			continue // ignore error and continue with the next image
		}

		request.Header.Add("Authorization", "Bearer "+r.ImageAPIKey)
		client := &http.Client{}

		response, err := client.Do(request)
		if err != nil {
			continue // ignore error and continue with the next image
		}
		defer response.Body.Close() // unhandled error

		if response.StatusCode != 200 {
			continue // ignore error and continue with the next image
		}

		body, err := io.ReadAll(response.Body)
		if err != nil {
			continue // ignore error and continue with the next image
		}

		uuid, err := uuid.Parse(string(body))
		if err != nil {
			continue // ignore error and continue with the next image
		}

		submittedImages = append(submittedImages, uuid)
	}

	return submittedImages
}

func (r *mutationResolver) approveImages(uuids []uuid.UUID) ([]uuid.UUID, error) {
	var approvedImages []uuid.UUID

	for _, imageUUID := range uuids {
		url := "http://localhost:3000/approve/" + imageUUID.String()
		request, err := http.NewRequest("POST", url, nil)
		if err != nil {
			return approvedImages, err
		}

		request.Header.Add("Authorization", "Bearer "+r.ImageAPIKey)
		client := &http.Client{}

		response, err := client.Do(request)
		if err != nil {
			return approvedImages, err
		}

		defer response.Body.Close() // unhandled error

		if response.StatusCode != 200 {
			return approvedImages, err
		}

		body, err := io.ReadAll(response.Body)
		if err != nil {
			return approvedImages, err
		}

		uuid, err := uuid.Parse(string(body))
		if err != nil {
			return approvedImages, err
		}

		approvedImages = append(approvedImages, uuid)
	}
	return approvedImages, nil
}

func (r *mutationResolver) unapproveImages(uuids []uuid.UUID) ([]uuid.UUID, error) {
	var unapprovedImages []uuid.UUID

	for _, imageUUID := range uuids {
		url := "http://localhost:3000/unapprove/" + imageUUID.String()
		request, err := http.NewRequest("POST", url, nil)
		if err != nil {
			return unapprovedImages, err
		}

		request.Header.Add("Authorization", "Bearer "+r.ImageAPIKey)
		client := &http.Client{}

		response, err := client.Do(request)
		if err != nil {
			return unapprovedImages, err
		}

		defer response.Body.Close() // unhandled error

		if response.StatusCode != 200 {
			return unapprovedImages, err
		}

		body, err := io.ReadAll(response.Body)
		if err != nil {
			return unapprovedImages, err
		}

		uuid, err := uuid.Parse(string(body))
		if err != nil {
			return unapprovedImages, err
		}

		unapprovedImages = append(unapprovedImages, uuid)
	}
	return unapprovedImages, nil
}

func (r *mutationResolver) deleteImages(ctx context.Context, images []*ent.Image) error {
	//for _, image := range images {
	//	err := r.ImageUploader.RemoveImage(image.ID)
	//	if err != nil {
	//		return err
	//	}
	//}
	//
	//return nil

	for _, image := range images {
		url := "http://localhost:3000/image/" + image.ID.String()
		request, err := http.NewRequest("DELETE", url, nil)
		if err != nil {
			return err // todo: maybe
		}

		// Add headers to the request
		request.Header.Add("Authorization", "Bearer "+r.ImageAPIKey)

		// Create a client
		client := &http.Client{}

		// Send the request
		response, err := client.Do(request)
		if err != nil {
			return err // todo: maybe
		}
		defer response.Body.Close() // unhandled error

		if response.StatusCode != 200 {
			return err // todo: maybe
		}
		err = r.Database.Image.DeleteOne(image).Exec(ctx)
		if err != nil {
			return err // todo: maybe
		}
	}
	return nil
}

func (r *queryResolver) filteredDishes(ctx context.Context, filter *models.DishFilter) ([]*ent.Dish, error) {
	queryBuilder := r.Database.Dish.Query()

	if filter.Dishes != nil {
		queryBuilder = queryBuilder.Where(dish.IDIn(filter.Dishes...))
	}

	if filter.NameDe != nil {
		queryBuilder = queryBuilder.Where(dish.NameDeEQ(*filter.NameDe))
	}

	if filter.NameEn != nil {
		queryBuilder = queryBuilder.Where(dish.NameEnEQ(*filter.NameEn))
	}

	return queryBuilder.All(ctx)
}

func (r *queryResolver) filteredOccurrences(ctx context.Context, filter *models.OccurrenceFilter) ([]*ent.Occurrence, error) {
	queryBuilder := r.Database.Occurrence.Query()

	if filter.Occurrences != nil {
		queryBuilder = queryBuilder.Where(occurrence.IDIn(filter.Occurrences...))
	}

	if filter.Dishes != nil {
		queryBuilder = queryBuilder.Where(occurrence.HasDishWith(dish.IDIn(filter.Dishes...)))
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

func (r *queryResolver) filteredLocations(ctx context.Context, filter *models.LocationFilter) ([]*ent.Location, error) {
	queryBuilder := r.Database.Location.Query()

	if filter.Ids != nil {
		queryBuilder = queryBuilder.Where(location.IDIn(filter.Ids...))
	}

	if filter.ExternalIds != nil {
		queryBuilder = queryBuilder.Where(location.ExternalIDIn(filter.ExternalIds...))
	}

	if filter.Names != nil {
		queryBuilder = queryBuilder.Where(location.NameIn(filter.Names...))
	}

	if filter.Visible != nil {
		queryBuilder = queryBuilder.Where(location.VisibleEQ(*filter.Visible))
	}

	return queryBuilder.All(ctx)
}
