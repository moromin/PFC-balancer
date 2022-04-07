package server

import (
	"context"
	"net/http"

	"github.com/moromin/PFC-balancer/services/food/db"
	"github.com/moromin/PFC-balancer/services/food/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Server struct {
	H db.Handler
}

const createFood = `
INSERT INTO foods (
	name,
	protein,
	fat,
	carbohydrate,
	category
) VALUES (
	$1, $2, $3, $4, $5
)
`

func (s *Server) CreateFood(ctx context.Context, req *proto.CreateFoodRequest) (*proto.CreateFoodResponse, error) {
	ins, err := s.H.DB.PrepareContext(ctx, createFood)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to prepare query:", err)
	}

	_, err = ins.ExecContext(ctx, req.Name, req.Protein, req.Fat, req.Carbohydrate, req.Category)
	if err != nil {
		return nil, status.Errorf(codes.AlreadyExists, "%s already exists", req.Name)
	}

	return &proto.CreateFoodResponse{
		Status: http.StatusCreated,
	}, nil
}

const findOne = `
SELECT *
FROM foods
WHERE name = $1
`

func (s *Server) FindOne(ctx context.Context, req *proto.FindOneRequest) (*proto.FindOneResponse, error) {
	var data proto.FindOneData

	row := s.H.DB.QueryRowContext(ctx, findOne, req.Name)
	err := row.Scan(
		&data.Id,
		&data.Name,
		&data.Protein,
		&data.Fat,
		&data.Carbohydrate,
		&data.Category,
	)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "%s is not found", req.Name)
	}

	return &proto.FindOneResponse{
		Status: http.StatusOK,
		Data:   &data,
	}, nil
}

const listFood = `
	SELECT *
	FROM foods
	ORDER BY id ASC;
`

func (s *Server) ListFood(ctx context.Context, req *proto.ListFoodRequest) (*proto.ListFoodResponse, error) {
	var foodList []*proto.FindOneData

	rows, err := s.H.DB.QueryContext(ctx, listFood)
	if err != nil {
		return &proto.ListFoodResponse{
			Status: http.StatusInternalServerError,
		}, err
	}
	defer rows.Close()

	for rows.Next() {
		var data proto.FindOneData

		err := rows.Scan(
			&data.Id,
			&data.Name,
			&data.Protein,
			&data.Fat,
			&data.Carbohydrate,
			&data.Category,
		)
		if err != nil {
			return nil, status.Errorf(codes.DataLoss, "database scan error")
		}

		foodList = append(foodList, &data)
	}

	if err := rows.Err(); err != nil {
		return nil, status.Errorf(codes.Internal, "some error during load database")
	}

	return &proto.ListFoodResponse{
		Status:   http.StatusOK,
		FoodList: foodList,
	}, nil
}
