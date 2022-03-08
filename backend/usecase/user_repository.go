package usecase

import (
	"github.com/moromin/go-svelte/backend/domain"
)

type UserRepository interface {
	Store(u domain.User) (domain.User, error)
}
