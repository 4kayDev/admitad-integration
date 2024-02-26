package rpc

import (
	"context"

	pb "github.com/4kayDev/admitad-integration/internal/generated/proto/admitad_integration"
)

func (s *Server) UpdateSavedOffer(ctx context.Context, req *pb.UpdateSavedOfferRequest) (*pb.UpdateSavedOfferResponse, error) {
	return &pb.UpdateSavedOfferResponse{}, nil
}
