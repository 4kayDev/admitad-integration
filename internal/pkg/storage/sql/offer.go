package sql

import (
	"context"
	"errors"

	"github.com/4kayDev/admitad-integration/internal/pkg/models"
	"github.com/4kayDev/admitad-integration/internal/pkg/storage"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type CreateOfferInput struct {
	AdmitadID   int
	Name        string
	Description string
	SharedValue int
	ImageURL    string
}

func (s *Storage) CreateOffer(ctx context.Context, input *CreateOfferInput) (*models.Offer, error) {
	offer := &models.Offer{
		AdmitadID:   input.AdmitadID,
		Name:        input.Name,
		Description: input.Description,
		ShareValue:  input.SharedValue,
		ImageURL:    input.ImageURL,
	}

	tr := s.getter.DefaultTrOrDB(ctx, s.db).WithContext(ctx)
	err := tr.Create(offer).Error
	if err != nil {
		switch {
		case errors.Is(err, gorm.ErrDuplicatedKey):
			return nil, storage.ErrEntityExists
		case errors.Is(err, gorm.ErrInvalidValue):
			return nil, storage.ErrInvalidValue
		default:
			return nil, err
		}
	}

	return offer, nil
}

type FindOfferInput struct {
	ID uuid.UUID
}

func (s *Storage) FindOffer(ctx context.Context, input *FindOfferInput) (*models.Offer, error) {
	offer := new(models.Offer)

	tr := s.getter.DefaultTrOrDB(ctx, s.db).WithContext(ctx)
	err := tr.First(offer, "id = ?", input.ID).Error
	if err != nil {
		switch {
		case errors.Is(err, gorm.ErrRecordNotFound):
			return nil, storage.ErrEntityNotFound
		default:
			return nil, err
		}
	}

	return offer, nil
}

type FindOffersByAdmitadIDInput struct {
	IDs []int
}

func (s *Storage) FindOffersByAdmitadID(ctx context.Context, input *FindOffersByAdmitadIDInput) ([]*models.Offer, error) {
	offers := make([]*models.Offer, 0, len(input.IDs))

	tr := s.getter.DefaultTrOrDB(ctx, s.db).WithContext(ctx)
	err := tr.Order("created_at ASC").Find(offers, input.IDs).Error
	if err != nil {
		return nil, err
	}

	return offers, nil
}

type UpdateOfferInput struct {
	ID          uuid.UUID
	Name        string
	Description string
	SharedValue int
	ImageURL    string
}

func (s *Storage) UpdateOffer(ctx context.Context, input *UpdateOfferInput) (*models.Offer, error) {
	offer := &models.Offer{
		ID:          input.ID,
		Name:        input.Name,
		Description: input.Description,
		ShareValue:  input.SharedValue,
		ImageURL:    input.ImageURL,
	}

	tr := s.getter.DefaultTrOrDB(ctx, s.db).WithContext(ctx)
	result := tr.Clauses(clause.Returning{}).Updates(offer)
	if result.Error != nil {
		switch {
		case errors.Is(result.Error, gorm.ErrInvalidValue):
			return nil, storage.ErrInvalidValue
		default:
			return nil, result.Error
		}
	}

	if result.RowsAffected == 0 {
		return nil, storage.ErrEntityNotFound
	}

	return offer, nil
}

type DeleteOfferInput struct {
	ID uuid.UUID
}

func (s *Storage) DeleteOffer(ctx context.Context, input *DeleteOfferInput) (*models.Offer, error) {
	offer := new(models.Offer)

	tr := s.getter.DefaultTrOrDB(ctx, s.db).WithContext(ctx)
	result := tr.Clauses(clause.Returning{}).Delete(offer, "id = ?", input.ID)
	if result.Error != nil {
		return nil, result.Error
	}

	if result.RowsAffected == 0 {
		return nil, storage.ErrEntityNotFound
	}

	return offer, nil
}
