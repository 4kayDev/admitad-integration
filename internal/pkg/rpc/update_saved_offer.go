package rpc

import (
	"context"
	"errors"

	pb "github.com/4kayDev/admitad-integration/internal/generated/proto/admitad_integration"
	"github.com/4kayDev/admitad-integration/internal/pkg/service"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) UpdateSavedOffer(ctx context.Context, req *pb.UpdateSavedOfferRequest) (*pb.UpdateSavedOfferResponse, error) {
	id, err := uuid.Parse(req.GetId())
	if err != nil {
		return &pb.UpdateSavedOfferResponse{ErrorCode: pb.UpdateSavedOfferErroCode_UPDATE_SAVED_OFFER_ERROR_CODE_VALIDATION}, nil
	}

	offer, err := s.service.UpdateOffer(ctx, &service.UpdateOfferInput{
		ID:          id,
		Name:        req.GetName(),
		Description: req.GetDescription(),
		SharedValue: int(req.GetSharedValue()),
	})
	if err != nil {
		switch {
		case errors.Is(err, service.ErrNotFound):
			return &pb.UpdateSavedOfferResponse{ErrorCode: pb.UpdateSavedOfferErroCode_UPDATE_SAVED_OFFER_ERROR_CODE_NOT_FOUND}, nil
		default:
			return nil, status.Error(codes.Internal, err.Error())
		}
	}

	return &pb.UpdateSavedOfferResponse{Offer: offer}, nil
}
