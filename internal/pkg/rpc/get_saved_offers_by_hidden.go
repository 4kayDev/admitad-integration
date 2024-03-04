package rpc

import (
	"context"

	pb "github.com/4kayDev/admitad-integration/internal/generated/proto/admitad_integration"
	"github.com/4kayDev/admitad-integration/internal/pkg/service"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) GetSavedOffersByHidden(ctx context.Context, req *pb.GetSavedOffersByHiddenRequest) (*pb.GetSavedOffersByHiddenResponse, error) {
	offers, err := s.service.GetSavedOffersByHidden(ctx, &service.GetSavedOffersByHiddenInput{
		PaginationInput: service.PaginationInput{
			Limit:  int(req.GetLimit()),
			Offset: int(req.GetOffset()),
		},
		IsHidden: req.IsHidden,
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &pb.GetSavedOffersByHiddenResponse{
		Offers: offers,
	}, err
}
