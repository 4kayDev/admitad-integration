package rpc

import (
	"context"

	pb "github.com/4kayDev/admitad-integration/internal/generated/proto/admitad_integration"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) GetOfferByNameOrDescription(ctx context.Context, req *pb.GetOfferByNameOrDescriptionRequest) (*pb.GetOfferByNameOrDescriptionResponse, error) {
	offers, err := s.service.FindOfferByNameOrDescription(ctx, req.Name)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &pb.GetOfferByNameOrDescriptionResponse{
		Offers: offers,
	}, nil
}
