package scalars

import (
	"github.com/99designs/gqlgen/graphql"
	"github.com/google/uuid"
)

func MarshalNullUUID(id uuid.NullUUID) graphql.Marshaler {
	if !id.Valid {
		return graphql.Null
	}
	return MarshalUUID(id.UUID)
}

func UnmarshalNullUUID(v interface{}) (uuid.NullUUID, error) {
	if v == nil {
		return uuid.NullUUID{Valid: false}, nil
	}
	id, err := UnmarshalUUID(v)
	return uuid.NullUUID{UUID: id, Valid: true}, err
}
