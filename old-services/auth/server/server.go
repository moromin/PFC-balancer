package server

import (
	"context"
	"net/http"

	"github.com/moromin/PFC-balancer/services/auth/db"
	"github.com/moromin/PFC-balancer/services/auth/models"
	"github.com/moromin/PFC-balancer/services/auth/proto"
	"github.com/moromin/PFC-balancer/services/auth/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Server struct {
	H   db.Handler
	Jwt utils.JwtWrapper
}

const registerUser = `
INSERT INTO users (
	email,
	password
) VALUES (
	$1, $2
)
`

func (s *Server) Register(ctx context.Context, req *proto.RegisterRequest) (*proto.RegisterResponse, error) {
	req.Password = utils.HashPassword(req.Password)

	ins, err := s.H.DB.PrepareContext(ctx, registerUser)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to prepare query:", err)
	}

	_, err = ins.ExecContext(ctx, req.Email, req.Password)
	if err != nil {
		return nil, status.Errorf(codes.AlreadyExists, "%s already exists", req.Email)
	}

	return &proto.RegisterResponse{
		Status: http.StatusCreated,
	}, nil
}

const findUser = `
SELECT *
FROM users
WHERE email = $1
`

func (s *Server) Login(ctx context.Context, req *proto.LoginRequest) (*proto.LoginResponse, error) {
	var user models.User

	row := s.H.DB.QueryRowContext(ctx, findUser, req.Email)
	err := row.Scan(
		&user.Id,
		&user.Email,
		&user.Password,
	)

	if err != nil {
		return nil, status.Error(codes.NotFound, "user not found")
	}

	if match := utils.CheckPasswordHash(req.Password, user.Password); !match {
		return nil, status.Error(codes.NotFound, "user not found")
	}

	token, err := s.Jwt.GenerateToken(user)

	if err != nil {
		return nil, status.Error(codes.Internal, "failed to generate token")
	}

	return &proto.LoginResponse{
		Status: http.StatusOK,
		Token:  token,
	}, nil
}

func (s *Server) Validate(ctx context.Context, req *proto.ValidateRequest) (*proto.ValidateResponse, error) {
	claims, err := s.Jwt.ValidateToken(req.Token)

	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "failed to validate token")
	}

	var user models.User

	row := s.H.DB.QueryRowContext(ctx, findUser, claims.Email)
	err = row.Scan(
		&user.Id,
		&user.Email,
		&user.Password,
	)

	if err != nil {
		return nil, status.Errorf(codes.NotFound, "%s is not found", claims.Email)
	}

	return &proto.ValidateResponse{
		Status: http.StatusOK,
		UserId: user.Id,
	}, nil
}
