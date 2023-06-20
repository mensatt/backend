package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.27

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/mensatt/backend/internal/database/ent"
	"github.com/mensatt/backend/internal/database/ent/tag"
	"github.com/mensatt/backend/internal/database/ent/user"
	"github.com/mensatt/backend/internal/graphql/gqlserver"
	"github.com/mensatt/backend/internal/graphql/models"
	"github.com/mensatt/backend/pkg/utils"
)

// LoginUser is the resolver for the loginUser field.
func (r *mutationResolver) LoginUser(ctx context.Context, input models.LoginUserInput) (string, error) {
	user, err := r.Database.User.Query().Where(user.Email(input.Email)).Only(ctx)
	if err != nil {
		return "", err
	}

	if !utils.CheckPasswordHash(input.Password, user.PasswordHash) {
		return "", errors.New("wrong password")
	}

	return r.JWTKeyStore.GenerateJWT(user.ID)
}

// CreateTag is the resolver for the createTag field.
func (r *mutationResolver) CreateTag(ctx context.Context, input models.CreateTagInput) (*ent.Tag, error) {
	queryBuilder := r.Database.Tag.Create().
		SetID(input.Key).
		SetName(input.Name).
		SetDescription(input.Description)

	if input.ShortName != nil {
		queryBuilder = queryBuilder.SetShortName(*input.ShortName)
	}

	if input.Priority != nil {
		queryBuilder = queryBuilder.SetPriority(*input.Priority)
	}

	if input.IsAllergy != nil {
		queryBuilder = queryBuilder.SetIsAllergy(*input.IsAllergy)
	}

	return queryBuilder.Save(ctx)
}

// CreateDish is the resolver for the createDish field.
func (r *mutationResolver) CreateDish(ctx context.Context, input models.CreateDishInput) (*ent.Dish, error) {
	queryBuilder := r.Database.Dish.Create().
		SetNameDe(input.NameDe)

	if input.NameEn != nil {
		queryBuilder = queryBuilder.SetNameEn(*input.NameEn)
	}

	return queryBuilder.Save(ctx)
}

// UpdateDish is the resolver for the updateDish field.
func (r *mutationResolver) UpdateDish(ctx context.Context, input models.UpdateDishInput) (*ent.Dish, error) {
	queryBuilder := r.Database.Dish.UpdateOneID(input.ID)

	if input.NameDe != nil {
		queryBuilder = queryBuilder.SetNameDe(*input.NameDe)
	}

	if input.NameEn != nil {
		queryBuilder = queryBuilder.SetNameEn(*input.NameEn)
	}

	return queryBuilder.Save(ctx)
}

// CreateDishAlias is the resolver for the createDishAlias field.
func (r *mutationResolver) CreateDishAlias(ctx context.Context, input models.CreateDishAliasInput) (*ent.DishAlias, error) {
	return r.Database.DishAlias.Create().
		SetID(input.AliasName).
		SetNormalizedAliasName(input.NormalizedAliasName).
		SetDishID(input.Dish).
		Save(ctx)
}

// DeleteDishAlias is the resolver for the deleteDishAlias field.
func (r *mutationResolver) DeleteDishAlias(ctx context.Context, input models.DeleteDishAliasInput) (*ent.DishAlias, error) {
	dishAlias, err := r.Database.DishAlias.Get(ctx, input.AliasName)
	if err != nil {
		return nil, err
	}

	err = r.Database.DishAlias.DeleteOne(dishAlias).Exec(ctx)
	if err != nil {
		return nil, err
	}

	return dishAlias, nil
}

