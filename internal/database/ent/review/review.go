// Code generated by ent, DO NOT EDIT.

package review

import (
	"time"

	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the review type in the database.
	Label = "review"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldDisplayName holds the string denoting the display_name field in the database.
	FieldDisplayName = "display_name"
	// FieldStars holds the string denoting the stars field in the database.
	FieldStars = "stars"
	// FieldText holds the string denoting the text field in the database.
	FieldText = "text"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldAcceptedAt holds the string denoting the accepted_at field in the database.
	FieldAcceptedAt = "accepted_at"
	// EdgeOccurrence holds the string denoting the occurrence edge name in mutations.
	EdgeOccurrence = "occurrence"
	// EdgeImages holds the string denoting the images edge name in mutations.
	EdgeImages = "images"
	// Table holds the table name of the review in the database.
	Table = "review"
	// OccurrenceTable is the table that holds the occurrence relation/edge.
	OccurrenceTable = "review"
	// OccurrenceInverseTable is the table name for the Occurrence entity.
	// It exists in this package in order to avoid circular dependency with the "occurrence" package.
	OccurrenceInverseTable = "occurrence"
	// OccurrenceColumn is the table column denoting the occurrence relation/edge.
	OccurrenceColumn = "occurrence"
	// ImagesTable is the table that holds the images relation/edge.
	ImagesTable = "image"
	// ImagesInverseTable is the table name for the Image entity.
	// It exists in this package in order to avoid circular dependency with the "image" package.
	ImagesInverseTable = "image"
	// ImagesColumn is the table column denoting the images relation/edge.
	ImagesColumn = "review"
)

// Columns holds all SQL columns for review fields.
var Columns = []string{
	FieldID,
	FieldDisplayName,
	FieldStars,
	FieldText,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldAcceptedAt,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "review"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"occurrence",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// DisplayNameValidator is a validator for the "display_name" field. It is called by the builders before save.
	DisplayNameValidator func(string) error
	// StarsValidator is a validator for the "stars" field. It is called by the builders before save.
	StarsValidator func(int) error
	// TextValidator is a validator for the "text" field. It is called by the builders before save.
	TextValidator func(string) error
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)
