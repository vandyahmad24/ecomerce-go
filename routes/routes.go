package routes

import (
	"go-ecommerce/constants"
	"go-ecommerce/controllers"

	echoMid "go-ecommerce/middlewares"

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
	//start login register
	l := e.Group("/api")
	l.POST("/register", controllers.AddUserControllers)
	l.POST("/login", controllers.LoginController)
	// untuk recreated OTP
	r.GET("/otp/resend", controllers.ResendOTPController)
	// verifikasi OTP
	r.POST("/otp/verif", controllers.VerifOTPController)

	isAdmin := r.Group("/admin")
	isAdmin.Use(echoMid.IsAdmin)

	// untuk user yg sudah aktif
	isActive := r.Group("")
	isActive.Use(echoMid.IsActive)

	// hanya test
	isAdmin.GET("/test", controllers.CobaSegalaController)
	//end login register

	//start provinces
	r.GET("/provinces", controllers.GetAllProvincesController)
	r.GET("/provinces/:id", controllers.GetProvincesByIDController)
	isAdmin.DELETE("/provinces/:id", controllers.DeleteProvincesByIDController)
	isAdmin.POST("/provinces", controllers.StoreProvincesController)
	isAdmin.PUT("/provinces/:id", controllers.PutProvincesByIDController)
	//end provinces
	//start city
	r.GET("/city/:id", controllers.GetCityByProvincesIdController)
	r.GET("/get-city/:id", controllers.GetCityByIdController)
	isAdmin.POST("/city", controllers.StoreCityController)
	isAdmin.DELETE("/delete-city/:id", controllers.DeleteCityController)
	isAdmin.PUT("/update-city/:id", controllers.PutCityByIDController)
	// end city
	// start user
	r.GET("/user", controllers.GetUserByTokenControllers)
	// end user

	// categori
	r.GET("/category/active", controllers.GetAllCategoryActive)
	isAdmin.GET("/category/inactive", controllers.GetAllCategoryInActive)
	isActive.POST("/category", controllers.InsertCategory)
	r.GET("/category/:id", controllers.GetCategoryByID)
	isAdmin.PUT("/category/:id", controllers.UpdateCategoryByID)
	// end categori

	// produk
	isActive.POST("/produk/create", controllers.CreateProdukController)
	// get produk by id
	isActive.GET("/produk/:id", controllers.GetProdukByIdController)
	// get produk by category id
	isActive.GET("/produk/category/:id", controllers.GetProdukByCategoryController)
	// produk ke shopping cart
	isActive.POST("/add-cart/:product_id", controllers.AddCartByProdukProdukController)
	// list produk dari cart
	isActive.GET("/carts", controllers.GetListCartController)
	// delete produk dari cart
	isActive.DELETE("/carts/:cart_id", controllers.DeleteCartController)
	// CART CHECKOUT
	isActive.POST("/checkout", controllers.CheckoutController)

	return e
}
