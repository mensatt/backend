package scalars

import (
	"database/sql"

	"github.com/99designs/gqlgen/graphql"
)

// MarshalNullString is a custom marshaller.
func MarshalNullString(ns sql.NullString) graphql.Marshaler {
	if !ns.Valid {
		return graphql.Null
	}
	return graphql.MarshalString(ns.String)
}

// UnmarshalNullString is a custom unmarshaller.
func UnmarshalNullString(v interface{}) (sql.NullString, error) {
	if v == nil {
		return sql.NullString{Valid: false}, nil
	}
	s, err := graphql.UnmarshalString(v)
	return sql.NullString{String: s, Valid: true}, err
}
