package routes

import (
	"go-ecommerce/controllers"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

func New() *echo.Echo {
	e := echo.New()
	e.Validator = &CustomValidator{validator: validator.New()}

	//start provinces
	e.GET("/provinces", controllers.GetAllProvincesController)
	e.GET("/provinces/:id", controllers.GetProvincesByIDController)
	e.DELETE("/provinces/:id", controllers.DeleteProvincesByIDController)
	e.POST("/provinces", controllers.StoreProvincesController)
	e.PUT("/provinces/:id", controllers.PutProvincesByIDController)
	//end provinces
	//start city
	e.GET("/city/:id", controllers.GetCityByProvincesIdController)
	e.GET("/get-city/:id", controllers.GetCityByIdController)
	e.POST("/city", controllers.StoreCityController)
	e.DELETE("/delete-city/:id", controllers.DeleteCityController)
	e.PUT("/update-city/:id", controllers.PutCityByIDController)

	// end city

	//start users
	e.GET("/users", controllers.GetUserControllers)
	e.POST("/users", controllers.AddUserControllers)
	//end users
	return e
}
