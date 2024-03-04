package sql

import (
	"context"
	"errors"

	"github.com/4kayDev/admitad-integration/internal/pkg/models"
	"github.com/4kayDev/admitad-integration/internal/pkg/storage"
	"github.com/4kayDev/admitad-integration/internal/utils/ref"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type CreateOfferInput struct {
	AdmitadID   int
	Data        string
	Name        string
	Description string
	ImageURL    string
	Link        string
	SharedValue int
}

func (s *Storage) CreateOffer(ctx context.Context, input *CreateOfferInput) (*models.Offer, error) {

	offer := &models.Offer{
		AdmitadID:   input.AdmitadID,
		Name:        input.Name,
		Description: input.Description,
		ImageURL:    input.ImageURL,
		ShareValue:  input.SharedValue,
		Data:        input.Data,
		Link:        input.Link,
		IsHidden:    ref.Ref[bool](true),
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

type FindOffersInput struct {
	Limit    int
	Offset   int
	IsHidden bool
}

func (s *Storage) FindOffers(ctx context.Context, input *FindOffersInput) ([]*models.Offer, error) {
	offers := make([]*models.Offer, 0)

	tr := s.getter.DefaultTrOrDB(ctx, s.db).WithContext(ctx)
	err := tr.Model(&offers).Where("is_hidden = ?", input.IsHidden).Offset(input.Offset).Limit(input.Limit + 1).Find(&offers).Error
	if err != nil {
		return nil, err
	}

	return offers, nil
}

type FindOffersByAdmitadIDInput struct {
	IDs []int
}

func (s *Storage) FindOffersByAdmitadID(ctx context.Context, input *FindOffersByAdmitadIDInput) ([]*models.Offer, error) {
	offers := make([]*models.Offer, 0)

	tr := s.getter.DefaultTrOrDB(ctx, s.db).WithContext(ctx)
	err := tr.Order("created_at ASC").Where("admitad_id IN ?", input.IDs).Find(&offers).Error
	if err != nil {
		return nil, err
	}

	return offers, nil
}

type FinOfferByNameOrDescriptionInput struct {
	Name string
}

func (s *Storage) FindOfferByNameOrDescription(ctx context.Context, input *FinOfferByNameOrDescriptionInput) ([]*models.Offer, error) {
	offers := make([]*models.Offer, 0)

	tr := s.getter.DefaultTrOrDB(ctx, s.db).WithContext(ctx)
	err := tr.Where("name LIKE ? OR description LIKE ?", "%"+input.Name+"%", "%"+input.Name+"%").Find(&offers).Error
	if err != nil {
		return nil, err
	}
	return offers, nil
}

type UpdateOfferInput struct {
	ID          uuid.UUID
	Name        string
	Description string
	ImageURL    string
	IsHidden    *bool
	SharedValue int
}

func (s *Storage) UpdateOffer(ctx context.Context, input *UpdateOfferInput) (*models.Offer, error) {
	offer := &models.Offer{
		ID:          input.ID,
		AdmitadID:   0,
		Name:        input.Name,
		Description: input.Description,
		ImageURL:    input.ImageURL,
		ShareValue:  input.SharedValue,
		IsHidden:    input.IsHidden,
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
