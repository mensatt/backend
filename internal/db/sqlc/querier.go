// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.14.0

package sqlc

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type Querier interface {
	AddImageToReview(ctx context.Context, arg *AddImageToReviewParams) (*Image, error)
	AddMultipleImagesToReview(ctx context.Context, arg []*AddMultipleImagesToReviewParams) (int64, error)
	AddMultipleOccurrenceSideDishes(ctx context.Context, arg []*AddMultipleOccurrenceSideDishesParams) (int64, error)
	AddMultipleOccurrenceTags(ctx context.Context, arg []*AddMultipleOccurrenceTagsParams) (int64, error)
	AddOccurrenceSideDish(ctx context.Context, arg *AddOccurrenceSideDishParams) (*OccurrenceSideDish, error)
	AddOccurrenceTag(ctx context.Context, arg *AddOccurrenceTagParams) (*OccurrenceTag, error)
	CreateDish(ctx context.Context, arg *CreateDishParams) (*Dish, error)
	CreateDishAlias(ctx context.Context, arg *CreateDishAliasParams) (*DishAlias, error)
	CreateOccurrence(ctx context.Context, arg *CreateOccurrenceParams) (*Occurrence, error)
	CreateReview(ctx context.Context, arg *CreateReviewParams) (*Review, error)
	CreateTag(ctx context.Context, arg *CreateTagParams) (*Tag, error)
	DeleteDishAlias(ctx context.Context, aliasName string) (*DishAlias, error)
	DeleteImage(ctx context.Context, id uuid.UUID) (*Image, error)
	DeleteOccurrence(ctx context.Context, id uuid.UUID) (*Occurrence, error)
	DeleteReview(ctx context.Context, id uuid.UUID) (*Review, error)
	GetAliasesForDish(ctx context.Context, dish uuid.UUID) ([]string, error)
	GetAllAliases(ctx context.Context) ([]*DishAlias, error)
	GetAllDishes(ctx context.Context) ([]*Dish, error)
	GetAllImages(ctx context.Context) ([]*Image, error)
	GetAllLocations(ctx context.Context) ([]*Location, error)
	GetAllOccurrences(ctx context.Context) ([]*Occurrence, error)
	GetAllReviews(ctx context.Context) ([]*Review, error)
	GetAllTags(ctx context.Context) ([]*Tag, error)
	GetDishByID(ctx context.Context, id uuid.UUID) (*Dish, error)
	GetDishReviewMetadata(ctx context.Context, id uuid.UUID) (*GetDishReviewMetadataRow, error)
	GetFilteredOccurrences(ctx context.Context, arg *GetFilteredOccurrencesParams) ([]*Occurrence, error)
	GetImageByID(ctx context.Context, id uuid.UUID) (*Image, error)
	GetImageStoreIDByID(ctx context.Context, id uuid.UUID) (string, error)
	GetImagesByDish(ctx context.Context, id uuid.UUID) ([]*Image, error)
	GetImagesByOccurrence(ctx context.Context, id uuid.UUID) ([]*Image, error)
	GetImagesByReview(ctx context.Context, id uuid.UUID) ([]*Image, error)
	GetImagesForOccurrence(ctx context.Context, id uuid.UUID) ([]*Image, error)
	GetLocationByID(ctx context.Context, id uuid.UUID) (*Location, error)
	GetOccurrenceByID(ctx context.Context, id uuid.UUID) (*Occurrence, error)
	GetOccurrenceReviewMetadata(ctx context.Context, occurrence uuid.UUID) (*GetOccurrenceReviewMetadataRow, error)
	GetOccurrencesAfterInclusiveDate(ctx context.Context, date time.Time) ([]*Occurrence, error)
	GetOccurrencesByDate(ctx context.Context, date time.Time) ([]*Occurrence, error)
	GetReviewByID(ctx context.Context, id uuid.UUID) (*Review, error)
	GetReviewByImage(ctx context.Context, id uuid.UUID) (*Review, error)
	GetReviewsByDish(ctx context.Context, id uuid.UUID) ([]*Review, error)
	GetReviewsForOccurrence(ctx context.Context, id uuid.UUID) ([]*Review, error)
	GetSideDishesForOccurrence(ctx context.Context, occurrence uuid.UUID) ([]*Dish, error)
	GetTagByKey(ctx context.Context, key string) (*Tag, error)
	GetTagsForOccurrence(ctx context.Context, occurrence uuid.UUID) ([]*Tag, error)
	GetUserByEmail(ctx context.Context, email string) (*User, error)
	GetUserByID(ctx context.Context, id uuid.UUID) (*User, error)
	RemoveOccurrenceSideDish(ctx context.Context, arg *RemoveOccurrenceSideDishParams) (*OccurrenceSideDish, error)
	RemoveOccurrenceTag(ctx context.Context, arg *RemoveOccurrenceTagParams) (*OccurrenceTag, error)
	SetReviewApproval(ctx context.Context, arg *SetReviewApprovalParams) (*Review, error)
	UpdateDish(ctx context.Context, arg *UpdateDishParams) (*Dish, error)
	UpdateDishAlias(ctx context.Context, arg *UpdateDishAliasParams) (*DishAlias, error)
	UpdateOccurrence(ctx context.Context, arg *UpdateOccurrenceParams) (*Occurrence, error)
	UpdateReview(ctx context.Context, arg *UpdateReviewParams) (*Review, error)
}

var _ Querier = (*Queries)(nil)
