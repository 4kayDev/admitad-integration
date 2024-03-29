package service

import (
	"context"

	"github.com/4kayDev/admitad-integration/internal/pkg/clients/admitad"
	"github.com/4kayDev/admitad-integration/internal/pkg/models"
	"github.com/4kayDev/admitad-integration/internal/pkg/storage/sql"
)

type Service struct {
	storage       Storage
	admitadClient *admitad.Client
}
type Storage interface {
	CreateOffer(ctx context.Context, input *sql.CreateOfferInput) (*models.Offer, error)
	FindOffer(ctx context.Context, input *sql.FindOfferInput) (*models.Offer, error)
	FindOffersByAdmitadID(ctx context.Context, input *sql.FindOffersByAdmitadIDInput) ([]*models.Offer, error)
	FindOffers(ctx context.Context, input *sql.FindOffersInput) ([]*models.Offer, error)
	UpdateOffer(ctx context.Context, input *sql.UpdateOfferInput) (*models.Offer, error)
	DeleteOffer(ctx context.Context, input *sql.DeleteOfferInput) (*models.Offer, error)
	CreateClick(ctx context.Context, input *sql.CreateClickInput) (*models.Click, error)
	FindOfferByNameOrDescription(ctx context.Context, input *sql.FinOfferByNameOrDescriptionInput) ([]*models.Offer, error)
	FindOffersByHidden(ctx context.Context, input *sql.FindOffersByHiddenInput) ([]*models.Offer, error)
}

func NewService(storage Storage, admitadClient *admitad.Client) *Service {
	return &Service{storage: storage, admitadClient: admitadClient}
}
