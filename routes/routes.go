package routes

import (
	"go-ecommerce/constants"
	"go-ecommerce/controllers"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
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
	r := e.Group("/api")
	r.Use(middleware.JWT([]byte(constants.SECRET_JWT)))

	//start provinces
	r.GET("/provinces", controllers.GetAllProvincesController)
	r.GET("/provinces/:id", controllers.GetProvincesByIDController)
	r.DELETE("/provinces/:id", controllers.DeleteProvincesByIDController)
	r.POST("/provinces", controllers.StoreProvincesController)
	r.PUT("/provinces/:id", controllers.PutProvincesByIDController)
	//end provinces
	//start city
	r.GET("/city/:id", controllers.GetCityByProvincesIdController)
	r.GET("/get-city/:id", controllers.GetCityByIdController)
	r.POST("/city", controllers.StoreCityController)
	r.DELETE("/delete-city/:id", controllers.DeleteCityController)
	r.PUT("/update-city/:id", controllers.PutCityByIDController)

	// end city

	//start users
	// e.GET("/users/all", controllers.GetUserControllers)
	l := e.Group("/api")
	l.POST("/register", controllers.AddUserControllers)
	l.POST("/login", controllers.LoginController)

	//end users
	return e
}
