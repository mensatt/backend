package scalars

import (
	"database/sql"

	"github.com/99designs/gqlgen/graphql"
)

// MarshalNullTime is a custom marshaller.
func MarshalNullTime(ns sql.NullTime) graphql.Marshaler {
	if !ns.Valid {
		return graphql.Null
	}
	return graphql.MarshalTime(ns.Time)
}

// UnmarshalNullTime is a custom unmarshaller.
func UnmarshalNullTime(v interface{}) (sql.NullTime, error) {
	if v == nil {
		return sql.NullTime{Valid: false}, nil
	}
	t, err := graphql.UnmarshalTime(v)
	return sql.NullTime{Time: t, Valid: true}, err
}
