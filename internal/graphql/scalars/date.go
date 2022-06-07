package scalars

import (
	"fmt"
	"io"
	"strconv"
	"time"

	"github.com/99designs/gqlgen/graphql"
)

func MarshalDate(d time.Time) graphql.Marshaler {
	return graphql.WriterFunc(func(w io.Writer) {
		w.Write([]byte(strconv.Quote(d.Format("2006-01-02"))))
	})
}

func UnmarshalDate(v interface{}) (time.Time, error) {
	if dateStr, ok := v.(string); ok {
		d, err := time.ParseInLocation("2006-01-02", dateStr, time.UTC)
		if err != nil {
			return time.Time{}, err
		}
		return d, nil
	}
	return time.Time{}, fmt.Errorf("%T is not a date string", v)
}
