package controllers

import (
	"fmt"
	"net/http"

	"github.com/moromin/go-svelte/backend/domain"
	"github.com/moromin/go-svelte/backend/interface/database"
	"github.com/moromin/go-svelte/backend/usecase"
)

type UserController struct {
	Interactor usecase.UserInteractor
}

func NewUserController(sqlhandler database.SQLHandler) *UserController {
	return &UserController{
		Interactor: usecase.UserInteractor{
			UserRepository: &database.UserRepository{
				Conn: sqlhandler,
			},
		},
	}
}

func (controller *UserController) Create(c Context) error {
	u := domain.User{}
	if err := c.Bind(&u); err != nil {
		return APICustomError(c, http.StatusInternalServerError, fmt.Sprintf("Bind user error: %v", err))
	}

	hashedPassword, err := HashPassword(u.HashedPassword)
	if err != nil {
		return APICustomError(c, http.StatusInternalServerError, fmt.Sprintf("Hashing password error: %v", err))
	}

	u.HashedPassword = hashedPassword

	user, err := controller.Interactor.Add(u)
	if err != nil {
		return APICustomError(c, http.StatusInternalServerError, fmt.Sprintf("Create user error: %v", err))
	}

	return c.JSON(http.StatusCreated, user)
}
