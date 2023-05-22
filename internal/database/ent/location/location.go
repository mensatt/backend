// Code generated by ent, DO NOT EDIT.

package location

import (
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the location type in the database.
	Label = "location"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldExternalID holds the string denoting the external_id field in the database.
	FieldExternalID = "external_id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// EdgeOccurrences holds the string denoting the occurrences edge name in mutations.
	EdgeOccurrences = "occurrences"
	// Table holds the table name of the location in the database.
	Table = "locations"
	// OccurrencesTable is the table that holds the occurrences relation/edge.
	OccurrencesTable = "occurrences"
	// OccurrencesInverseTable is the table name for the Occurrence entity.
	// It exists in this package in order to avoid circular dependency with the "occurrence" package.
	OccurrencesInverseTable = "occurrences"
	// OccurrencesColumn is the table column denoting the occurrences relation/edge.
	OccurrencesColumn = "location"
)

// Columns holds all SQL columns for location fields.
var Columns = []string{
	FieldID,
	FieldExternalID,
	FieldName,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)
