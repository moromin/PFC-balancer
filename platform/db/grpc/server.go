package grpc

import (
	"context"
	"errors"

	"github.com/moromin/PFC-balancer/platform/db/db"
	"github.com/moromin/PFC-balancer/platform/db/models"
	"github.com/moromin/PFC-balancer/platform/db/proto"
	food "github.com/moromin/PFC-balancer/services/food/proto"
	recipe "github.com/moromin/PFC-balancer/services/recipe/proto"
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

func (s *server) CreateRecipe(ctx context.Context, req *proto.CreateRecipeRequest) (*proto.CreateRecipeResponse, error) {
	foodAmounts := make([]*models.FoodAmount, len(req.FoodAmounts))
	for i, fa := range foodAmounts {
		foodAmounts[i] = &models.FoodAmount{
			FoodId: fa.FoodId,
			Amount: fa.Amount,
		}
	}

	id, err := s.db.CreateRecipe(ctx, req.Name, foodAmounts, req.Procedures, req.UserId)
	if err != nil {
		if errors.Is(err, db.ErrAlreadyExists) {
			return nil, status.Error(codes.AlreadyExists, "")
		} else if errors.Is(err, db.ErrNotFound) {
			return nil, status.Error(codes.NotFound, "")
		}
		return nil, status.Error(codes.Internal, "internal error")
	}

	return &proto.CreateRecipeResponse{
		Id: id,
	}, nil
}

func (s *server) FindRecipeById(ctx context.Context, req *proto.FindRecipeByIdRequest) (*proto.FindRecipeByIdResponse, error) {
	r, err := s.db.FindRecipeById(ctx, req.Id)
	if err != nil {
		if errors.Is(err, db.ErrNotFound) {
			return nil, status.Error(codes.NotFound, "")
		}
		return nil, status.Error(codes.Internal, "internal error")
	}

	foodAmounts := make([]*recipe.FoodAmount, len(r.FoodAmounts))
	for i, fa := range r.FoodAmounts {
		foodAmounts[i] = &recipe.FoodAmount{
			FoodId: fa.FoodId,
			Amount: fa.Amount,
		}
	}

	return &proto.FindRecipeByIdResponse{
		Recipe: &recipe.Recipe{
			Id:          r.Id,
			Name:        r.Name,
			FoodAmounts: foodAmounts,
			Procedures:  r.Procedures,
			UserId:      r.UserId,
		},
	}, nil
}

func (s *server) ListRecipes(ctx context.Context, req *proto.ListRecipesRequest) (*proto.ListRecipesResponse, error) {
	recipes, err := s.db.ListRecipes(ctx)
	if err != nil {
		if errors.Is(err, db.ErrNotFound) {
			return nil, status.Error(codes.NotFound, "")
		}
		return nil, status.Error(codes.Internal, "internal error")
	}

	res := make([]*recipe.Recipe, len(recipes))
	for i, r := range recipes {
		foodAmounts := make([]*recipe.FoodAmount, len(r.FoodAmounts))
		for j, fa := range r.FoodAmounts {
			foodAmounts[j] = &recipe.FoodAmount{
				FoodId: fa.FoodId,
				Amount: fa.Amount,
			}
		}
		res[i] = &recipe.Recipe{
			Id:          r.Id,
			Name:        r.Name,
			FoodAmounts: foodAmounts,
			Procedures:  r.Procedures,
			UserId:      r.UserId,
		}
	}

	return &proto.ListRecipesResponse{
		Recipes: res,
	}, nil
}
