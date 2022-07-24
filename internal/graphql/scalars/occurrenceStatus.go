package scalars

import (
	"fmt"
	"io"
	"strconv"

	"github.com/99designs/gqlgen/graphql"
	"github.com/mensatt/backend/internal/db/sqlc"
)

func MarshalOccurrenceStatus(p sqlc.OccurrenceStatus) graphql.Marshaler {
	return graphql.WriterFunc(func(w io.Writer) {
		w.Write([]byte(strconv.Quote(string(p))))
	})
}

func UnmarshalOccurrenceStatus(v interface{}) (sqlc.OccurrenceStatus, error) {
	var p sqlc.OccurrenceStatus
	if pStr, ok := v.(string); ok {
		err := p.Scan(pStr)
		if err != nil {
			return p, err
		}
		return p, nil
	}
	return p, fmt.Errorf("%T is not a OccurrenceStatus string", v)
}
