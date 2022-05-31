package scalars

import (
	"fmt"
	"io"
	"strconv"

	"github.com/99designs/gqlgen/graphql"
	"github.com/mensatt/backend/internal/db/sqlc"
)

func MarshalReviewStatus(p sqlc.ReviewStatus) graphql.Marshaler {
	return graphql.WriterFunc(func(w io.Writer) {
		w.Write([]byte(strconv.Quote(string(p))))
	})
}

func UnmarshalReviewStatus(v interface{}) (sqlc.ReviewStatus, error) {
	var p sqlc.ReviewStatus
	if pStr, ok := v.(string); ok {
		err := p.Scan(pStr)
		if err != nil {
			return p, err
		}
		return p, nil
	}
	return p, fmt.Errorf("%T is not a ReviewStatus string", v)
}
