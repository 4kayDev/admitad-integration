package rpc

import (
	"context"

	pb "github.com/4kayDev/admitad-integration/internal/generated/proto/admitad_integration"
	"github.com/4kayDev/admitad-integration/internal/pkg/service"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) GetSavedOffers(ctx context.Context, req *pb.GetSavedOffersRequest) (*pb.GetSavedOffersResponse, error) {
	offers, err := s.service.GetSavedOffers(ctx, &service.PaginationInput{
		Limit:  int(req.GetLimit()),
		Offset: int(req.GetOffset()),
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &pb.GetSavedOffersResponse{
		Offers: offers,
	}, err
}
