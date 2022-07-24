// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.14.0
// source: copyfrom.go

package sqlc

import (
	"context"
)

// iteratorForAddMultipleImagesToReview implements pgx.CopyFromSource.
type iteratorForAddMultipleImagesToReview struct {
	rows                 []*AddMultipleImagesToReviewParams
	skippedFirstNextCall bool
}

func (r *iteratorForAddMultipleImagesToReview) Next() bool {
	if len(r.rows) == 0 {
		return false
	}
	if !r.skippedFirstNextCall {
		r.skippedFirstNextCall = true
		return true
	}
	r.rows = r.rows[1:]
	return len(r.rows) > 0
}

func (r iteratorForAddMultipleImagesToReview) Values() ([]interface{}, error) {
	return []interface{}{
		r.rows[0].Review,
		r.rows[0].ImageStoreID,
	}, nil
}

func (r iteratorForAddMultipleImagesToReview) Err() error {
	return nil
}

func (q *Queries) AddMultipleImagesToReview(ctx context.Context, arg []*AddMultipleImagesToReviewParams) (int64, error) {
	return q.db.CopyFrom(ctx, []string{"image"}, []string{"review", "image_store_id"}, &iteratorForAddMultipleImagesToReview{rows: arg})
}

// iteratorForAddMultipleOccurrenceSideDishes implements pgx.CopyFromSource.
type iteratorForAddMultipleOccurrenceSideDishes struct {
	rows                 []*AddMultipleOccurrenceSideDishesParams
	skippedFirstNextCall bool
}

func (r *iteratorForAddMultipleOccurrenceSideDishes) Next() bool {
	if len(r.rows) == 0 {
		return false
	}
	if !r.skippedFirstNextCall {
		r.skippedFirstNextCall = true
		return true
	}
	r.rows = r.rows[1:]
	return len(r.rows) > 0
}

func (r iteratorForAddMultipleOccurrenceSideDishes) Values() ([]interface{}, error) {
	return []interface{}{
		r.rows[0].Occurrence,
		r.rows[0].Dish,
	}, nil
}

func (r iteratorForAddMultipleOccurrenceSideDishes) Err() error {
	return nil
}

func (q *Queries) AddMultipleOccurrenceSideDishes(ctx context.Context, arg []*AddMultipleOccurrenceSideDishesParams) (int64, error) {
	return q.db.CopyFrom(ctx, []string{"occurrence_side_dishes"}, []string{"occurrence", "dish"}, &iteratorForAddMultipleOccurrenceSideDishes{rows: arg})
}

// iteratorForAddMultipleOccurrenceTags implements pgx.CopyFromSource.
type iteratorForAddMultipleOccurrenceTags struct {
	rows                 []*AddMultipleOccurrenceTagsParams
	skippedFirstNextCall bool
}

func (r *iteratorForAddMultipleOccurrenceTags) Next() bool {
	if len(r.rows) == 0 {
		return false
	}
	if !r.skippedFirstNextCall {
		r.skippedFirstNextCall = true
		return true
	}
	r.rows = r.rows[1:]
	return len(r.rows) > 0
}

func (r iteratorForAddMultipleOccurrenceTags) Values() ([]interface{}, error) {
	return []interface{}{
		r.rows[0].Occurrence,
		r.rows[0].Tag,
	}, nil
}

func (r iteratorForAddMultipleOccurrenceTags) Err() error {
	return nil
}

func (q *Queries) AddMultipleOccurrenceTags(ctx context.Context, arg []*AddMultipleOccurrenceTagsParams) (int64, error) {
	return q.db.CopyFrom(ctx, []string{"occurrence_tag"}, []string{"occurrence", "tag"}, &iteratorForAddMultipleOccurrenceTags{rows: arg})
}
