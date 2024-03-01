package rpc

import (
	pb "github.com/4kayDev/admitad-integration/internal/generated/proto/admitad_integration"
	"github.com/4kayDev/admitad-integration/internal/pkg/service"
)

type Server struct {
	*pb.UnimplementedAdmitadIntegrationServer
	service *service.Service
}

func NewServer(service *service.Service) *Server {
	return &Server{service: service}
}
