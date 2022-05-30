package db

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/mensatt/backend/internal/db/sqlc"
)

type ExtendedQuerier interface {
	sqlc.Querier
	CreateOccurrenceWithSideDishesAndTags(ctx context.Context, occParams *sqlc.CreateOccurrenceParams, sideDishes []uuid.UUID, tags []string) (*sqlc.Occurrence, error)
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
