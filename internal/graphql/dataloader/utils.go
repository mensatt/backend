package dataloader

import (
	"context"
	"fmt"

	gophers_dl "github.com/graph-gophers/dataloader/v7"
)

func newBatchedLoaderFromDB[K comparable, V any](getKeyFunc func(*V) K, bulkDBFunc func(context.Context, []K) ([]*V, error)) *gophers_dl.Loader[K, *V] {
	batcher := func(ctx context.Context, keys []K) []*gophers_dl.Result[*V] {
		return getBulkFromDB(ctx, keys, getKeyFunc, bulkDBFunc)
	}
	return gophers_dl.NewBatchedLoader(batcher)
}

func getBulkFromDB[K comparable, V any](ctx context.Context, keys []K, getKeyFunc func(*V) K, bulkDBFunc func(context.Context, []K) ([]*V, error)) []*gophers_dl.Result[*V] {
	keyOrder := make(map[K]int, len(keys))
	for i, key := range keys {
		keyOrder[key] = i
	}

	results := make([]*gophers_dl.Result[*V], len(keys))

	dbRecords, err := bulkDBFunc(ctx, keys)

	if err != nil {
		for i := range results {
			results[i] = &gophers_dl.Result[*V]{
				Data:  nil,
				Error: err,
			}
		}
	} else {
		for _, record := range dbRecords {
			// if found, remove from index lookup map so we know elements were found
			key := getKeyFunc(record)
			if idx, ok := keyOrder[key]; ok {
				results[idx] = &gophers_dl.Result[*V]{
					Data:  record,
					Error: nil,
				}
				delete(keyOrder, key)
			}
		}
		// fill array positions with errors where not found in DB
		for recordKey, idx := range keyOrder {
			err := fmt.Errorf("record with key: '%v' not found", recordKey)
			results[idx] = &gophers_dl.Result[*V]{
				Data:  nil,
				Error: err,
			}
		}
	}

	return results
}
