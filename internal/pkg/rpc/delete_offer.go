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

func (s *Server) DeleteOffer(ctx context.Context, req *pb.DeleteOfferRequest) (*pb.DeleteOfferResponse, error) {
	id, err := uuid.Parse(req.Id)
	if err != nil {
		return &pb.DeleteOfferResponse{
			ErrorCode: pb.DeleteOfferErrorCode_DELETE_OFFER_ERROR_CODE_VALIDATION,
		}, nil
	}

	offer, err := s.service.DeleteOffer(ctx, &service.DeleteOfferInput{
		ID: id,
	})
	if err != nil {
		switch {
		case errors.Is(err, service.ErrNotFound):
			return &pb.DeleteOfferResponse{ErrorCode: pb.DeleteOfferErrorCode_DELETE_OFFER_ERROR_CODE_NOT_FOUND}, nil
		default:
			return nil, status.Error(codes.Internal, err.Error())
		}
	}

	return &pb.DeleteOfferResponse{
		Offer: offer,
	}, nil
}
