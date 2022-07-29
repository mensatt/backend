package scalars

import (
	"database/sql"

	"github.com/99designs/gqlgen/graphql"
)

// MarshalNullBool is a custom marshaller.
func MarshalNullBool(v sql.NullBool) graphql.Marshaler {
	if !v.Valid {
		return graphql.Null
	}
	return graphql.MarshalBoolean(v.Bool)
}

// UnmarshalNullBool is a custom unmarshaller.
func UnmarshalNullBool(v interface{}) (sql.NullBool, error) {
	if v == nil {
		return sql.NullBool{Valid: false}, nil
	}
	s, err := graphql.UnmarshalBoolean(v)
	return sql.NullBool{Bool: s, Valid: true}, err
}
