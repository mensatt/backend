package scalars

import (
	"database/sql"

	"github.com/99designs/gqlgen/graphql"
)

// MarshalNullInt is a custom marshaller.
func MarshalNullInt(ns sql.NullInt32) graphql.Marshaler {
	if !ns.Valid {
		return graphql.Null
	}
	return graphql.MarshalInt32(ns.Int32)
}

// UnmarshalNullInt is a custom unmarshaller.
func UnmarshalNullInt(v interface{}) (sql.NullInt32, error) {
	if v == nil {
		return sql.NullInt32{Valid: false}, nil
	}
	i, err := graphql.UnmarshalInt32(v)
	return sql.NullInt32{Int32: i, Valid: true}, err
}
