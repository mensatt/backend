// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/mensatt/backend/internal/database/ent/dish"
)

// Dish is the model entity for the Dish schema.
type Dish struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// NameDe holds the value of the "name_de" field.
	NameDe string `json:"name_de,omitempty"`
	// NameEn holds the value of the "name_en" field.
	NameEn *string `json:"name_en,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the DishQuery when eager-loading is set.
	Edges DishEdges `json:"edges"`
}

// DishEdges holds the relations/edges for other nodes in the graph.
type DishEdges struct {
	// DishOccurrences holds the value of the dish_occurrences edge.
	DishOccurrences []*Occurrence `json:"dish_occurrences,omitempty"`
	// Aliases holds the value of the aliases edge.
	Aliases []*DishAlias `json:"aliases,omitempty"`
	// SideDishOccurrence holds the value of the side_dish_occurrence edge.
	SideDishOccurrence []*Occurrence `json:"side_dish_occurrence,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
}

// DishOccurrencesOrErr returns the DishOccurrences value or an error if the edge
// was not loaded in eager-loading.
func (e DishEdges) DishOccurrencesOrErr() ([]*Occurrence, error) {
	if e.loadedTypes[0] {
		return e.DishOccurrences, nil
	}
	return nil, &NotLoadedError{edge: "dish_occurrences"}
}

// AliasesOrErr returns the Aliases value or an error if the edge
// was not loaded in eager-loading.
func (e DishEdges) AliasesOrErr() ([]*DishAlias, error) {
	if e.loadedTypes[1] {
		return e.Aliases, nil
	}
	return nil, &NotLoadedError{edge: "aliases"}
}

// SideDishOccurrenceOrErr returns the SideDishOccurrence value or an error if the edge
// was not loaded in eager-loading.
func (e DishEdges) SideDishOccurrenceOrErr() ([]*Occurrence, error) {
	if e.loadedTypes[2] {
		return e.SideDishOccurrence, nil
	}
	return nil, &NotLoadedError{edge: "side_dish_occurrence"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Dish) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case dish.FieldNameDe, dish.FieldNameEn:
			values[i] = new(sql.NullString)
		case dish.FieldID:
			values[i] = new(uuid.UUID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Dish", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Dish fields.
func (d *Dish) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case dish.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				d.ID = *value
			}
		case dish.FieldNameDe:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name_de", values[i])
			} else if value.Valid {
				d.NameDe = value.String
			}
		case dish.FieldNameEn:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name_en", values[i])
			} else if value.Valid {
				d.NameEn = new(string)
				*d.NameEn = value.String
			}
		}
	}
	return nil
}

// QueryDishOccurrences queries the "dish_occurrences" edge of the Dish entity.
func (d *Dish) QueryDishOccurrences() *OccurrenceQuery {
	return (&DishClient{config: d.config}).QueryDishOccurrences(d)
}

// QueryAliases queries the "aliases" edge of the Dish entity.
func (d *Dish) QueryAliases() *DishAliasQuery {
	return (&DishClient{config: d.config}).QueryAliases(d)
}

// QuerySideDishOccurrence queries the "side_dish_occurrence" edge of the Dish entity.
func (d *Dish) QuerySideDishOccurrence() *OccurrenceQuery {
	return (&DishClient{config: d.config}).QuerySideDishOccurrence(d)
}

// Update returns a builder for updating this Dish.
// Note that you need to call Dish.Unwrap() before calling this method if this Dish
// was returned from a transaction, and the transaction was committed or rolled back.
func (d *Dish) Update() *DishUpdateOne {
	return (&DishClient{config: d.config}).UpdateOne(d)
}

// Unwrap unwraps the Dish entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (d *Dish) Unwrap() *Dish {
	_tx, ok := d.config.driver.(*txDriver)
	if !ok {
		panic("ent: Dish is not a transactional entity")
	}
	d.config.driver = _tx.drv
	return d
}

// String implements the fmt.Stringer.
func (d *Dish) String() string {
	var builder strings.Builder
	builder.WriteString("Dish(")
	builder.WriteString(fmt.Sprintf("id=%v, ", d.ID))
	builder.WriteString("name_de=")
	builder.WriteString(d.NameDe)
	builder.WriteString(", ")
	if v := d.NameEn; v != nil {
		builder.WriteString("name_en=")
		builder.WriteString(*v)
	}
	builder.WriteByte(')')
	return builder.String()
}

// Dishes is a parsable slice of Dish.
type Dishes []*Dish

func (d Dishes) config(cfg config) {
	for _i := range d {
		d[_i].config = cfg
	}
}
