// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/mensatt/backend/internal/database/ent/image"
	"github.com/mensatt/backend/internal/database/ent/review"
)

// Image is the model entity for the Image schema.
type Image struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// ImageHash holds the value of the "image_hash" field.
	ImageHash string `json:"image_hash,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ImageQuery when eager-loading is set.
	Edges         ImageEdges `json:"edges"`
	review_images *uuid.UUID
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
	if e.loadedTypes[0] {
		if e.Review == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: review.Label}
		}
		return e.Review, nil
	}
	return nil, &NotLoadedError{edge: "review"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Image) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case image.FieldImageHash:
			values[i] = new(sql.NullString)
		case image.FieldID:
			values[i] = new(uuid.UUID)
		case image.ForeignKeys[0]: // review_images
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		default:
			return nil, fmt.Errorf("unexpected column %q for type Image", columns[i])
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
		case image.FieldImageHash:
			if value, ok := values[j].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field image_hash", values[j])
			} else if value.Valid {
				i.ImageHash = value.String
			}
		case image.ForeignKeys[0]:
			if value, ok := values[j].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field review_images", values[j])
			} else if value.Valid {
				i.review_images = new(uuid.UUID)
				*i.review_images = *value.S.(*uuid.UUID)
			}
		}
	}
	return nil
}

// QueryReview queries the "review" edge of the Image entity.
func (i *Image) QueryReview() *ReviewQuery {
	return (&ImageClient{config: i.config}).QueryReview(i)
}

// Update returns a builder for updating this Image.
// Note that you need to call Image.Unwrap() before calling this method if this Image
// was returned from a transaction, and the transaction was committed or rolled back.
func (i *Image) Update() *ImageUpdateOne {
	return (&ImageClient{config: i.config}).UpdateOne(i)
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
	builder.WriteString(fmt.Sprintf("id=%v, ", i.ID))
	builder.WriteString("image_hash=")
	builder.WriteString(i.ImageHash)
	builder.WriteByte(')')
	return builder.String()
}

// Images is a parsable slice of Image.
type Images []*Image

func (i Images) config(cfg config) {
	for _i := range i {
		i[_i].config = cfg
	}
}
