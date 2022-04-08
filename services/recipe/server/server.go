package server

import (
	"context"

	"github.com/moromin/PFC-balancer/services/recipe/db"
	"github.com/moromin/PFC-balancer/services/recipe/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Server struct {
	H db.Handler
}

const registerRecipe = `
INSERT INTO recipes (
	name,
	user_id
) VALUES (
	$1, $2
) RETURNING id;
`

const registerProcedure = `
INSERT INTO procedures (
	text,
	recipe_id
) VALUES (
	$1, $2
)
`

const registerRecipeFood = `
INSERT INTO recipe_food (
	recipe_id,
	food_id,
	amount
) VALUES (
	$1, $2, $3
)
`

func (s *Server) CreateRecipe(ctx context.Context, req *proto.CreateRecipeRequest) (*proto.CreateRecipeResponse, error) {
	var recipe_id int64

	row := s.H.DB.QueryRowContext(ctx, registerRecipe, req.Data.Name, req.UserId)
	if err := row.Scan(&recipe_id); err != nil {
		return nil, status.Errorf(codes.AlreadyExists, "%s already exists", req.Data.Name)
	}

	for i := 0; i < len(req.Data.ProcedureList); i++ {
		_, err := s.H.DB.ExecContext(ctx, registerProcedure, req.Data.ProcedureList[i], recipe_id)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "failed to register procedure %q", req.Data.ProcedureList[i])
		}
	}

	for i := 0; i < len(req.Data.FoodAmount); i++ {
		_, err := s.H.DB.ExecContext(ctx, registerRecipeFood, recipe_id, req.Data.FoodAmount[i].FoodId, req.Data.FoodAmount[i].Amount)
		if err != nil {
			return nil, status.Error(codes.Internal, "failed to register recipe_food")
		}
	}

	return &proto.CreateRecipeResponse{
		Id: recipe_id,
	}, nil
}

const findRecipe = `
SELECT name
FROM recipes
WHERE id = $1 AND user_id = $2
`

const findProcedure = `
SELECT text
FROM procedures
WHERE recipe_id = $1
`

const findRecipeFood = `
SELECT food_id, amount
FROM recipe_food
WHERE recipe_id = $1
`

func (s *Server) ReadRecipe(ctx context.Context, req *proto.ReadRecipeRequest) (*proto.ReadRecipeResponse, error) {
	var data proto.RecipeData

	// findRecipe
	if err := s.H.DB.QueryRowContext(ctx, findRecipe, req.Id, req.UserId).Scan(&data.Name); err != nil {
		return nil, status.Error(codes.InvalidArgument, "recipe or user is not found")
	}

	// findProcedures
	rows, err := s.H.DB.QueryContext(ctx, findProcedure, req.Id)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "%q procedure is not found", data.Name)
	}
	defer rows.Close()

	for rows.Next() {
		var procedure string
		if err := rows.Scan(&procedure); err != nil {
			return nil, status.Errorf(codes.DataLoss, "database scan error")
		}
		data.ProcedureList = append(data.ProcedureList, procedure)
	}
	if err := rows.Err(); err != nil {
		return nil, status.Error(codes.Internal, "some error during load database")
	}

	// findRecipeFood
	rows, err = s.H.DB.QueryContext(ctx, findRecipeFood, req.Id)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "%q food is not found", data.Name)
	}
	defer rows.Close()

	for rows.Next() {
		var foodAmount proto.FoodAmount
		if err := rows.Scan(&foodAmount.FoodId, &foodAmount.Amount); err != nil {
			return nil, status.Errorf(codes.DataLoss, "database scan error")
		}
		data.FoodAmount = append(data.FoodAmount, &foodAmount)
	}
	if err := rows.Err(); err != nil {
		return nil, status.Error(codes.Internal, "some error during load database")
	}

	return &proto.ReadRecipeResponse{
		Data: &data,
	}, nil
}
