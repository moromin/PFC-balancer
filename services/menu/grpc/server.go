package grpc

import (
	"context"

	food "github.com/moromin/PFC-balancer/services/food/proto"
	"github.com/moromin/PFC-balancer/services/menu/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var _ proto.MenuServiceServer = (*server)(nil)

type server struct {
	foodClient food.FoodServiceClient
}

func (s *server) FindFoodById(ctx context.Context, req *proto.FindFoodByIdRequest) (*proto.FindFoodByIdResponse, error) {
	res, err := s.foodClient.FindFoodById(ctx, &food.FindFoodByIdRequest{
		Id: req.Id,
	})
	if err != nil {
		st, ok := status.FromError(err)
		if ok {
			return nil, status.Error(st.Code(), err.Error())
		}
		return nil, status.Error(codes.Internal, err.Error())
	}

	food := res.GetFood()

	return &proto.FindFoodByIdResponse{
		Food: &proto.Food{
			Id:           food.Id,
			Name:         food.Name,
			Protein:      food.Protein,
			Fat:          food.Fat,
			Carbohydrate: food.Carbohydrate,
			Category:     food.Category,
		},
	}, nil
}

func (s *server) ListFoods(ctx context.Context, req *proto.ListFoodsRequest) (*proto.ListFoodsResponse, error) {
	foodRes, err := s.foodClient.ListFoods(ctx, &food.ListFoodsRequest{})
	if err != nil {
		st, ok := status.FromError(err)
		if ok {
			return nil, status.Error(st.Code(), err.Error())
		}
		return nil, status.Error(codes.Internal, err.Error())
	}

	res := &proto.ListFoodsResponse{
		FoodList: make([]*proto.Food, len(foodRes.FoodList)),
	}

	for i, food := range foodRes.FoodList {
		res.FoodList[i] = &proto.Food{
			Id:           food.Id,
			Name:         food.Name,
			Protein:      food.Protein,
			Fat:          food.Fat,
			Carbohydrate: food.Carbohydrate,
			Category:     food.Category,
		}
	}

	return res, nil
}

func (s *server) SearchFoods(ctx context.Context, req *proto.SearchFoodsRequest) (*proto.SearchFoodsResponse, error) {
	foodRes, err := s.foodClient.SearchFoods(ctx, &food.SearchFoodsRequest{
		Name: req.Name,
	})
	if err != nil {
		st, ok := status.FromError(err)
		if ok {
			return nil, status.Error(st.Code(), err.Error())
		}
		return nil, status.Error(codes.Internal, err.Error())
	}

	res := &proto.SearchFoodsResponse{
		FoodList: make([]*proto.Food, len(foodRes.FoodList)),
	}

	for i, food := range foodRes.FoodList {
		res.FoodList[i] = &proto.Food{
			Id:           food.Id,
			Name:         food.Name,
			Protein:      food.Protein,
			Fat:          food.Fat,
			Carbohydrate: food.Carbohydrate,
			Category:     food.Category,
		}
	}

	return res, nil
}
