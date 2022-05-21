// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.13.0

package sqlc

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/google/uuid"
)

type Priority string

const (
	PriorityUNSET  Priority = "UNSET"
	PriorityLOW    Priority = "LOW"
	PriorityMEDIUM Priority = "MEDIUM"
	PriorityHIGH   Priority = "HIGH"
)

func (e *Priority) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = Priority(s)
	case string:
		*e = Priority(s)
	default:
		return fmt.Errorf("unsupported scan type for Priority: %T", src)
	}
	return nil
}

type Dish struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

type Image struct {
	ID          uuid.UUID      `json:"id"`
	Occurrence  uuid.UUID      `json:"occurrence"`
	DisplayName string         `json:"display_name"`
	Description sql.NullString `json:"description"`
	UpVotes     int32          `json:"up_votes"`
	DownVotes   int32          `json:"down_votes"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	AcceptedAt  sql.NullTime   `json:"accepted_at"`
}

type Occurrence struct {
	ID            uuid.UUID       `json:"id"`
	Dish          uuid.UUID       `json:"dish"`
	Date          time.Time       `json:"date"`
	Kj            sql.NullFloat64 `json:"kj"`
	Kcal          sql.NullFloat64 `json:"kcal"`
	Fat           sql.NullFloat64 `json:"fat"`
	SaturatedFat  sql.NullFloat64 `json:"saturated_fat"`
	Carbohydrates sql.NullFloat64 `json:"carbohydrates"`
	Sugar         sql.NullFloat64 `json:"sugar"`
	Fiber         sql.NullFloat64 `json:"fiber"`
	Protein       sql.NullFloat64 `json:"protein"`
	Salt          sql.NullFloat64 `json:"salt"`
	PriceStudent  int32           `json:"price_student"`
	PriceStaff    int32           `json:"price_staff"`
	PriceGuest    int32           `json:"price_guest"`
}

type OccurrenceSideDish struct {
	OccurrenceID uuid.UUID `json:"occurrence_id"`
	DishID       uuid.UUID `json:"dish_id"`
}

type OccurrenceTag struct {
	OccurrenceID uuid.UUID `json:"occurrence_id"`
	TagKey       string    `json:"tag_key"`
}

type Review struct {
	ID          uuid.UUID      `json:"id"`
	Occurrence  uuid.UUID      `json:"occurrence"`
	DisplayName sql.NullString `json:"display_name"`
	Stars       int32          `json:"stars"`
	Text        sql.NullString `json:"text"`
	UpVotes     int32          `json:"up_votes"`
	DownVotes   int32          `json:"down_votes"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	AcceptedAt  sql.NullTime   `json:"accepted_at"`
}

type Tag struct {
	Key         string         `json:"key"`
	Name        string         `json:"name"`
	Description string         `json:"description"`
	ShortName   sql.NullString `json:"short_name"`
	Priority    Priority       `json:"priority"`
	IsAllergy   bool           `json:"is_allergy"`
}