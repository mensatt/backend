// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/mensatt/backend/internal/database/ent/image"
	"github.com/mensatt/backend/internal/database/ent/review"
)

// Image is the model entity for the Image schema.
type Image struct {
	config
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ImageQuery when eager-loading is set.
	Edges        ImageEdges `json:"edges"`
	review       *uuid.UUID
	selectValues sql.SelectValues
}

// ImageEdges holds the relations/edges for other nodes in the graph.
type ImageEdges struct {
	// Review holds the value of the review edge.
	Review *Review `json:"review,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// ReviewOrErr returns the Review value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ImageEdges) ReviewOrErr() (*Review, error) {
	if e.Review != nil {
		return e.Review, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: review.Label}
	}
	return nil, &NotLoadedError{edge: "review"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Image) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case image.FieldID:
			values[i] = new(uuid.UUID)
		case image.ForeignKeys[0]: // review
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Image fields.
func (i *Image) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for j := range columns {
		switch columns[j] {
		case image.FieldID:
			if value, ok := values[j].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[j])
			} else if value != nil {
				i.ID = *value
			}
		case image.ForeignKeys[0]:
			if value, ok := values[j].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field review", values[j])
			} else if value.Valid {
				i.review = new(uuid.UUID)
				*i.review = *value.S.(*uuid.UUID)
			}
		default:
			i.selectValues.Set(columns[j], values[j])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Image.
// This includes values selected through modifiers, order, etc.
func (i *Image) Value(name string) (ent.Value, error) {
	return i.selectValues.Get(name)
}

// QueryReview queries the "review" edge of the Image entity.
func (i *Image) QueryReview() *ReviewQuery {
	return NewImageClient(i.config).QueryReview(i)
}

// Update returns a builder for updating this Image.
// Note that you need to call Image.Unwrap() before calling this method if this Image
// was returned from a transaction, and the transaction was committed or rolled back.
func (i *Image) Update() *ImageUpdateOne {
	return NewImageClient(i.config).UpdateOne(i)
}

// Unwrap unwraps the Image entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (i *Image) Unwrap() *Image {
	_tx, ok := i.config.driver.(*txDriver)
	if !ok {
		panic("ent: Image is not a transactional entity")
	}
	i.config.driver = _tx.drv
	return i
}

// String implements the fmt.Stringer.
func (i *Image) String() string {
	var builder strings.Builder
	builder.WriteString("Image(")
	builder.WriteString(fmt.Sprintf("id=%v", i.ID))
	builder.WriteByte(')')
	return builder.String()
}

// Images is a parsable slice of Image.
type Images []*Image
