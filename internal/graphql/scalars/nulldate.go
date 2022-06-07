package scalars

import (
	"database/sql"

	"github.com/99designs/gqlgen/graphql"
)

func MarshalNullDate(d sql.NullTime) graphql.Marshaler {
	if !d.Valid {
		return graphql.Null
	}
	return MarshalDate(d.Time)
}

func UnmarshalNullDate(v interface{}) (sql.NullTime, error) {
	if v == nil {
		return sql.NullTime{Valid: false}, nil
	}
	d, err := UnmarshalDate(v)
	return sql.NullTime{Time: d, Valid: true}, err
}
