package grpc

import (
	"context"

	food "github.com/moromin/PFC-balancer/services/food/proto"
	"github.com/moromin/PFC-balancer/services/menu/proto"
	recipe "github.com/moromin/PFC-balancer/services/recipe/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var _ proto.MenuServiceServer = (*server)(nil)

type server struct {
	foodClient   food.FoodServiceClient
	recipeClient recipe.RecipeServiceClient
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

func (s *server) CreateRecipe(ctx context.Context, req *proto.CreateRecipeRequest) (*proto.CreateRecipeResponse, error) {
	// TODO: validate user
	res, err := s.recipeClient.CreateRecipe(ctx, &recipe.CreateRecipeRequest{
		Name:        req.Name,
		FoodAmounts: req.FoodAmounts,
		Procedures:  req.Procedures,
		UserId:      req.UserId,
	})

	if err != nil {
		st, ok := status.FromError(err)
		if ok {
			return nil, status.Error(st.Code(), err.Error())
		}
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &proto.CreateRecipeResponse{
		Id: res.Id,
	}, nil
}

func (s *server) FindRecipeById(ctx context.Context, req *proto.FindRecipeByIdRequest) (*proto.FindRecipeByIdResponse, error) {
	res, err := s.recipeClient.FindRecipeById(ctx, &recipe.FindRecipeByIdRequest{
		Id: req.Id,
	})

	if err != nil {
		st, ok := status.FromError(err)
		if ok {
			return nil, status.Error(st.Code(), err.Error())
		}
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &proto.FindRecipeByIdResponse{
		Recipe: res.GetRecipe(),
	}, nil
}

func (s *server) ListRecipes(ctx context.Context, req *proto.ListRecipesRequest) (*proto.ListRecipesResponse, error) {
	res, err := s.recipeClient.ListRecipes(ctx, &recipe.ListRecipesRequest{})

	if err != nil {
		st, ok := status.FromError(err)
		if ok {
			return nil, status.Error(st.Code(), err.Error())
		}
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &proto.ListRecipesResponse{
		Recipes: res.GetRecipes(),
	}, nil
}
