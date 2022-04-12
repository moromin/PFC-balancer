package grpc

import (
	"context"

	"github.com/moromin/PFC-balancer/platform/db/models"
	"github.com/moromin/PFC-balancer/services/auth/proto"
	"github.com/moromin/PFC-balancer/services/auth/utils"
	user "github.com/moromin/PFC-balancer/services/user/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var _ proto.AuthServiceServer = (*server)(nil)

type server struct {
	userClient user.UserServiceClient
	jwt        utils.JwtWrapper
}

func (s *server) Register(ctx context.Context, req *proto.RegisterRequest) (*proto.RegisterResponse, error) {
	res, err := s.userClient.CreateUser(ctx, &user.CreateUserRequest{
		Email:    req.Email,
		Password: utils.HashPassword(req.Password),
	})
	if err != nil {
		st, ok := status.FromError(err)
		if ok && st.Code() == codes.AlreadyExists {
			return nil, status.Errorf(codes.AlreadyExists, "%s already exists", req.Email)
		}
		return nil, status.Errorf(codes.Internal, "internal error")
	}

	user := res.GetUser()

	return &proto.RegisterResponse{
		User: user,
	}, nil
}

func (s *server) Login(ctx context.Context, req *proto.LoginRequest) (*proto.LoginResponse, error) {
	res, err := s.userClient.FindUserByEmail(ctx, &user.FindUserByEmailRequest{
		Email: req.Email,
	})
	if err != nil {
		st, ok := status.FromError(err)
		if ok && st.Code() == codes.NotFound {
			return nil, status.Error(codes.NotFound, "not found")
		}
		return nil, status.Error(codes.Internal, "internal error")
	}

	if match := utils.CheckPasswordHash(req.Password, res.Password); !match {
		return nil, status.Error(codes.NotFound, "not found")
	}

	user := res.GetUser()
	token, err := s.jwt.GenerateToken(models.User{
		Id:    user.Id,
		Email: user.Email,
	})
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to generate token")
	}

	return &proto.LoginResponse{
		Token: token,
	}, nil
}
