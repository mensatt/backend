// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/mensatt/backend/internal/database/ent/tag"
	"github.com/mensatt/backend/internal/database/schema"
)

// Tag is the model entity for the Tag schema.
type Tag struct {
	config `json:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Description holds the value of the "description" field.
	Description string `json:"description,omitempty"`
	// ShortName holds the value of the "short_name" field.
	ShortName *string `json:"short_name,omitempty"`
	// Priority holds the value of the "priority" field.
	Priority schema.TagPriority `json:"priority,omitempty"`
	// IsAllergy holds the value of the "is_allergy" field.
	IsAllergy bool `json:"is_allergy,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the TagQuery when eager-loading is set.
	Edges        TagEdges `json:"edges"`
	selectValues sql.SelectValues
}

// TagEdges holds the relations/edges for other nodes in the graph.
type TagEdges struct {
	// Occurrence holds the value of the occurrence edge.
	Occurrence []*Occurrence `json:"occurrence,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// OccurrenceOrErr returns the Occurrence value or an error if the edge
// was not loaded in eager-loading.
func (e TagEdges) OccurrenceOrErr() ([]*Occurrence, error) {
	if e.loadedTypes[0] {
		return e.Occurrence, nil
	}
	return nil, &NotLoadedError{edge: "occurrence"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Tag) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case tag.FieldIsAllergy:
			values[i] = new(sql.NullBool)
		case tag.FieldID, tag.FieldName, tag.FieldDescription, tag.FieldShortName, tag.FieldPriority:
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Tag fields.
func (t *Tag) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case tag.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				t.ID = value.String
			}
		case tag.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				t.Name = value.String
			}
		case tag.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				t.Description = value.String
			}
		case tag.FieldShortName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field short_name", values[i])
			} else if value.Valid {
				t.ShortName = new(string)
				*t.ShortName = value.String
			}
		case tag.FieldPriority:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field priority", values[i])
			} else if value.Valid {
				t.Priority = schema.TagPriority(value.String)
			}
		case tag.FieldIsAllergy:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field is_allergy", values[i])
			} else if value.Valid {
				t.IsAllergy = value.Bool
			}
		default:
			t.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Tag.
// This includes values selected through modifiers, order, etc.
func (t *Tag) Value(name string) (ent.Value, error) {
	return t.selectValues.Get(name)
}

// QueryOccurrence queries the "occurrence" edge of the Tag entity.
func (t *Tag) QueryOccurrence() *OccurrenceQuery {
	return NewTagClient(t.config).QueryOccurrence(t)
}

// Update returns a builder for updating this Tag.
// Note that you need to call Tag.Unwrap() before calling this method if this Tag
// was returned from a transaction, and the transaction was committed or rolled back.
func (t *Tag) Update() *TagUpdateOne {
	return NewTagClient(t.config).UpdateOne(t)
}

// Unwrap unwraps the Tag entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (t *Tag) Unwrap() *Tag {
	_tx, ok := t.config.driver.(*txDriver)
	if !ok {
		panic("ent: Tag is not a transactional entity")
	}
	t.config.driver = _tx.drv
	return t
}

// String implements the fmt.Stringer.
func (t *Tag) String() string {
	var builder strings.Builder
	builder.WriteString("Tag(")
	builder.WriteString(fmt.Sprintf("id=%v, ", t.ID))
	builder.WriteString("name=")
	builder.WriteString(t.Name)
	builder.WriteString(", ")
	builder.WriteString("description=")
	builder.WriteString(t.Description)
	builder.WriteString(", ")
	if v := t.ShortName; v != nil {
		builder.WriteString("short_name=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	builder.WriteString("priority=")
	builder.WriteString(fmt.Sprintf("%v", t.Priority))
	builder.WriteString(", ")
	builder.WriteString("is_allergy=")
	builder.WriteString(fmt.Sprintf("%v", t.IsAllergy))
	builder.WriteByte(')')
	return builder.String()
}

// Tags is a parsable slice of Tag.
type Tags []*Tag
