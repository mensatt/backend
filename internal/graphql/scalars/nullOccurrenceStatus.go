package scalars

import (
	"github.com/99designs/gqlgen/graphql"
	"github.com/mensatt/backend/internal/db/sqlc"
)

func MarshalNullOccurrenceStatus(v sqlc.NullOccurrenceStatus) graphql.Marshaler {
	if !v.Valid {
		return graphql.Null
	}
	return MarshalOccurrenceStatus(v.OccurrenceStatus)
}

func UnmarshalNullOccurrenceStatus(v interface{}) (sqlc.NullOccurrenceStatus, error) {
	if v == nil {
		return sqlc.NullOccurrenceStatus{Valid: false}, nil
	}
	r, err := UnmarshalOccurrenceStatus(v)
	return sqlc.NullOccurrenceStatus{OccurrenceStatus: r, Valid: true}, err
}
