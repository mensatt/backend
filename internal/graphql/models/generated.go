// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package models

import (
	"time"

	"github.com/99designs/gqlgen/graphql"
	"github.com/google/uuid"
	"github.com/mensatt/backend/internal/database/ent"
	"github.com/mensatt/backend/internal/database/schema"
)

type AddImagesToReviewInput struct {
	Review uuid.UUID     `json:"review"`
	Images []*ImageInput `json:"images"`
}

type AddSideDishToOccurrenceInput struct {
	Occurrence uuid.UUID `json:"occurrence"`
	Dish       uuid.UUID `json:"dish"`
}

type AddTagToOccurrenceInput struct {
	Occurrence uuid.UUID `json:"occurrence"`
	Tag        string    `json:"tag"`
}

type CreateDishAliasInput struct {
	AliasName           string    `json:"aliasName"`
	NormalizedAliasName string    `json:"normalizedAliasName"`
	Dish                uuid.UUID `json:"dish"`
}

type CreateDishInput struct {
	NameDe string  `json:"nameDe"`
	NameEn *string `json:"nameEn,omitempty"`
}

type CreateOccurrenceInput struct {
	Location      uuid.UUID   `json:"location"`
	Dish          uuid.UUID   `json:"dish"`
	SideDishes    []uuid.UUID `json:"sideDishes,omitempty"`
	Date          *time.Time  `json:"date,omitempty"`
	Kj            *int        `json:"kj,omitempty"`
	Kcal          *int        `json:"kcal,omitempty"`
	Fat           *int        `json:"fat,omitempty"`
	SaturatedFat  *int        `json:"saturatedFat,omitempty"`
	Carbohydrates *int        `json:"carbohydrates,omitempty"`
	Sugar         *int        `json:"sugar,omitempty"`
	Fiber         *int        `json:"fiber,omitempty"`
	Protein       *int        `json:"protein,omitempty"`
	Salt          *int        `json:"salt,omitempty"`
	PriceStudent  *int        `json:"priceStudent,omitempty"`
	PriceStaff    *int        `json:"priceStaff,omitempty"`
	PriceGuest    *int        `json:"priceGuest,omitempty"`
	Tags          []string    `json:"tags,omitempty"`
}

type CreateReviewInput struct {
	Occurrence  uuid.UUID     `json:"occurrence"`
	DisplayName *string       `json:"displayName,omitempty"`
	Images      []*ImageInput `json:"images,omitempty"`
	Stars       int           `json:"stars"`
	Text        *string       `json:"text,omitempty"`
}

type CreateTagInput struct {
	Key         string              `json:"key"`
	Name        string              `json:"name"`
	Description string              `json:"description"`
	ShortName   *string             `json:"shortName,omitempty"`
	Priority    *schema.TagPriority `json:"priority,omitempty"`
	IsAllergy   *bool               `json:"isAllergy,omitempty"`
}

type DeleteDishAliasInput struct {
	AliasName string `json:"aliasName"`
}

type DeleteImageToReviewInput struct {
	Review uuid.UUID `json:"review"`
	ID     uuid.UUID `json:"id"`
}

type DeleteOccurrenceInput struct {
	ID uuid.UUID `json:"id"`
}

type DeleteReviewInput struct {
	ID uuid.UUID `json:"id"`
}

type DishFilter struct {
	Dishes []uuid.UUID `json:"dishes,omitempty"`
	NameDe *string     `json:"nameDe,omitempty"`
	NameEn *string     `json:"nameEn,omitempty"`
}

type ImageInput struct {
	Image    graphql.Upload `json:"image"`
	Rotation *int           `json:"rotation,omitempty"`
}

type LocationFilter struct {
	Ids         []uuid.UUID `json:"ids,omitempty"`
	ExternalIds []int       `json:"externalIds,omitempty"`
	Names       []string    `json:"names,omitempty"`
	Visible     *bool       `json:"visible,omitempty"`
}

type LoginUserInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type OccurrenceFilter struct {
	Occurrences []uuid.UUID `json:"occurrences,omitempty"`
	Dishes      []uuid.UUID `json:"dishes,omitempty"`
	StartDate   *time.Time  `json:"startDate,omitempty"`
	EndDate     *time.Time  `json:"endDate,omitempty"`
	Location    *uuid.UUID  `json:"location,omitempty"`
}

type OccurrenceSideDish struct {
	Occurrence *ent.Occurrence `json:"occurrence"`
	Dish       *ent.Dish       `json:"dish"`
}

type OccurrenceTag struct {
	Occurrence *ent.Occurrence `json:"occurrence"`
	Tag        *ent.Tag        `json:"tag"`
}

type RemoveSideDishFromOccurrenceInput struct {
	Occurrence uuid.UUID `json:"occurrence"`
	Dish       uuid.UUID `json:"dish"`
}

type RemoveTagFromOccurrenceInput struct {
	Occurrence uuid.UUID `json:"occurrence"`
	Tag        string    `json:"tag"`
}

type ReviewDataDish struct {
	Reviews  []*ent.Review       `json:"reviews"`
	Images   []*ent.Image        `json:"images"`
	Metadata *ReviewMetadataDish `json:"metadata"`
}

type ReviewDataOccurrence struct {
	Reviews  []*ent.Review             `json:"reviews"`
	Images   []*ent.Image              `json:"images"`
	Metadata *ReviewMetadataOccurrence `json:"metadata"`
}

type ReviewFilter struct {
	Approved *bool `json:"approved,omitempty"`
}

type ReviewMetadataDish struct {
	AverageStars *float64 `json:"averageStars,omitempty"`
	ReviewCount  int      `json:"reviewCount"`
}

type ReviewMetadataOccurrence struct {
	AverageStars *float64 `json:"averageStars,omitempty"`
	ReviewCount  int      `json:"reviewCount"`
}

type UpdateDishInput struct {
	ID     uuid.UUID `json:"id"`
	NameDe *string   `json:"nameDe,omitempty"`
	NameEn *string   `json:"nameEn,omitempty"`
}

type UpdateOccurrenceInput struct {
	ID            uuid.UUID  `json:"id"`
	Dish          *uuid.UUID `json:"dish,omitempty"`
	Date          *time.Time `json:"date,omitempty"`
	Kj            *int       `json:"kj,omitempty"`
	Kcal          *int       `json:"kcal,omitempty"`
	Fat           *int       `json:"fat,omitempty"`
	SaturatedFat  *int       `json:"saturatedFat,omitempty"`
	Carbohydrates *int       `json:"carbohydrates,omitempty"`
	Sugar         *int       `json:"sugar,omitempty"`
	Fiber         *int       `json:"fiber,omitempty"`
	Protein       *int       `json:"protein,omitempty"`
	Salt          *int       `json:"salt,omitempty"`
	PriceStudent  *int       `json:"priceStudent,omitempty"`
	PriceStaff    *int       `json:"priceStaff,omitempty"`
	PriceGuest    *int       `json:"priceGuest,omitempty"`
}

type UpdateReviewInput struct {
	ID          uuid.UUID  `json:"id"`
	Occurrence  *uuid.UUID `json:"occurrence,omitempty"`
	DisplayName *string    `json:"displayName,omitempty"`
	Stars       *int       `json:"stars,omitempty"`
	Text        *string    `json:"text,omitempty"`
	Approved    *bool      `json:"approved,omitempty"`
}

type UpdateUserInput struct {
	ID       uuid.UUID        `json:"id"`
	Email    *string          `json:"email,omitempty"`
	Password *string          `json:"password,omitempty"`
	Username *string          `json:"username,omitempty"`
	Role     *schema.UserRole `json:"role,omitempty"`
}
