package usecase

import (
	"github.com/moromin/go-svelte/backend/domain"
)

type UserInteractor struct {
	UserRepository UserRepository
}

func (interactor *UserInteractor) Add(u domain.User) (domain.User, error) {
	user, err := interactor.UserRepository.Store(u)
	if err != nil {
		return domain.User{}, err
	}
	return user, nil
}
