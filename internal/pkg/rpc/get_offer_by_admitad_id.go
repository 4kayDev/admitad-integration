package rpc

import (
	"context"
	"errors"

	pb "github.com/4kayDev/admitad-integration/internal/generated/proto/admitad_integration"
	"github.com/4kayDev/admitad-integration/internal/pkg/service"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) GetOfferByAdmitadId(ctx context.Context, req *pb.GetOfferByAdmitadIdRequest) (*pb.GetOfferByAdmitadIdResponse, error) {
	offer, err := s.service.GetSavedOfferByAdmitadId(ctx, &service.GetOfferByAdmitadId{AdmitadID: req.GetAdmitadId()})
	if err != nil {
		if errors.Is(err, service.ErrNotFound) {
			return &pb.GetOfferByAdmitadIdResponse{ErrorCode: pb.GetOfferByAdmitadIdErrorCode_GET_OFFER_BY_ADMITAD_ID_ERROR_CODE_NOT_FOUND}, nil
		}

		return nil, status.Error(codes.Internal, err.Error())
	}

	return &pb.GetOfferByAdmitadIdResponse{
		Offer: offer,
	}, nil
}
