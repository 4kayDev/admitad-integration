package service

import (
	"context"

	"github.com/4kayDev/admitad-integration/internal/pkg/clients/admitad"
	"github.com/4kayDev/admitad-integration/internal/pkg/clients/admitad/models"
)

type ToggleSaveOfferInput struct {
	AdmitadID int
	Limit     int
	Offset    int
}

func (s *Service) ToggleSaveOffer(ctx context.Context, input *ToggleSaveOfferInput) (*ToggleSaveOfferOutput, error) {
	affilates, err := s.admitadClient.GetAffiliates(&admitad.GetAffiliatesInput{
		Limit:  input.Limit,
		Offset: input.Offset,
	})
	if err != nil {
		return nil, err
	}
	return &ToggleSaveOfferOutput{
		Affiliates: affilates,
	}, nil
}

type GetSavedOffersInput struct {
	Limit  int
	Offset int
}
type GetSavedOffersOutput struct {
	Affiliates []models.Affiliate
}

func (s *Service) GetSavedOffers(ctx context.Context, input *GetSavedOffersInput) (*GetSavedOffersOutput, error) {
	affilates, err := s.admitadClient.GetAffiliates(&admitad.GetAffiliatesInput{
		Limit:  input.Limit,
		Offset: input.Offset,
	})
	if err != nil {
		return nil, err
	}
	return &GetSavedOffersOutput{
		Affiliates: affilates,
	}, nil
}

func (s *Service) GetOffers(ctx context.Context, input *GetSavedOffersInput) (*GetSavedOffersOutput, error) {
	return nil, nil
}
