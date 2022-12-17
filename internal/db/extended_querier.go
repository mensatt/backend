package db

import (
	"context"
	"errors"
	"github.com/getsentry/sentry-go"
	"github.com/google/uuid"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/mensatt/backend/internal/db/sqlc"
)

type ExtendedQuerier interface {
	sqlc.Querier
	CreateOccurrenceWithSideDishesAndTags(ctx context.Context, occParams *sqlc.CreateOccurrenceParams, sideDishes []uuid.UUID, tags []string) (*sqlc.Occurrence, error)
	CreateReviewWithImages(ctx context.Context, reviewParams *CreateReviewWithImagesParams) (*sqlc.Review, error)
	MergeDishes(ctx context.Context, keep *sqlc.Dish, merge *sqlc.Dish) (*sqlc.Dish, error)
}

var _ ExtendedQuerier = (*ExtendedQueries)(nil)

type ExtendedQueries struct {
	*sqlc.Queries
	pool *pgxpool.Pool
}

func NewExtended(pool *pgxpool.Pool) *ExtendedQueries {
	return &ExtendedQueries{
		Queries: sqlc.New(pool),
		pool:    pool,
	}
}

func (eq *ExtendedQueries) CreateOccurrenceWithSideDishesAndTags(ctx context.Context, occParams *sqlc.CreateOccurrenceParams, sideDishes []uuid.UUID, tags []string) (*sqlc.Occurrence, error) {
	tx, err := eq.pool.Begin(ctx)
	if err != nil {
		return nil, err
	}

	defer func() {
		if err != nil {
			err := tx.Rollback(ctx)
			if err != nil {
				sentry.CaptureException(err)
				return
			}
		} else {
			err := tx.Commit(ctx)
			if err != nil {
				sentry.CaptureException(err)
				return
			}
		}
	}()

	qtx := eq.WithTx(tx)

	occ, err := qtx.CreateOccurrence(ctx, occParams)
	if err != nil {
		return nil, err
	}

	addSideDishesParams := make([]*sqlc.AddMultipleOccurrenceSideDishesParams, len(sideDishes))
	addTagsParams := make([]*sqlc.AddMultipleOccurrenceTagsParams, len(tags))

	for i, dishID := range sideDishes {
		addSideDishesParams[i] = &sqlc.AddMultipleOccurrenceSideDishesParams{
			Occurrence: occ.ID,
			Dish:       dishID,
		}
	}
	for i, tagKey := range tags {
		addTagsParams[i] = &sqlc.AddMultipleOccurrenceTagsParams{
			Occurrence: occ.ID,
			Tag:        tagKey,
		}
	}

	_, err = qtx.AddMultipleOccurrenceSideDishes(ctx, addSideDishesParams)
	if err != nil {
		return nil, err
	}
	_, err = qtx.AddMultipleOccurrenceTags(ctx, addTagsParams)
	if err != nil {
		return nil, err
	}

	return occ, nil
}

type CreateImageParams struct {
	ImageHash string `json:"image_hash"`
}

type CreateReviewWithImagesParams struct {
	sqlc.CreateReviewParams
	Images []*CreateImageParams
}

func (eq *ExtendedQueries) CreateReviewWithImages(ctx context.Context, reviewParams *CreateReviewWithImagesParams) (*sqlc.Review, error) {
	tx, err := eq.pool.Begin(ctx)
	if err != nil {
		return nil, err
	}

	defer func() {
		if err != nil {
			err := tx.Rollback(ctx)
			if err != nil {
				sentry.CaptureException(err)
				return
			}
		} else {
			err := tx.Commit(ctx)
			if err != nil {
				sentry.CaptureException(err)
				return
			}
		}
	}()

	qtx := eq.WithTx(tx)

	review, err := qtx.CreateReview(ctx, &reviewParams.CreateReviewParams)
	if err != nil {
		return nil, err
	}

	addImagesToReviewParams := make([]*sqlc.AddMultipleImagesToReviewParams, len(reviewParams.Images))

	for i, imageParam := range reviewParams.Images {
		addImagesToReviewParams[i] = &sqlc.AddMultipleImagesToReviewParams{
			Review:    review.ID,
			ImageHash: imageParam.ImageHash,
		}
	}

	_, err = qtx.AddMultipleImagesToReview(ctx, addImagesToReviewParams)
	if err != nil {
		return nil, err
	}

	return review, nil
}

func (eq *ExtendedQueries) MergeDishes(ctx context.Context, keep *sqlc.Dish, merge *sqlc.Dish) (*sqlc.Dish, error) {
	tx, err := eq.pool.Begin(ctx)
	if err != nil {
		return nil, err
	}

	defer func() {
		if err != nil {
			err := tx.Rollback(ctx)
			if err != nil {
				sentry.CaptureException(err)
				return
			}
		} else {
			err := tx.Commit(ctx)
			if err != nil {
				sentry.CaptureException(err)
				return
			}
		}
	}()

	qtx := eq.WithTx(tx)

	// re-write all existing dish_aliases to point to the keep dish instead of the merge dish
	mergeAliasesParams := &sqlc.MergeAliasesParams{
		Keep:  keep.ID,
		Merge: merge.ID,
	}
	err = qtx.MergeAliases(ctx, mergeAliasesParams)
	if err != nil {
		return nil, err
	}

	// if keep and merge dishes are on the same day, we need to update the reviews to point to the keep dish
	mergeReviewsParams := &sqlc.MergeReviewsParams{
		Keep:  keep.ID,
		Merge: merge.ID,
	}
	err = qtx.MergeReviews(ctx, mergeReviewsParams)
	if err != nil {
		return nil, err
	}

	// delete all occurrences of the merge dish
	mergeSameDayOccurrencesParams := &sqlc.MergeSameDayOccurrencesParams{
		Keep:  keep.ID,
		Merge: merge.ID,
	}
	err = qtx.MergeSameDayOccurrences(ctx, mergeSameDayOccurrencesParams)
	if err != nil {
		return nil, err
	}

	// update all occurrences to point to the keep dish instead of the merge dish
	mergeOccurrencesParams := &sqlc.MergeOccurrencesParams{
		Keep:  keep.ID,
		Merge: merge.ID,
	}
	err = qtx.MergeOccurrences(ctx, mergeOccurrencesParams)
	if err != nil {
		return nil, err
	}

	// update the side dishes - this can result in a violation of the unique constraint as some side dishes might already exist which should be ignored
	mergeSideDishesParams := &sqlc.MergeSideDishesParams{
		Keep:  keep.ID,
		Merge: merge.ID,
	}
	err = qtx.MergeSideDishes(ctx, mergeSideDishesParams)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code != "23505" {
				return nil, err // not a unique constraint violation - return the error
			}
		}
	}

	// delete the old (to merge) dish as it is not referenced anymore
	_, err = qtx.DeleteDish(ctx, merge.ID)
	if err != nil {
		return nil, err
	}

	return keep, nil
}
