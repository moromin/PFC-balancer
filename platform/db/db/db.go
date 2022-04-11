package db

import (
	"context"
	"database/sql"
	"errors"
	"log"

	_ "github.com/lib/pq"
	"github.com/moromin/PFC-balancer/platform/db/config"
	"github.com/moromin/PFC-balancer/platform/db/models"
	"github.com/moromin/PFC-balancer/platform/db/utils"
)

var _ DB = (*dbWrapper)(nil)

var (
	ErrNotFound      = errors.New("not found")
	ErrAlreadyExists = errors.New("already exists")
)

type DB interface {
	CreateUser(ctx context.Context, email, password string) (*models.User, error)
	FindUserByEmail(ctx context.Context, email string) (*models.User, error)
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
	password = utils.HashPassword(password)
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
SELECT id, email
FROM users
WHERE email = $1
`

func (w *dbWrapper) FindUserByEmail(ctx context.Context, email string) (*models.User, error) {
	var user models.User

	if err := w.db.QueryRowContext(ctx, findUserByEmail, email).Scan(&user.Id, &user.Email); err != nil {
		return nil, ErrNotFound
	}

	return &user, nil
}
