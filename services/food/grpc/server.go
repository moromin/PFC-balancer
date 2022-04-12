package grpc

import (
	"context"

	db "github.com/moromin/PFC-balancer/platform/db/proto"
	"github.com/moromin/PFC-balancer/services/food/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var _ proto.FoodServiceServer = (*server)(nil)

type server struct {
	dbClient db.DBServiceClient
}

func (s *server) FindFoodById(ctx context.Context, req *proto.FindFoodByIdRequest) (*proto.FindFoodByIdResponse, error) {
	res, err := s.dbClient.FindFoodById(ctx, &db.FindFoodByIdRequest{
		Id: req.Id,
	})
	if err != nil {
		st, ok := status.FromError(err)
		if ok && st.Code() == codes.NotFound {
			return nil, status.Error(codes.AlreadyExists, err.Error())
		}
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &proto.FindFoodByIdResponse{
		Food: res.GetFood(),
	}, nil
}

func (s *server) ListFoods(ctx context.Context, req *proto.ListFoodsRequest) (*proto.ListFoodsResponse, error) {
	res, err := s.dbClient.ListFoods(ctx, &db.ListFoodsRequest{})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &proto.ListFoodsResponse{
		FoodList: res.GetFoodList(),
	}, nil
}

func (s *server) SearchFoods(ctx context.Context, req *proto.SearchFoodsRequest) (*proto.SearchFoodsResponse, error) {
	res, err := s.dbClient.SearchFoods(ctx, &db.SearchFoodsRequest{
		Name: req.Name,
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &proto.SearchFoodsResponse{
		FoodList: res.GetFoodList(),
	}, nil
}
