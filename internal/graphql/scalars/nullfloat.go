package scalars

import (
	"database/sql"

	"github.com/99designs/gqlgen/graphql"
)

// MarshalNullFloat is a custom marshaller.
func MarshalNullFloat(ns sql.NullFloat64) graphql.Marshaler {
	if !ns.Valid {
		return graphql.Null
	}
	return graphql.MarshalFloat(ns.Float64)
}

// UnmarshalNullFloat is a custom unmarshaller.
func UnmarshalNullFloat(v interface{}) (sql.NullFloat64, error) {
	if v == nil {
		return sql.NullFloat64{Valid: false}, nil
	}
	f, err := graphql.UnmarshalFloat(v)
	return sql.NullFloat64{Float64: f, Valid: true}, err
}
