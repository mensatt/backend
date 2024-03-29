// Code generated by ent, DO NOT EDIT.

package ent

import (
	"time"

	"github.com/google/uuid"
	"github.com/mensatt/backend/internal/database/ent/dish"
	"github.com/mensatt/backend/internal/database/ent/dishalias"
	"github.com/mensatt/backend/internal/database/ent/image"
	"github.com/mensatt/backend/internal/database/ent/location"
	"github.com/mensatt/backend/internal/database/ent/occurrence"
	"github.com/mensatt/backend/internal/database/ent/review"
	"github.com/mensatt/backend/internal/database/ent/tag"
	"github.com/mensatt/backend/internal/database/ent/user"
	"github.com/mensatt/backend/internal/database/schema"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	dishFields := schema.Dish{}.Fields()
	_ = dishFields
	// dishDescNameDe is the schema descriptor for name_de field.
	dishDescNameDe := dishFields[1].Descriptor()
	// dish.NameDeValidator is a validator for the "name_de" field. It is called by the builders before save.
	dish.NameDeValidator = dishDescNameDe.Validators[0].(func(string) error)
	// dishDescNameEn is the schema descriptor for name_en field.
	dishDescNameEn := dishFields[2].Descriptor()
	// dish.NameEnValidator is a validator for the "name_en" field. It is called by the builders before save.
	dish.NameEnValidator = dishDescNameEn.Validators[0].(func(string) error)
	// dishDescID is the schema descriptor for id field.
	dishDescID := dishFields[0].Descriptor()
	// dish.DefaultID holds the default value on creation for the id field.
	dish.DefaultID = dishDescID.Default.(func() uuid.UUID)
	dishaliasFields := schema.DishAlias{}.Fields()
	_ = dishaliasFields
	// dishaliasDescNormalizedAliasName is the schema descriptor for normalized_alias_name field.
	dishaliasDescNormalizedAliasName := dishaliasFields[1].Descriptor()
	// dishalias.NormalizedAliasNameValidator is a validator for the "normalized_alias_name" field. It is called by the builders before save.
	dishalias.NormalizedAliasNameValidator = dishaliasDescNormalizedAliasName.Validators[0].(func(string) error)
	imageFields := schema.Image{}.Fields()
	_ = imageFields
	// imageDescID is the schema descriptor for id field.
	imageDescID := imageFields[0].Descriptor()
	// image.DefaultID holds the default value on creation for the id field.
	image.DefaultID = imageDescID.Default.(func() uuid.UUID)
	locationFields := schema.Location{}.Fields()
	_ = locationFields
	// locationDescName is the schema descriptor for name field.
	locationDescName := locationFields[2].Descriptor()
	// location.NameValidator is a validator for the "name" field. It is called by the builders before save.
	location.NameValidator = locationDescName.Validators[0].(func(string) error)
	// locationDescVisible is the schema descriptor for visible field.
	locationDescVisible := locationFields[3].Descriptor()
	// location.DefaultVisible holds the default value on creation for the visible field.
	location.DefaultVisible = locationDescVisible.Default.(bool)
	// locationDescID is the schema descriptor for id field.
	locationDescID := locationFields[0].Descriptor()
	// location.DefaultID holds the default value on creation for the id field.
	location.DefaultID = locationDescID.Default.(func() uuid.UUID)
	occurrenceFields := schema.Occurrence{}.Fields()
	_ = occurrenceFields
	// occurrenceDescDate is the schema descriptor for date field.
	occurrenceDescDate := occurrenceFields[1].Descriptor()
	// occurrence.DefaultDate holds the default value on creation for the date field.
	occurrence.DefaultDate = occurrenceDescDate.Default.(func() time.Time)
	// occurrenceDescID is the schema descriptor for id field.
	occurrenceDescID := occurrenceFields[0].Descriptor()
	// occurrence.DefaultID holds the default value on creation for the id field.
	occurrence.DefaultID = occurrenceDescID.Default.(func() uuid.UUID)
	reviewFields := schema.Review{}.Fields()
	_ = reviewFields
	// reviewDescDisplayName is the schema descriptor for display_name field.
	reviewDescDisplayName := reviewFields[1].Descriptor()
	// review.DisplayNameValidator is a validator for the "display_name" field. It is called by the builders before save.
	review.DisplayNameValidator = func() func(string) error {
		validators := reviewDescDisplayName.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(display_name string) error {
			for _, fn := range fns {
				if err := fn(display_name); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// reviewDescStars is the schema descriptor for stars field.
	reviewDescStars := reviewFields[2].Descriptor()
	// review.StarsValidator is a validator for the "stars" field. It is called by the builders before save.
	review.StarsValidator = func() func(int) error {
		validators := reviewDescStars.Validators
		fns := [...]func(int) error{
			validators[0].(func(int) error),
			validators[1].(func(int) error),
		}
		return func(stars int) error {
			for _, fn := range fns {
				if err := fn(stars); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// reviewDescText is the schema descriptor for text field.
	reviewDescText := reviewFields[3].Descriptor()
	// review.TextValidator is a validator for the "text" field. It is called by the builders before save.
	review.TextValidator = reviewDescText.Validators[0].(func(string) error)
	// reviewDescCreatedAt is the schema descriptor for created_at field.
	reviewDescCreatedAt := reviewFields[4].Descriptor()
	// review.DefaultCreatedAt holds the default value on creation for the created_at field.
	review.DefaultCreatedAt = reviewDescCreatedAt.Default.(func() time.Time)
	// reviewDescUpdatedAt is the schema descriptor for updated_at field.
	reviewDescUpdatedAt := reviewFields[5].Descriptor()
	// review.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	review.DefaultUpdatedAt = reviewDescUpdatedAt.Default.(func() time.Time)
	// review.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	review.UpdateDefaultUpdatedAt = reviewDescUpdatedAt.UpdateDefault.(func() time.Time)
	// reviewDescID is the schema descriptor for id field.
	reviewDescID := reviewFields[0].Descriptor()
	// review.DefaultID holds the default value on creation for the id field.
	review.DefaultID = reviewDescID.Default.(func() uuid.UUID)
	tagFields := schema.Tag{}.Fields()
	_ = tagFields
	// tagDescName is the schema descriptor for name field.
	tagDescName := tagFields[1].Descriptor()
	// tag.NameValidator is a validator for the "name" field. It is called by the builders before save.
	tag.NameValidator = tagDescName.Validators[0].(func(string) error)
	// tagDescDescription is the schema descriptor for description field.
	tagDescDescription := tagFields[2].Descriptor()
	// tag.DescriptionValidator is a validator for the "description" field. It is called by the builders before save.
	tag.DescriptionValidator = tagDescDescription.Validators[0].(func(string) error)
	// tagDescShortName is the schema descriptor for short_name field.
	tagDescShortName := tagFields[3].Descriptor()
	// tag.ShortNameValidator is a validator for the "short_name" field. It is called by the builders before save.
	tag.ShortNameValidator = tagDescShortName.Validators[0].(func(string) error)
	// tagDescIsAllergy is the schema descriptor for is_allergy field.
	tagDescIsAllergy := tagFields[5].Descriptor()
	// tag.DefaultIsAllergy holds the default value on creation for the is_allergy field.
	tag.DefaultIsAllergy = tagDescIsAllergy.Default.(bool)
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescEmail is the schema descriptor for email field.
	userDescEmail := userFields[1].Descriptor()
	// user.EmailValidator is a validator for the "email" field. It is called by the builders before save.
	user.EmailValidator = userDescEmail.Validators[0].(func(string) error)
	// userDescPasswordHash is the schema descriptor for password_hash field.
	userDescPasswordHash := userFields[2].Descriptor()
	// user.PasswordHashValidator is a validator for the "password_hash" field. It is called by the builders before save.
	user.PasswordHashValidator = userDescPasswordHash.Validators[0].(func(string) error)
	// userDescCreatedAt is the schema descriptor for created_at field.
	userDescCreatedAt := userFields[3].Descriptor()
	// user.DefaultCreatedAt holds the default value on creation for the created_at field.
	user.DefaultCreatedAt = userDescCreatedAt.Default.(func() time.Time)
	// userDescUpdatedAt is the schema descriptor for updated_at field.
	userDescUpdatedAt := userFields[4].Descriptor()
	// user.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	user.DefaultUpdatedAt = userDescUpdatedAt.Default.(func() time.Time)
	// user.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	user.UpdateDefaultUpdatedAt = userDescUpdatedAt.UpdateDefault.(func() time.Time)
	// userDescID is the schema descriptor for id field.
	userDescID := userFields[0].Descriptor()
	// user.DefaultID holds the default value on creation for the id field.
	user.DefaultID = userDescID.Default.(func() uuid.UUID)
}