// CreateOccurrence is the resolver for the createOccurrence field.
func (r *mutationResolver) CreateOccurrence(ctx context.Context, input models.CreateOccurrenceInput) (*ent.Occurrence, error) {
	queryBuilder := r.Database.Occurrence.Create().
		SetLocationID(input.Location).
		SetDishID(input.Dish).
		SetDate(input.Date).
		SetStatus(input.Status)

	if len(input.SideDishes) > 0 {
		queryBuilder = queryBuilder.AddSideDishIDs(input.SideDishes...)
	}

	if input.Kj != nil {
		queryBuilder = queryBuilder.SetKj(*input.Kj)
	}

	if input.Kcal != nil {
		queryBuilder = queryBuilder.SetKcal(*input.Kcal)
	}

	if input.Fat != nil {
		queryBuilder = queryBuilder.SetFat(*input.Fat)
	}

	if input.SaturatedFat != nil {
		queryBuilder = queryBuilder.SetSaturatedFat(*input.SaturatedFat)
	}

	if input.Carbohydrates != nil {
		queryBuilder = queryBuilder.SetCarbohydrates(*input.Carbohydrates)
	}

	if input.Sugar != nil {
		queryBuilder = queryBuilder.SetSugar(*input.Sugar)
	}

	if input.Fiber != nil {
		queryBuilder = queryBuilder.SetFiber(*input.Fiber)
	}

	if input.Protein != nil {
		queryBuilder = queryBuilder.SetProtein(*input.Protein)
	}

	if input.Salt != nil {
		queryBuilder = queryBuilder.SetSalt(*input.Salt)
	}

	if input.PriceStudent != nil {
		queryBuilder = queryBuilder.SetPriceStudent(*input.PriceStudent)
	}

	if input.PriceStaff != nil {
		queryBuilder = queryBuilder.SetPriceStaff(*input.PriceStaff)
	}

	if input.PriceGuest != nil {
		queryBuilder = queryBuilder.SetPriceGuest(*input.PriceGuest)
	}

	if len(input.Tags) > 0 {
		queryBuilder = queryBuilder.AddTagIDs(input.Tags...)
	}

	return queryBuilder.Save(ctx)
}

// UpdateOccurrence is the resolver for the updateOccurrence field.
func (r *mutationResolver) UpdateOccurrence(ctx context.Context, input models.UpdateOccurrenceInput) (*ent.Occurrence, error) {
	queryBuilder := r.Database.Occurrence.UpdateOneID(input.ID)

	if input.Dish != nil {
		queryBuilder = queryBuilder.SetDishID(*input.Dish)
	}

	if input.Date != nil {
		queryBuilder = queryBuilder.SetDate(*input.Date)
	}

	if input.Status != nil {
		queryBuilder = queryBuilder.SetStatus(*input.Status)
	}

	if input.Kj != nil {
		queryBuilder = queryBuilder.SetKj(*input.Kj)
	}

	if input.Kcal != nil {
		queryBuilder = queryBuilder.SetKcal(*input.Kcal)
	}

	if input.Fat != nil {
		queryBuilder = queryBuilder.SetFat(*input.Fat)
	}

	if input.SaturatedFat != nil {
		queryBuilder = queryBuilder.SetSaturatedFat(*input.SaturatedFat)
	}

	if input.Carbohydrates != nil {
		queryBuilder = queryBuilder.SetCarbohydrates(*input.Carbohydrates)
	}

	if input.Sugar != nil {
		queryBuilder = queryBuilder.SetSugar(*input.Sugar)
	}

	if input.Fiber != nil {
		queryBuilder = queryBuilder.SetFiber(*input.Fiber)
	}

	if input.Protein != nil {
		queryBuilder = queryBuilder.SetProtein(*input.Protein)
	}

	if input.Salt != nil {
		queryBuilder = queryBuilder.SetSalt(*input.Salt)
	}

	if input.PriceStudent != nil {
		queryBuilder = queryBuilder.SetPriceStudent(*input.PriceStudent)
	}

	if input.PriceStaff != nil {
		queryBuilder = queryBuilder.SetPriceStaff(*input.PriceStaff)
	}

	if input.PriceGuest != nil {
		queryBuilder = queryBuilder.SetPriceGuest(*input.PriceGuest)
	}

	return queryBuilder.Save(ctx)
}

// DeleteOccurrence is the resolver for the deleteOccurrence field.
func (r *mutationResolver) DeleteOccurrence(ctx context.Context, input models.DeleteOccurrenceInput) (*ent.Occurrence, error) {
	occurrence, err := r.Database.Occurrence.Get(ctx, input.ID)
	if err != nil {
		return nil, err
	}

	err = r.Database.Occurrence.DeleteOne(occurrence).Exec(ctx)
	if err != nil {
		return nil, err
	}

	return occurrence, nil
}

