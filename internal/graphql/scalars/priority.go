package scalars

import (
	"fmt"
	"io"
	"strconv"

	"github.com/99designs/gqlgen/graphql"
	"github.com/mensatt/mensatt-backend/internal/db"
)

func MarshalPriority(p db.Priority) graphql.Marshaler {
	return graphql.WriterFunc(func(w io.Writer) {
		w.Write([]byte(strconv.Quote(string(p))))
	})
}

func UnmarshalPriority(v interface{}) (db.Priority, error) {
	var p db.Priority
	if pStr, ok := v.(string); ok {
		err := p.Scan(pStr)
		if err != nil {
			return p, err
		}
		return p, nil
	}
	return p, fmt.Errorf("%T is not a Priority string", v)
}
