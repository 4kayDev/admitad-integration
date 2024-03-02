package service

import (
	"context"
	"encoding/json"
	"errors"

	pb "github.com/4kayDev/admitad-integration/internal/generated/proto/admitad_integration"
	"github.com/4kayDev/admitad-integration/internal/pkg/clients/admitad"
	"github.com/4kayDev/admitad-integration/internal/pkg/clients/admitad/models"
	"github.com/4kayDev/admitad-integration/internal/pkg/storage"
	"github.com/4kayDev/admitad-integration/internal/pkg/storage/sql"
	"github.com/4kayDev/admitad-integration/internal/utils/jsoner"
	"github.com/dr3dnought/gospadi"
)

func (s *Service) DeleteOffer(ctx context.Context, input *DeleteOfferInput) (*pb.Offer, error) {
	offer, err := s.storage.DeleteOffer(ctx, &sql.DeleteOfferInput{
		ID: input.ID,
	})
	if err != nil {
		return nil, err
	}

	return &pb.Offer{
		Id:        offer.ID.String(),
		AdmitadId: int64(offer.AdmitadID),
		Data:      offer.Data,
		IsSaved:   true,
	}, nil
}

func (s *Service) SaveOffer(ctx context.Context, input *SaveOfferInput) (*pb.Offer, error) {
	affliliate, exerr := s.admitadClient.GetAffiliateById(&admitad.GetAffiliateByIdInput{
		AdmiatdId: input.AdmitadId,
	})
	if exerr != nil {
		return nil, exerr.Error()
	}

	data, err := json.Marshal(affliliate)
	if err != nil {
		return nil, err
	}

	offer, err := s.storage.CreateOffer(ctx, &sql.CreateOfferInput{
		AdmitadID:   affliliate.Id,
		Data:        string(data),
		Name:        affliliate.Name,
		Description: affliliate.Description,
		Link:        affliliate.SiteURL,
		SharedValue: 0,
	})
	if err != nil {
		if errors.Is(err, storage.ErrEntityExists) {
			return nil, ErrEntityExists
		}

		return nil, err
	}

	return &pb.Offer{
		Id:          offer.ID.String(),
		AdmitadId:   int64(offer.AdmitadID),
		SharedValue: int32(offer.ShareValue),
		Name:        offer.Name,
		Description: offer.Description,
		Data:        offer.Data,
		IsSaved:     true,
	}, nil
}

func (s *Service) UpdateOffer(ctx context.Context, input *UpdateOfferInput) (*pb.Offer, error) {
	offer, err := s.storage.UpdateOffer(ctx, &sql.UpdateOfferInput{
		ID:          input.ID,
		Name:        input.Name,
		Description: input.Description,
		SharedValue: input.SharedValue,
	})
	if err != nil {
		if errors.Is(err, storage.ErrEntityNotFound) {
			return nil, ErrNotFound
		}

		return nil, err
	}

	return &pb.Offer{
		Id:          offer.ID.String(),
		AdmitadId:   int64(offer.AdmitadID),
		SharedValue: int32(offer.ShareValue),
		Name:        offer.Name,
		Description: offer.Description,
		Data:        offer.Data,
		IsSaved:     true,
	}, nil
}

func (s *Service) GetSavedOffers(ctx context.Context, input *PaginationInput) ([]*pb.Offer, error) {
	offers, err := s.storage.FindOffers(ctx, &sql.FindOffersInput{
		Limit:  input.Limit,
		Offset: input.Offset,
	})
	if err != nil {
		return nil, err
	}

	result := make([]*pb.Offer, 0, len(offers))
	for _, e := range offers {
		result = append(result, &pb.Offer{
			Id:          e.ID.String(),
			AdmitadId:   int64(e.AdmitadID),
			SharedValue: 0,
			Name:        e.Name,
			Description: e.Description,
			Data:        e.Data,
			IsSaved:     true,
		})
	}

	return result, nil
}

func (s *Service) GetOffers(ctx context.Context, input *GetOffersInput) ([]*pb.Offer, error) {
	affiliates, exerr := s.admitadClient.GetAffiliates(&admitad.GetAffiliatesInput{
		Limit:  input.Limit,
		Offset: input.Offset,
	})
	if exerr != nil {
		return nil, exerr.Error()
	}

	savedOffers, err := s.storage.FindOffersByAdmitadID(ctx, &sql.FindOffersByAdmitadIDInput{
		IDs: gospadi.Map(affiliates, func(m models.Affiliate) int {
			return m.Id
		}),
	})
	if err != nil {
		return nil, err
	}

	offers := make([]*pb.Offer, 0, len(affiliates))
	for _, e := range affiliates {
		isSaved := false
		name := ""
		description := ""
		id := ""
		for _, o := range savedOffers {
			if o.AdmitadID == e.Id {
				isSaved = true
				id = o.ID.String()
				name = e.Name
				description = e.Description
			}
		}

		offers = append(offers, &pb.Offer{
			Id:          id,
			AdmitadId:   int64(e.Id),
			SharedValue: 0,
			Data:        jsoner.Jsonify(e),
			Name:        name,
			Description: description,
			IsSaved:     isSaved,
		})
	}

	return offers, nil
}

func (s *Service) InitLink(ctx context.Context, input *InitLinkInput) (string, error) {
	offer, err := s.storage.FindOffer(ctx, &sql.FindOfferInput{ID: input.ID})
	if err != nil {
		if errors.Is(err, storage.ErrEntityNotFound) {
			return "", ErrNotFound
		}
	}

	_, err = s.storage.CreateClick(ctx, &sql.CreateClickInput{
		RequestId: input.RequestId,
		OfferID:   input.ID,
	})
	if err != nil {
		if errors.Is(err, storage.ErrEntityExists) {
			return "", ErrEntityExists
		}
		return "", err
	}

	return offer.Link, nil
}
