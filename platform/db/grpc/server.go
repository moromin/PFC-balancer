package grpc

import (
	"context"
	"errors"

	"github.com/moromin/PFC-balancer/platform/db/db"
	"github.com/moromin/PFC-balancer/platform/db/proto"
	food "github.com/moromin/PFC-balancer/services/food/proto"
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
			Id:       u.Id,
			Email:    u.Email,
			Password: u.Password,
		},
	}, nil
}

func (s *server) FindFoodById(ctx context.Context, req *proto.FindFoodByIdRequest) (*proto.FindFoodByIdResponse, error) {
	f, err := s.db.FindFoodById(ctx, req.Id)
	if err != nil {
		if errors.Is(err, db.ErrNotFound) {
			return nil, status.Error(codes.NotFound, "not found")
		}
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &proto.FindFoodByIdResponse{
		Food: &food.Food{
			Id:           f.Id,
			Name:         f.Name,
			Protein:      f.Protein,
			Fat:          f.Fat,
			Carbohydrate: f.Carbohydrate,
			Category:     f.Category,
		},
	}, nil
}

func (s *server) ListFoods(ctx context.Context, req *proto.ListFoodsRequest) (*proto.ListFoodsResponse, error) {
	fl, err := s.db.ListFoods(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	var foodList []*food.Food
	for _, f := range fl {
		food := &food.Food{
			Id:           f.Id,
			Name:         f.Name,
			Protein:      f.Protein,
			Fat:          f.Fat,
			Carbohydrate: f.Carbohydrate,
			Category:     f.Category,
		}
		foodList = append(foodList, food)
	}

	return &proto.ListFoodsResponse{
		FoodList: foodList,
	}, nil
}

func (s *server) SearchFoods(ctx context.Context, req *proto.SearchFoodsRequest) (*proto.SearchFoodsResponse, error) {
	fl, err := s.db.SearchFoods(ctx, req.Name)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	var foodList []*food.Food
	for _, f := range fl {
		food := &food.Food{
			Id:           f.Id,
			Name:         f.Name,
			Protein:      f.Protein,
			Fat:          f.Fat,
			Carbohydrate: f.Carbohydrate,
			Category:     f.Category,
		}
		foodList = append(foodList, food)
	}

	return &proto.SearchFoodsResponse{
		FoodList: foodList,
	}, nil
}
