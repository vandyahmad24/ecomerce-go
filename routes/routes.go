package routes

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo"
	"go-ecommerce/controllers"
)


type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

func New() *echo.Echo{
	e := echo.New()
	e.Validator = &CustomValidator{validator: validator.New()}
	e.GET("/users",controllers.GetUserControllers)
	e.POST("/users",controllers.AddUserControllers)
	return e
}
