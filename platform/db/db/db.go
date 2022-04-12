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
	CreateUser(context.Context, string, string) (*models.User, error)
	FindUserByEmail(context.Context, string) (*models.User, error)

	FindFoodById(context.Context, int64) (*models.Food, error)
	ListFoods(context.Context) ([]*models.Food, error)
	SearchFoods(context.Context, string) ([]*models.Food, error)

	CreateRecipe(context.Context, string, []*models.FoodAmount, []string, int64) (int64, error)
	FindRecipeById(context.Context, int64) (*models.Recipe, error)
	ListRecipes(context.Context) ([]*models.Recipe, error)
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

// user
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

// food
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

// recipe
const createRecipe = `
INSERT INTO recipes (
	name,
	user_id
) VALUES (
	$1, $2
)
`

const createProcedures = `
INSERT INTO procedures (
	text,
	recipe_id
) VALUES (
	$1, $2
)
`

const createRecipeFood = `
INSERT INTO recipe_food (
	recipe_id,
	food_id,
	amount
) VALUES (
	$1, $2, $3
)
`

const findRecipeByName = `
SELECT *
FROM recipes
WHERE name = $1
`

func (w *dbWrapper) CreateRecipe(ctx context.Context, name string, foodAmounts []*models.FoodAmount, procedures []string, userId int64) (int64, error) {
	_, err := w.db.ExecContext(ctx, createRecipe, name, userId)
	if err != nil {
		return 0, ErrAlreadyExists
	}

	var recipe models.Recipe
	if err := w.db.QueryRowContext(ctx, findRecipeByName, name).Scan(&recipe.Id, &recipe.Name, &recipe.UserId); err != nil {
		return 0, ErrNotFound
	}

	for _, procedure := range procedures {
		_, err := w.db.ExecContext(ctx, createProcedures, procedure, recipe.Id)
		if err != nil {
			return 0, fmt.Errorf("failed to create procedures: %w", err)
		}
	}

	for _, foodAmount := range foodAmounts {
		_, err := w.db.ExecContext(ctx, createRecipeFood, recipe.Id, foodAmount.FoodId, foodAmount.Amount)
		if err != nil {
			return 0, fmt.Errorf("failed to create food_recipe: %w", err)
		}
	}

	return recipe.Id, nil
}

const findRecipe = `
SELECT *
FROM recipes
WHERE id = $1
`

const findProcedures = `
SELECT text
FROM procedures
WHERE recipe_id = $1
`

const findRecipeFood = `
SELECT food_id, amount
FROM recipe_food
WHERE recipe_id = $1
`

func (w *dbWrapper) FindRecipeById(ctx context.Context, id int64) (*models.Recipe, error) {
	var recipe models.Recipe

	// Find recipe
	if err := w.db.QueryRowContext(ctx, findRecipe, id).Scan(&recipe.Id, &recipe.Name, &recipe.UserId); err != nil {
		return nil, ErrNotFound
	}

	// Find procedures
	prows, err := w.db.QueryContext(ctx, findProcedures, recipe.Id)
	if err != nil {
		return nil, ErrNotFound
	}
	defer prows.Close()

	recipe.Procedures = make([]string, 0)
	for prows.Next() {
		var procedure string
		if err := prows.Scan(&procedure); err != nil {
			return nil, fmt.Errorf("failed to scan procedure: %w", err)
		}
		recipe.Procedures = append(recipe.Procedures, procedure)
	}
	if err := prows.Err(); err != nil {
		return nil, fmt.Errorf("internal error at loading procedures: %w", err)
	}

	// Find recipe_food
	rfrows, err := w.db.QueryContext(ctx, findRecipeFood, recipe.Id)
	if err != nil {
		return nil, ErrNotFound
	}
	defer rfrows.Close()

	recipe.FoodAmounts = make([]*models.FoodAmount, 0)
	for rfrows.Next() {
		var foodAmount models.FoodAmount
		if err := rfrows.Scan(&foodAmount.FoodId, &foodAmount.Amount); err != nil {
			return nil, fmt.Errorf("failed to scan food_amunt: %w", err)
		}
		recipe.FoodAmounts = append(recipe.FoodAmounts, &foodAmount)
	}
	if err := rfrows.Err(); err != nil {
		return nil, fmt.Errorf("internal error at loading food_amount: %w", err)
	}

	return &recipe, nil
}

const listRecipes = `
SELECT id
FROM recipes
ORDER BY id ASC
`

func (w *dbWrapper) ListRecipes(ctx context.Context) ([]*models.Recipe, error) {
	rows, err := w.db.QueryContext(ctx, listRecipes)
	if err != nil {
		return nil, ErrNotFound
	}
	defer rows.Close()

	recipes := make([]*models.Recipe, 0)
	for rows.Next() {
		var id int64
		if err := rows.Scan(&id); err != nil {
			return nil, fmt.Errorf("failed to scan recipes: %w", err)
		}
		recipe, err := w.FindRecipeById(ctx, id)
		if err != nil {
			return nil, err
		}
		recipes = append(recipes, recipe)
	}

	return recipes, nil
}
