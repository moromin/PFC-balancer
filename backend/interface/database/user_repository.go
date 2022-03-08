package database

import (
	"github.com/moromin/go-svelte/backend/domain"
)

type UserRepository struct {
	Conn SQLHandler
}

const store = `-- name: Store :one
INSERT INTO users (
    nick_name,
    email,
    hashed_password
) VALUES (
    $1, $2, $3
) RETURNING *
`

func (repo *UserRepository) Store(u domain.User) (domain.User, error) {
	row := repo.Conn.QueryRow(store, u.NickName, u.Email, u.HashedPassword)
	var user domain.User
	err := row.Scan(
		&user.UserID,
		&user.NickName,
		&user.Email,
		&user.HashedPassword,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	return user, err
}