// AddTagToOccurrence is the resolver for the addTagToOccurrence field.
func (r *mutationResolver) AddTagToOccurrence(ctx context.Context, input models.AddTagToOccurrenceInput) (*models.OccurrenceTag, error) {
	occurrence, err := r.Database.Occurrence.Get(ctx, input.Occurrence)
	if err != nil {
		return nil, err
	}

	tag, err := r.Database.Tag.Query().Where(tag.IDEQ(input.Tag)).Only(ctx)
	if err != nil {
		return nil, err
	}

	err = r.Database.Occurrence.UpdateOne(occurrence).AddTags(tag).Exec(ctx) // only difference to remove tag
	if err != nil {
		return nil, err
	}

	return &models.OccurrenceTag{
		Occurrence: occurrence,
		Tag:        tag,
	}, nil
}

// RemoveTagFromOccurrence is the resolver for the removeTagFromOccurrence field.
func (r *mutationResolver) RemoveTagFromOccurrence(ctx context.Context, input models.RemoveTagFromOccurrenceInput) (*models.OccurrenceTag, error) {
	occurrence, err := r.Database.Occurrence.Get(ctx, input.Occurrence)
	if err != nil {
		return nil, err
	}

	tag, err := r.Database.Tag.Query().Where(tag.IDEQ(input.Tag)).Only(ctx)
	if err != nil {
		return nil, err
	}

	err = r.Database.Occurrence.UpdateOne(occurrence).RemoveTags(tag).Exec(ctx) // only difference to add tag
	if err != nil {
		return nil, err
	}

	return &models.OccurrenceTag{
		Occurrence: occurrence,
		Tag:        tag,
	}, nil
}

// AddSideDishToOccurrence is the resolver for the addSideDishToOccurrence field.
func (r *mutationResolver) AddSideDishToOccurrence(ctx context.Context, input models.AddSideDishToOccurrenceInput) (*models.OccurrenceSideDish, error) {
	occurrence, err := r.Database.Occurrence.Get(ctx, input.Occurrence)
	if err != nil {
		return nil, err
	}

	sideDish, err := r.Database.Dish.Get(ctx, input.Dish)
	if err != nil {
		return nil, err
	}

	err = r.Database.Occurrence.UpdateOne(occurrence).AddSideDishes(sideDish).Exec(ctx) // only difference to remove side dish
	if err != nil {
		return nil, err
	}

	return &models.OccurrenceSideDish{
		Occurrence: occurrence,
		Dish:       sideDish,
	}, nil
}

// RemoveSideDishFromOccurrence is the resolver for the removeSideDishFromOccurrence field.
func (r *mutationResolver) RemoveSideDishFromOccurrence(ctx context.Context, input models.RemoveSideDishFromOccurrenceInput) (*models.OccurrenceSideDish, error) {
	occurrence, err := r.Database.Occurrence.Get(ctx, input.Occurrence)
	if err != nil {
		return nil, err
	}

	sideDish, err := r.Database.Dish.Get(ctx, input.Dish)
	if err != nil {
		return nil, err
	}

	err = r.Database.Occurrence.UpdateOne(occurrence).RemoveSideDishes(sideDish).Exec(ctx) // only difference to add side dish
	if err != nil {
		return nil, err
	}

	return &models.OccurrenceSideDish{
		Occurrence: occurrence,
		Dish:       sideDish,
	}, nil
}

// CreateReview is the resolver for the createReview field.
func (r *mutationResolver) CreateReview(ctx context.Context, input models.CreateReviewInput) (*ent.Review, error) {
	tx, err := r.Database.Tx(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to start transaction: %w", err)
	}

	queryBuilder := tx.Review.Create().
		SetOccurrenceID(input.Occurrence).
		SetStars(input.Stars)

	if input.DisplayName != nil {
		queryBuilder = queryBuilder.SetDisplayName(*input.DisplayName)
	}

	if input.Text != nil {
		queryBuilder = queryBuilder.SetText(*input.Text)
	}

	review, err := queryBuilder.Save(ctx)
	if err != nil {
		txErr := tx.Rollback()
		if txErr != nil {
			return nil, fmt.Errorf("failed to rollback transaction: %w", txErr)
		}
		return nil, err
	}

	// Process & store images
	if input.Images != nil {
		_, err := r.storeImages(tx, ctx, review, input.Images) // only require error (if no error: images are stored)
		if err != nil {
			txErr := tx.Rollback()
			if txErr != nil {
				return nil, fmt.Errorf("failed to rollback transaction: %w", txErr)
			}
			return nil, err
		}
	}

	err = tx.Commit()
	if err != nil {
		return nil, fmt.Errorf("failed to commit transaction: %w", err)
	}

	return review, err
}

