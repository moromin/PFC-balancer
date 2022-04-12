package db

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"github.com/moromin/PFC-balancer/platform/db/config"
	"github.com/moromin/PFC-balancer/platform/db/models"
)

var _ DB = (*dbWrapper)(nil)

var (
	ErrNotFound      = errors.New("not found")
	ErrAlreadyExists = errors.New("already exists")
)

type DB interface {
	CreateUser(ctx context.Context, email, password string) (*models.User, error)
	FindUserByEmail(ctx context.Context, email string) (*models.User, error)

	FindFoodById(ctx context.Context, id int64) (*models.Food, error)
	ListFoods(ctx context.Context) ([]*models.Food, error)
	SearchFoods(ctx context.Context, name string) ([]*models.Food, error)
}

type dbWrapper struct {
	db *sql.DB
}

func New(c *config.Config) (DB, error) {
	db, err := sql.Open("postgres", c.GetDBUrl())
	if err != nil {
		return nil, err
	}

	return &dbWrapper{db: db}, nil
}

const createUser = `
INSERT INTO users (
	email,
	password
) VALUES (
	$1, $2
)
`

func (w *dbWrapper) CreateUser(ctx context.Context, email, password string) (*models.User, error) {
	_, err := w.db.ExecContext(ctx, createUser, email, password)
	if err != nil {
		return nil, ErrAlreadyExists
	}

	user, err := w.FindUserByEmail(ctx, email)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return user, nil
}

const findUserByEmail = `
SELECT *
FROM users
WHERE email = $1
`

func (w *dbWrapper) FindUserByEmail(ctx context.Context, email string) (*models.User, error) {
	var user models.User

	if err := w.db.QueryRowContext(ctx, findUserByEmail, email).Scan(&user.Id, &user.Email, &user.Password); err != nil {
		return nil, ErrNotFound
	}

	return &user, nil
}

const findFoodById = `
SELECT *
FROM foods
WHERE id = $1
`

func (w *dbWrapper) FindFoodById(ctx context.Context, id int64) (*models.Food, error) {
	var food models.Food

	if err := w.db.QueryRowContext(ctx, findFoodById, id).Scan(&food.Id, &food.Name, &food.Protein, &food.Fat, &food.Carbohydrate, &food.Category); err != nil {
		return nil, ErrNotFound
	}

	return &food, nil
}

const listFoods = `
SELECT *
FROM foods
ORDER BY id ASC
`

func (w *dbWrapper) ListFoods(ctx context.Context) ([]*models.Food, error) {
	var foodList []*models.Food

	rows, err := w.db.QueryContext(ctx, listFoods)
	if err != nil {
		return nil, fmt.Errorf("failed to query: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var food models.Food

		err := rows.Scan(&food.Id, &food.Name, &food.Protein, &food.Fat, &food.Carbohydrate, &food.Category)
		if err != nil {
			return nil, fmt.Errorf("failed to scan: %w", err)
		}

		foodList = append(foodList, &food)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("failed to close: %w", err)
	}

	return foodList, nil
}

const searchFoods = `
	SELECT *
	FROM foods
	WHERE name LIKE $1
	ORDER BY id ASC;
`

func (w *dbWrapper) SearchFoods(ctx context.Context, name string) ([]*models.Food, error) {
	var foodList []*models.Food

	rows, err := w.db.QueryContext(ctx, searchFoods, "%"+name+"%")
	if err != nil {
		return nil, fmt.Errorf("failed to query: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var food models.Food

		err := rows.Scan(&food.Id, &food.Name, &food.Protein, &food.Fat, &food.Carbohydrate, &food.Category)
		if err != nil {
			return nil, fmt.Errorf("failed to scan: %w", err)
		}

		foodList = append(foodList, &food)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("failed to close: %w", err)
	}

	return foodList, nil
}
