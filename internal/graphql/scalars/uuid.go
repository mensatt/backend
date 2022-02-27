package scalars

import (
	"fmt"
	"io"
	"strconv"

	"github.com/99designs/gqlgen/graphql"
	"github.com/google/uuid"
)

func MarshalUUID(id uuid.UUID) graphql.Marshaler {
	return graphql.WriterFunc(func(w io.Writer) {
		w.Write([]byte(strconv.Quote(id.String())))
	})
}

func UnmarshalUUID(v interface{}) (uuid.UUID, error) {
	if idStr, ok := v.(string); ok {
		id, err := uuid.Parse(idStr)
		if err != nil {
			return uuid.Nil, err
		}
		return id, nil
	}
	return uuid.Nil, fmt.Errorf("%T is not a uuid string", v)
}
