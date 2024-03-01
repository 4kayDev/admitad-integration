package rpc

import (
	"context"
	"errors"

	pb "github.com/4kayDev/admitad-integration/internal/generated/proto/admitad_integration"
	"github.com/4kayDev/admitad-integration/internal/pkg/service"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) InitLink(ctx context.Context, req *pb.InitLinkRequest) (*pb.InitLinkResponse, error) {
	id, err := uuid.Parse(req.GetId())
	if err != nil {
		return &pb.InitLinkResponse{
			ErrorCode: pb.InitLinkErrorCode_INIT_LINK_ERROR_CODE_VALIDATION,
		}, nil
	}

	requestId, err := uuid.Parse(req.GetRequestId())
	if err != nil {
		return &pb.InitLinkResponse{
			ErrorCode: pb.InitLinkErrorCode_INIT_LINK_ERROR_CODE_VALIDATION,
		}, nil
	}

	link, err := s.service.InitLink(ctx, &service.InitLinkInput{
		RequestId: requestId,
		ID:        id,
	})
	if err != nil {
		switch {
		case errors.Is(err, service.ErrNotFound):
			return &pb.InitLinkResponse{ErrorCode: pb.InitLinkErrorCode_INIT_LINK_ERROR_CODE_NOT_FOUND}, nil
		case errors.Is(err, service.ErrEntityExists):
			return &pb.InitLinkResponse{ErrorCode: pb.InitLinkErrorCode_INIT_LINK_ERROR_CODE_ALREADY_EXISTS}, nil
		default:
			return nil, status.Error(codes.Internal, err.Error())
		}
	}

	return &pb.InitLinkResponse{Link: link}, nil
}
