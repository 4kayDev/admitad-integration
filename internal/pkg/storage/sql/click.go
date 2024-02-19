package sql

import (
	"context"
	"errors"

	"github.com/4kayDev/admitad-integration/internal/pkg/models"
	"github.com/4kayDev/admitad-integration/internal/pkg/storage"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CreateClickInput struct {
	OfferID uuid.UUID
}

func (s *Storage) CreateClick(ctx context.Context, input *CreateClickInput) (*models.Click, error) {
	click := &models.Click{
		OfferID: input.OfferID,
	}

	tr := s.getter.DefaultTrOrDB(ctx, s.db).WithContext(ctx)
	err := tr.Create(click).Error
	if err != nil {
		switch {
		case errors.Is(err, gorm.ErrForeignKeyViolated):
			return nil, storage.ErrForeignKey
		default:
			return nil, err
		}
	}

	return click, nil
}
