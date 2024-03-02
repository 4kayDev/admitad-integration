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

func (s *Server) GetSavedOffer(ctx context.Context, req *pb.GetSavedOfferRequest) (*pb.GetSavedOfferResponse, error) {
	id, err := uuid.Parse(req.GetId())
	if err != nil {
		return &pb.GetSavedOfferResponse{ErrorCode: pb.GetSavedOfferErrorCode_GET_SAVED_OFFER_ERROR_CODE_VALIDATION}, nil
	}

	offer, err := s.service.GetOffer(ctx, &service.GetOfferInput{ID: id})
	if err != nil {
		if errors.Is(err, service.ErrNotFound) {
			return &pb.GetSavedOfferResponse{ErrorCode: pb.GetSavedOfferErrorCode_GET_SAVED_OFFER_ERROR_CODE_NOT_FOUND}, nil
		}

		return nil, status.Error(codes.Internal, err.Error())
	}

	return &pb.GetSavedOfferResponse{
		Offer: offer,
	}, nil
}
