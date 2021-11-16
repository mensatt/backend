// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package models

type Allergy struct {
	Abbreviation string `json:"abbreviation"`
	Name         string `json:"name"`
}

type Dish struct {
	ID        string     `json:"id"`
	Name      string     `json:"name"`
	Allergies []*Allergy `json:"allergies"`
}

type Image struct {
	ID          string      `json:"id"`
	Occurrence  *Occurrence `json:"occurrence"`
	AcceptedAt  *string     `json:"acceptedAt"`
	CreatedAt   string      `json:"createdAt"`
	Votes       int         `json:"votes"`
	Description *string     `json:"description"`
	DisplayName *string     `json:"displayName"`
}

type Occurrence struct {
	ID           string  `json:"id"`
	Dish         *Dish   `json:"dish"`
	SideDishes   []*Dish `json:"sideDishes"`
	Date         string  `json:"date"`
	PriceStudent float64 `json:"priceStudent"`
	PriceStaff   float64 `json:"priceStaff"`
	PriceGuest   float64 `json:"priceGuest"`
	Tags         []*Tag  `json:"tags"`
}

type Review struct {
	ID          string      `json:"id"`
	DisplayName string      `json:"displayName"`
	Stars       int         `json:"stars"`
	Text        *string     `json:"text"`
	AcceptedAt  *string     `json:"acceptedAt"`
	CreatedAt   string      `json:"createdAt"`
	UpdatedAt   string      `json:"updatedAt"`
	Occurrence  *Occurrence `json:"occurrence"`
	Votes       int         `json:"votes"`
}

type ReviewInput struct {
	DisplayName  string  `json:"displayName"`
	Stars        int     `json:"stars"`
	Text         *string `json:"text"`
	OccurrenceID string  `json:"occurrenceId"`
}

type Tag struct {
	Abbreviation string `json:"abbreviation"`
	Name         string `json:"name"`
}
