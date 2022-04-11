package grpc

import (
	"context"
	"errors"

	"github.com/moromin/PFC-balancer/platform/db/db"
	"github.com/moromin/PFC-balancer/platform/db/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var _ proto.DBServiceServer = (*server)(nil)

type server struct {
	db db.DB
}

func (s *server) CreateUser(ctx context.Context, req *proto.CreateUserRequest) (*proto.CreateUserResponse, error) {
	u, err := s.db.CreateUser(ctx, req.Email, req.Password)
	if err != nil {
		if errors.Is(err, db.ErrAlreadyExists) {
			return nil, status.Error(codes.AlreadyExists, "")
		}
		return nil, status.Error(codes.Internal, "internal error")
	}

	return &proto.CreateUserResponse{
		User: &proto.User{
			Id:    u.Id,
			Email: u.Email,
		},
	}, nil
}

func (s *server) FindUserByEmail(ctx context.Context, req *proto.FindUserByEmailRequest) (*proto.FindUserByEmailResponse, error) {
	u, err := s.db.FindUserByEmail(ctx, req.Email)
	if err != nil {
		if errors.Is(err, db.ErrAlreadyExists) {
			return nil, status.Error(codes.AlreadyExists, "")
		}
		return nil, status.Error(codes.Internal, "internal error")
	}

	return &proto.FindUserByEmailResponse{
		User: &proto.User{
			Id:    u.Id,
			Email: u.Email,
		},
	}, nil
}
