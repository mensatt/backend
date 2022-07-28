package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/mensatt/backend/internal/db"
	"github.com/mensatt/backend/internal/db/sqlc"
	"github.com/mensatt/backend/internal/graphql/gqlserver"
	"github.com/mensatt/backend/internal/graphql/models"
)

// Images is the resolver for the images field.
func (r *createReviewInputResolver) Images(ctx context.Context, obj *db.CreateReviewWithImagesParams, data []*models.ImageInput) error {
	if len(data) == 0 {
		return nil
	}

	imagesParams := make([]*db.CreateImageParams, len(data))
	for i, image := range data {
		imageHash, err := r.transcodeAndStoreImage(ctx, image)
		if err != nil {
			// TODO: delete previous images
			return err
		}
		imagesParams[i] = &db.CreateImageParams{
			ImageHash: imageHash,
		}
	}
	obj.Images = imagesParams

	return nil
}

// Approved is the resolver for the approved field.
func (r *updateReviewInputResolver) Approved(ctx context.Context, obj *sqlc.UpdateReviewParams, data *bool) error {
	obj.AcceptedAt = approvedBoolToNullTime(data == nil && *data)
	return nil
}

// CreateReviewInput returns gqlserver.CreateReviewInputResolver implementation.
func (r *Resolver) CreateReviewInput() gqlserver.CreateReviewInputResolver {
	return &createReviewInputResolver{r}
}

// UpdateReviewInput returns gqlserver.UpdateReviewInputResolver implementation.
func (r *Resolver) UpdateReviewInput() gqlserver.UpdateReviewInputResolver {
	return &updateReviewInputResolver{r}
}

type createReviewInputResolver struct{ *Resolver }
type updateReviewInputResolver struct{ *Resolver }
