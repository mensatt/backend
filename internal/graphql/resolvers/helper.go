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

func (r *mutationResolver) rotateImage(uuid uuid.UUID, angle int) error {
	url := r.ImageAPIURL + "rotate?id=" + uuid.String() + "&angle=" + string(rune(angle))
	request, err := http.NewRequest("POST", url, nil)
	if err != nil {
		return err
	}

	request.Header.Add("Authorization", "Bearer "+r.ImageAPIKey)
	client := &http.Client{}

	response, err := client.Do(request)
	if err != nil {
		return err
	}

	defer response.Body.Close() // unhandled error

	if response.StatusCode != 200 {
		return err
	}

	return nil
}

func (r *mutationResolver) submitImages(images []*models.ImageInput) []uuid.UUID {
	var submittedImages []uuid.UUID

	for _, image := range images {
		url := r.ImageAPIURL + "submit/" + image.ID.String()
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
		url := r.ImageAPIURL + "approve/" + imageUUID.String()
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
		url := r.ImageAPIURL + "unapprove/" + imageUUID.String()
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

func (r *mutationResolver) deleteImages(uuids []uuid.UUID) []uuid.UUID {
	var deletedImages []uuid.UUID

	for _, imageUUID := range uuids {
		url := r.ImageAPIURL + "image/" + imageUUID.String()
		request, err := http.NewRequest("DELETE", url, nil)
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

		deletedImages = append(deletedImages, uuid)
	}
	return deletedImages
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