// UpdateReview is the resolver for the updateReview field.
func (r *mutationResolver) UpdateReview(ctx context.Context, input models.UpdateReviewInput) (*ent.Review, error) {
	review, err := r.Database.Review.Get(ctx, input.ID)
	if err != nil {
		return nil, err
	}

	queryBuilder := r.Database.Review.UpdateOne(review)

	if input.Occurrence != nil {
		queryBuilder = queryBuilder.SetOccurrenceID(*input.Occurrence)
	}

	if input.DisplayName != nil {
		queryBuilder = queryBuilder.SetDisplayName(*input.DisplayName)
	}

	if input.Stars != nil {
		queryBuilder = queryBuilder.SetStars(*input.Stars)
	}

	if input.Text != nil {
		queryBuilder = queryBuilder.SetText(*input.Text)
	}

	if input.Approved != nil {
		queryBuilder = queryBuilder.SetAcceptedAt(time.Now()) // approved (not null) at current time
	}

	return queryBuilder.Save(ctx)
}

// DeleteReview is the resolver for the deleteReview field.
func (r *mutationResolver) DeleteReview(ctx context.Context, input models.DeleteReviewInput) (*ent.Review, error) {
	review, err := r.Database.Review.Get(ctx, input.ID)
	if err != nil {
		return nil, err
	}

	images, err := review.QueryImages().All(ctx)
	if err != nil {
		return nil, err
	}

	err = r.Database.Review.DeleteOne(review).Exec(ctx)
	if err != nil {
		return nil, err
	}

	err = r.deleteImages(ctx, images) // TODO: if a single image fails to delete, there will be remaining images in the fs
	if err != nil {
		return nil, err
	}

	return review, nil
}

// AddImagesToReview is the resolver for the addImagesToReview field.
func (r *mutationResolver) AddImagesToReview(ctx context.Context, input models.AddImagesToReviewInput) (*ent.Review, error) {
	// start a transaction
	tx, err := r.Database.Tx(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to start transaction: %w", err)
	}

	// get the review we want to add images to
	review, err := tx.Review.Get(ctx, input.Review)
	if err != nil {
		txErr := tx.Rollback()
		if txErr != nil {
			return nil, fmt.Errorf("failed to rollback transaction: %w", txErr)
		}
		return nil, err
	}

	// store the images
	images, err := r.storeImages(tx, ctx, review, input.Images)
	if err != nil {
		txErr := tx.Rollback()
		if txErr != nil {
			return nil, fmt.Errorf("failed to rollback transaction: %w", txErr)
		}
		return nil, err
	}

	// add the images to the review (in db)
	updatedReview, err := review.Update().AddImages(images...).Save(ctx)
	if err != nil {
		txErr := tx.Rollback()
		if txErr != nil {
			return nil, fmt.Errorf("failed to rollback transaction: %w", txErr)
		}
		return nil, err
	}

	// commit the transaction
	err = tx.Commit()
	if err != nil {
		return nil, fmt.Errorf("failed to commit transaction: %w", err)
	}

	return updatedReview, nil
}

// DeleteImageFromReview is the resolver for the deleteImageFromReview field.
func (r *mutationResolver) DeleteImageFromReview(ctx context.Context, input models.DeleteImageToReviewInput) (*ent.Image, error) {
	review, err := r.Database.Review.Get(ctx, input.Review)
	if err != nil {
		return nil, err
	}

	image, err := r.Database.Image.Get(ctx, input.ID)
	if err != nil {
		return nil, err
	}

	err = r.deleteImages(ctx, []*ent.Image{image})
	if err != nil {
		return nil, err
	}

	return image, review.Update().RemoveImages(image).Exec(ctx) // todo: perhaps use transaction here or keep images on error
}

// Mutation returns gqlserver.MutationResolver implementation.
func (r *Resolver) Mutation() gqlserver.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
