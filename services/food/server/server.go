package server

import (
	"context"
	"log"
	"net/http"

	"github.com/moromin/PFC-balancer/services/food/db"
	"github.com/moromin/PFC-balancer/services/food/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Server struct {
	H db.Handler
}

const findOne = `
SELECT *
FROM foods
WHERE name = $1
`

func (s *Server) FindOne(ctx context.Context, req *proto.FindOneRequest) (*proto.FindOneResponse, error) {
	var data proto.FoodData

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

func (s *Server) ListFoods(ctx context.Context, req *proto.ListFoodsRequest) (*proto.ListFoodsResponse, error) {
	var foodList []*proto.FoodData

	rows, err := s.H.DB.QueryContext(ctx, listFood)
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to query")
	}
	defer rows.Close()

	for rows.Next() {
		var data proto.FoodData

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

	return &proto.ListFoodsResponse{
		Status:   http.StatusOK,
		FoodList: foodList,
	}, nil
}

const searchFoods = `
	SELECT *
	FROM foods
	WHERE name LIKE $1
	ORDER BY id ASC;
`

func (s *Server) SearchFoods(ctx context.Context, req *proto.SearchFoodsRequest) (*proto.SearchFoodsResponse, error) {
	var foodList []*proto.FoodData

	rows, err := s.H.DB.QueryContext(ctx, searchFoods, "%"+req.Name+"%")
	if err != nil {
		log.Println(err)
		return nil, status.Error(codes.Internal, "failed to query")
	}
	defer rows.Close()

	for rows.Next() {
		var data proto.FoodData

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

	return &proto.SearchFoodsResponse{
		Status:   http.StatusOK,
		FoodList: foodList,
	}, nil
}
