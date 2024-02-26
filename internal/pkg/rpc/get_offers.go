package rpc

import (
	"context"

	pb "github.com/4kayDev/admitad-integration/internal/generated/proto/admitad_integration"
)

func (s *Server) GetOffers(ctx context.Context, req *pb.GetOffersRequest) (*pb.GetOffersResponse, error) {
	return &pb.GetOffersResponse{}, nil
}
