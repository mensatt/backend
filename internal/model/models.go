package models

import (
	"github.com/google/uuid"
	"time"
)

/*
 * Bun generates table names and aliases from struct names by underscoring them. It also pluralizes table
 * names, for example, struct ArticleCategory gets table name article_categories and alias
 * article_category.
 */

type Dish struct {
	ID   uuid.UUID `json:"id" bun:"type:uuid,default:uuid_generate_v4(),pk"`
	Name string    `json:"name"`
}

type Allergy struct {
	Abbreviation string `json:"abbreviation" bun:"pk"`
	Name         string `json:"name"`
}

type Tag struct {
	Abbreviation string `json:"abbreviation" bun:"pk"`
	Name         string `json:"name"`
}

type Occurrence struct {
	ID           uuid.UUID  `json:"id" bun:"type:uuid,default:uuid_generate_v4(),pk"`
	Dish         *Dish      `json:"dish" bun:"rel:has-one,join:id=dish_id"`
	SideDishes   []*Dish    `json:"sideDishes" bun:"rel:has-many,join:id=dish_id"` // is the has-many relationship correct?
	Date         time.Time  `json:"date"`
	PriceStudent uint16     `json:"priceStudent"`
	PriceStaff   uint16     `json:"priceStaff"`
	PriceGuest   uint16     `json:"priceGuest"`
	Allergies    []*Allergy `json:"allergies" bun:"rel:has-many,join:abbreviation=allergy_abbreviation"` // is the has-many relationship correct?
	Tags         []*Tag     `json:"tags" bun:"rel:has-many,join:abbreviation=tag_abbreviation"`          // is the has-many relationship correct?
}

type Review struct {
	ID          uuid.UUID   `json:"id" bun:"type:uuid,default:uuid_generate_v4(),pk"`
	Occurrence  *Occurrence `json:"occurrence" bun:"rel:has-one,join:id=occurrence_id"`
	DisplayName *string     `json:"displayName"`
	Stars       int         `json:"stars"`
	Text        *string     `json:"text"`
	UpVotes     uint32      `json:"upVotes"`
	DownVotes   uint32      `json:"downVotes"`
	AcceptedAt  time.Time   `json:"acceptedAt"`
	CreatedAt   time.Time   `json:"createdAt"`
	UpdatedAt   time.Time   `json:"updatedAt"`
}

type Image struct {
	ID          uuid.UUID   `json:"id" bun:"type:uuid,default:uuid_generate_v4(),pk"`
	Occurrence  *Occurrence `json:"occurrence" bun:"rel:has-one,join:id=occurrence_id"`
	DisplayName *string     `json:"displayName"`
	Description *string     `json:"description"`
	UpVotes     uint32      `json:"upVotes"`
	DownVotes   uint32      `json:"downVotes"`
	AcceptedAt  time.Time   `json:"acceptedAt"`
	CreatedAt   time.Time   `json:"createdAt"`
	UpdatedAt   time.Time   `json:"updatedAt"`
}
