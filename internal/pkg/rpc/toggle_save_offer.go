package rpc

import (
	"context"

	pb "github.com/4kayDev/admitad-integration/internal/generated/proto/admitad_integration"
)

func (s *Server) ToggleSaveOffer(ctx context.Context, req *pb.ToggleSaveOfferRequest) (*pb.ToggleSaveOfferResponse, error) {
	return &pb.ToggleSaveOfferResponse{}, nil
}