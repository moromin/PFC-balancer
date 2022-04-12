package grpc

import (
	"context"

	db "github.com/moromin/PFC-balancer/platform/db/proto"
	food "github.com/moromin/PFC-balancer/services/food/proto"
	"github.com/moromin/PFC-balancer/services/recipe/proto"
	user "github.com/moromin/PFC-balancer/services/user/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var _ proto.RecipeServiceServer = (*server)(nil)

type server struct {
	dbClient   db.DBServiceClient
	userClient user.UserServiceClient
	foodClient food.FoodServiceClient
}

func (s *server) CreateRecipe(ctx context.Context, req *proto.CreateRecipeRequest) (*proto.CreateRecipeResponse, error) {
	res, err := s.dbClient.CreateRecipe(ctx, &db.CreateRecipeRequest{
		Name:        req.Name,
		FoodAmounts: req.FoodAmounts,
		Procedures:  req.Procedures,
		UserId:      req.UserId,
	})
	if err != nil {
		st, ok := status.FromError(err)
		if ok && st.Code() == codes.AlreadyExists {
			return nil, status.Errorf(codes.AlreadyExists, "%s already exists", req.Name)
		}
		return nil, status.Errorf(codes.Internal, "internal error")
	}

	return &proto.CreateRecipeResponse{
		Id: res.Id,
	}, nil
}

func (s *server) FindRecipeById(ctx context.Context, req *proto.FindRecipeByIdRequest) (*proto.FindRecipeByIdResponse, error) {
	dbRes, err := s.dbClient.FindRecipeById(ctx, &db.FindRecipeByIdRequest{
		Id: req.Id,
	})
	if err != nil {
		st, ok := status.FromError(err)
		if ok && st.Code() == codes.NotFound {
			return nil, status.Error(codes.NotFound, "not found")
		}
		return nil, status.Errorf(codes.Internal, "internal error")
	}

	r := dbRes.GetRecipe()

	userRes, err := s.userClient.FindUserById(ctx, &user.FindUserByIdRequest{
		Id: r.UserId,
	})
	if err != nil {
		st, ok := status.FromError(err)
		if ok && st.Code() == codes.NotFound {
			return nil, status.Error(codes.NotFound, "not found")
		}
		return nil, status.Errorf(codes.Internal, "internal error")
	}

	foodPFCAmounts := make([]*proto.FoodPFCAmount, len(r.FoodAmounts))
	for i, fa := range r.FoodAmounts {
		foodRes, err := s.foodClient.FindFoodById(ctx, &food.FindFoodByIdRequest{
			Id: fa.FoodId,
		})
		if err != nil {
			st, ok := status.FromError(err)
			if ok && st.Code() == codes.NotFound {
				return nil, status.Error(codes.NotFound, "not found")
			}
			return nil, status.Errorf(codes.Internal, "internal error")
		}
		foodPFCAmounts[i] = &proto.FoodPFCAmount{
			Food:   foodRes.GetFood(),
			Amount: fa.Amount,
		}
	}

	return &proto.FindRecipeByIdResponse{
		Recipe: &proto.Recipe{
			Id:             r.Id,
			RecipeName:     r.Name,
			FoodPfcAmounts: foodPFCAmounts,
			Procedures:     r.Procedures,
			UserName:       userRes.User.Email,
		},
	}, nil
}

func (s *server) ListRecipes(ctx context.Context, req *proto.ListRecipesRequest) (*proto.ListRecipesResponse, error) {
	res, err := s.dbClient.ListRecipes(ctx, &db.ListRecipesRequest{})
	if err != nil {
		st, ok := status.FromError(err)
		if ok && st.Code() == codes.NotFound {
			return nil, status.Error(codes.NotFound, "not found")
		}
		return nil, status.Errorf(codes.Internal, "internal error")
	}

	rs := res.GetRecipes()

	recipes := make([]*proto.Recipe, len(rs))
	for i, r := range rs {
		recipe, err := s.FindRecipeById(ctx, &proto.FindRecipeByIdRequest{
			Id: r.Id,
		})
		if err != nil {
			return nil, err
		}
		recipes[i] = recipe.GetRecipe()
	}

	return &proto.ListRecipesResponse{
		Recipes: recipes,
	}, nil
}
