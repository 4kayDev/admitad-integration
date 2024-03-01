package rpc

import (
	"context"

	pb "github.com/4kayDev/admitad-integration/internal/generated/proto/admitad_integration"
	"github.com/4kayDev/admitad-integration/internal/pkg/service"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) GetOffers(ctx context.Context, req *pb.GetOffersRequest) (*pb.GetOffersResponse, error) {
	offers, err := s.service.GetOffers(ctx, &service.GetOffersInput{
		PaginationInput: service.PaginationInput{
			Limit:  int(req.GetLimit()),
			Offset: int(req.GetOffset()),
		},
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &pb.GetOffersResponse{
		Offers: offers,
	}, nil
}
