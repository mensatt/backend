package scalars

import (
	"fmt"
	"io"
	"strconv"

	"github.com/99designs/gqlgen/graphql"
	"github.com/mensatt/mensatt-backend/internal/db/sqlc"
)

func MarshalPriority(p sqlc.Priority) graphql.Marshaler {
	return graphql.WriterFunc(func(w io.Writer) {
		w.Write([]byte(strconv.Quote(string(p))))
	})
}

func UnmarshalPriority(v interface{}) (sqlc.Priority, error) {
	var p sqlc.Priority
	if pStr, ok := v.(string); ok {
		err := p.Scan(pStr)
		if err != nil {
			return p, err
		}
		return p, nil
	}
	return p, fmt.Errorf("%T is not a Priority string", v)
}
