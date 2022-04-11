package grpc

import (
	"context"

	db "github.com/moromin/PFC-balancer/platform/db/proto"
	"github.com/moromin/PFC-balancer/services/user/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// check implementation server methods
var _ proto.UserServiceServer = (*server)(nil)

type server struct {
	dbClient db.DBServiceClient
}

func (s *server) CreateUser(ctx context.Context, req *proto.CreateUserRequest) (*proto.CreateUserResponse, error) {
	res, err := s.dbClient.CreateUser(ctx, &db.CreateUserRequest{
		Email:    req.Email,
		Password: req.Password,
	})
	if err != nil {
		st, ok := status.FromError(err)
		if ok && st.Code() == codes.AlreadyExists {
			return nil, status.Errorf(codes.AlreadyExists, "%s already exists", req.Email)
		}
		return nil, status.Errorf(codes.Internal, "internal error")
	}

	user := res.GetUser()

	return &proto.CreateUserResponse{
		User: &proto.User{
			Id:    user.Id,
			Email: user.Email,
		},
	}, nil
}

func (s *server) FindUserByEmail(ctx context.Context, req *proto.FindUserByEmailRequest) (*proto.FindUserByEmailResponse, error) {
	res, err := s.dbClient.FindUserByEmail(ctx, &db.FindUserByEmailRequest{
		Email: req.Email,
	})
	if err != nil {
		st, ok := status.FromError(err)
		if ok && st.Code() == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "%s not found", req.Email)
		}
		return nil, status.Errorf(codes.Internal, "internal error")
	}

	user := res.GetUser()

	return &proto.FindUserByEmailResponse{
		User: &proto.User{
			Id:    user.Id,
			Email: user.Email,
		},
		Password: user.Password,
	}, nil
}
