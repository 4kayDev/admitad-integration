package rpc

import (
	"context"

	pb "github.com/4kayDev/admitad-integration/internal/generated/proto/admitad_integration"
)

func (s *Server) GetSavedOffers(ctx context.Context, req *pb.GetSavedOffersRequest) (*pb.GetSavedOffersResponse, error) {
	return &pb.GetSavedOffersResponse{}, nil
}
