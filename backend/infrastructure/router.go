package infrastructure

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/moromin/go-svelte/backend/interface/controllers"
)

func Serve() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	userController := controllers.NewUserController(NewSQLHandler())

	e.POST("/signup", func(c echo.Context) error { return userController.Create(c) })

	e.Logger.Fatal(e.Start(":8000"))
}
