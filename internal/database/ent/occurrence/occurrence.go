// Code generated by ent, DO NOT EDIT.

package occurrence

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/mensatt/backend/internal/database/schema"
)

const (
	// Label holds the string label denoting the occurrence type in the database.
	Label = "occurrence"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldDate holds the string denoting the date field in the database.
	FieldDate = "date"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// FieldKj holds the string denoting the kj field in the database.
	FieldKj = "kj"
	// FieldKcal holds the string denoting the kcal field in the database.
	FieldKcal = "kcal"
	// FieldFat holds the string denoting the fat field in the database.
	FieldFat = "fat"
	// FieldSaturatedFat holds the string denoting the saturated_fat field in the database.
	FieldSaturatedFat = "saturated_fat"
	// FieldCarbohydrates holds the string denoting the carbohydrates field in the database.
	FieldCarbohydrates = "carbohydrates"
	// FieldSugar holds the string denoting the sugar field in the database.
	FieldSugar = "sugar"
	// FieldFiber holds the string denoting the fiber field in the database.
	FieldFiber = "fiber"
	// FieldProtein holds the string denoting the protein field in the database.
	FieldProtein = "protein"
	// FieldSalt holds the string denoting the salt field in the database.
	FieldSalt = "salt"
	// FieldPriceStudent holds the string denoting the price_student field in the database.
	FieldPriceStudent = "price_student"
	// FieldPriceStaff holds the string denoting the price_staff field in the database.
	FieldPriceStaff = "price_staff"
	// FieldPriceGuest holds the string denoting the price_guest field in the database.
	FieldPriceGuest = "price_guest"
	// EdgeLocation holds the string denoting the location edge name in mutations.
	EdgeLocation = "location"
	// EdgeDish holds the string denoting the dish edge name in mutations.
	EdgeDish = "dish"
	// EdgeTags holds the string denoting the tags edge name in mutations.
	EdgeTags = "tags"
	// EdgeSideDishes holds the string denoting the side_dishes edge name in mutations.
	EdgeSideDishes = "side_dishes"
	// EdgeReviews holds the string denoting the reviews edge name in mutations.
	EdgeReviews = "reviews"
	// TagFieldID holds the string denoting the ID field of the Tag.
	TagFieldID = "key"
	// Table holds the table name of the occurrence in the database.
	Table = "occurrence"
	// LocationTable is the table that holds the location relation/edge.
	LocationTable = "occurrence"
	// LocationInverseTable is the table name for the Location entity.
	// It exists in this package in order to avoid circular dependency with the "location" package.
	LocationInverseTable = "location"
	// LocationColumn is the table column denoting the location relation/edge.
	LocationColumn = "location"
	// DishTable is the table that holds the dish relation/edge.
	DishTable = "occurrence"
	// DishInverseTable is the table name for the Dish entity.
	// It exists in this package in order to avoid circular dependency with the "dish" package.
	DishInverseTable = "dish"
	// DishColumn is the table column denoting the dish relation/edge.
	DishColumn = "dish"
	// TagsTable is the table that holds the tags relation/edge. The primary key declared below.
	TagsTable = "occurrence_tags"
	// TagsInverseTable is the table name for the Tag entity.
	// It exists in this package in order to avoid circular dependency with the "tag" package.
	TagsInverseTable = "tag"
	// SideDishesTable is the table that holds the side_dishes relation/edge.
	SideDishesTable = "dish"
	// SideDishesInverseTable is the table name for the Dish entity.
	// It exists in this package in order to avoid circular dependency with the "dish" package.
	SideDishesInverseTable = "dish"
	// SideDishesColumn is the table column denoting the side_dishes relation/edge.
	SideDishesColumn = "occurrence_side_dishes"
	// ReviewsTable is the table that holds the reviews relation/edge.
	ReviewsTable = "review"
	// ReviewsInverseTable is the table name for the Review entity.
	// It exists in this package in order to avoid circular dependency with the "review" package.
	ReviewsInverseTable = "review"
	// ReviewsColumn is the table column denoting the reviews relation/edge.
	ReviewsColumn = "occurrence"
)

// Columns holds all SQL columns for occurrence fields.
var Columns = []string{
	FieldID,
	FieldDate,
	FieldStatus,
	FieldKj,
	FieldKcal,
	FieldFat,
	FieldSaturatedFat,
	FieldCarbohydrates,
	FieldSugar,
	FieldFiber,
	FieldProtein,
	FieldSalt,
	FieldPriceStudent,
	FieldPriceStaff,
	FieldPriceGuest,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "occurrence"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"dish",
	"location",
}

var (
	// TagsPrimaryKey and TagsColumn2 are the table columns denoting the
	// primary key for the tags relation (M2M).
	TagsPrimaryKey = []string{"occurrence_id", "tag_id"}
)

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
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)

const DefaultStatus schema.OccurrenceStatus = "AWAITING_APPROVAL"

// StatusValidator is a validator for the "status" field enum values. It is called by the builders before save.
func StatusValidator(s schema.OccurrenceStatus) error {
	switch s {
	case "CONFIRMED", "APPROVED", "AWAITING_APPROVAL", "UPDATED", "PENDING_DELETION":
		return nil
	default:
		return fmt.Errorf("occurrence: invalid enum value for status field: %q", s)
	}
}
