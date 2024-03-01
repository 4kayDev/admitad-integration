package rpc

import (
	"context"

	pb "github.com/4kayDev/admitad-integration/internal/generated/proto/admitad_integration"
	"github.com/4kayDev/admitad-integration/internal/pkg/service"
)

func (s *Server) SaveOffer(ctx context.Context, req *pb.SaveOfferRequest) (*pb.SaveOfferResponse, error) {
	offer, err := s.service.SaveOffer(ctx, &service.SaveOfferInput{
		AdmitadId: int(req.GetAdmitadId()),
	})
	if err != nil {
		return nil, err
	}

	return &pb.SaveOfferResponse{
		Offer: offer,
	}, err
}
