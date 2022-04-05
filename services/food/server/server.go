package server

import (
	"context"
	"net/http"

	"github.com/moromin/PFC-balancer/services/food/db"
	"github.com/moromin/PFC-balancer/services/food/models"
	"github.com/moromin/PFC-balancer/services/food/proto"
)

type Server struct {
	H db.Handler
}

const createFood = `
INSERT INTO food (
	name,
	protein,
	fat,
	carbohydrate,
	category
) VALUES (
	$1, $2, $3, $4, $5
) RETURNING *;
`

func (s *Server) CreateFood(ctx context.Context, req *proto.CreateFoodRequest) (*proto.CreateFoodResponse, error) {
	ins, err := s.H.DB.PrepareContext(ctx, createFood)
	if err != nil {
		return &proto.CreateFoodResponse{
			Status: http.StatusInternalServerError,
			Error:  err.Error(),
		}, nil
	}

	res, err := ins.ExecContext(ctx, req.Name, req.Protein, req.Fat, req.Carbohydrate, req.Category)
	if err != nil {
		return &proto.CreateFoodResponse{
			Status: http.StatusConflict,
			Error:  err.Error(),
		}, nil
	}

	id, err := res.LastInsertId()
	if err != nil {
		return &proto.CreateFoodResponse{
			Status: http.StatusConflict,
			Error:  err.Error(),
		}, nil
	}

	return &proto.CreateFoodResponse{
		Status: http.StatusCreated,
		Id:     id,
	}, nil
}

const findOne = `
SELECT *
FROM food
WHERE name = $1;
`

func (s *Server) FindOne(ctx context.Context, req *proto.FindOneRequest) (*proto.FindOneResponse, error) {
	var food models.Food

	if err := s.H.DB.QueryRowContext(ctx, findOne, req.Name).Scan(&food); err != nil {
		return &proto.FindOneResponse{
			Status: http.StatusNotFound,
			Error:  err.Error(),
		}, nil
	}

	data := &proto.FindOneData{
		Id:           food.Id,
		Name:         food.Name,
		Protein:      food.Protein,
		Fat:          food.Fat,
		Carbohydrate: food.Carbohydrate,
		Category:     food.Category,
	}

	return &proto.FindOneResponse{
		Status: http.StatusOK,
		Data:   data,
	}, nil
}
