package routes

import (
	"github.com/labstack/echo"
	"go-ecommerce/controllers"
)

func New() *echo.Echo{
	e := echo.New()
	e.GET("/users",controllers.GetUserControllers)
	return e
}
