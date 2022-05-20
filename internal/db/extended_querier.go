package db

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
)

type ExtendedQuerier interface {
	Querier
	CreateOccurrenceWithSideDishesAndTags(ctx context.Context, occParams *CreateOccurrenceParams, sideDishes []uuid.UUID, tags []string) (*Occurrence, error)
}

var _ ExtendedQuerier = (*ExtendedQueries)(nil)

type ExtendedQueries struct {
	Queries
	pool *pgxpool.Pool
}

func NewExtended(pool *pgxpool.Pool) *ExtendedQueries {
	return &ExtendedQueries{
		Queries: Queries{
			db: pool,
		},
		pool: pool,
	}
}

func (eq *ExtendedQueries) CreateOccurrenceWithSideDishesAndTags(ctx context.Context, occParams *CreateOccurrenceParams, sideDishes []uuid.UUID, tags []string) (*Occurrence, error) {
	tx, err := eq.pool.Begin(ctx)
	if err != nil {
		return nil, err
	}

	defer func() {
        if err != nil {
            tx.Rollback(ctx)
        } else {
            tx.Commit(ctx)
        }
    }()

	qtx := eq.WithTx(tx)

	occ, err := qtx.CreateOccurrence(ctx, occParams)
	if err != nil {
		return nil, err
	}

	var addSideDishesParams []*AddMultipleOccurrenceSideDishesParams
	var addTagsParams []*AddMultipleOccurrenceTagsParams

	for _, dishID := range sideDishes {
		addSideDishesParams = append(addSideDishesParams, &AddMultipleOccurrenceSideDishesParams{
			OccurrenceID: occ.ID,
			DishID:       dishID,
		})
	}
	for _, tagKey := range tags {
		addTagsParams = append(addTagsParams, &AddMultipleOccurrenceTagsParams{
			OccurrenceID: occ.ID,
			TagKey:       tagKey,
		})
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
