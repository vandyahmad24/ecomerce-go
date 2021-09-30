package controllers

import (
	"go-ecommerce/helpers"
	"go-ecommerce/models"

	"github.com/labstack/echo"
	"golang.org/x/crypto/bcrypt"

	//"go-ecommerce/lib"
	"go-ecommerce/lib/database"
	"net/http"
)

func GetUserControllers(c echo.Context) error {
	users, e := database.GetUsers()
	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}
	response := helpers.ResponseFormater(http.StatusOK, "success", "success get users", users)
	return c.JSON(http.StatusOK, response)
}

func AddUserControllers(c echo.Context) error {
	u := new(models.UserRequest)
	c.Bind(u)
	if err := c.Validate(u); err != nil {
		response := helpers.ResponseFormater(http.StatusBadRequest, "error", "invalid request", helpers.ErrorFormater(err))
		return echo.NewHTTPError(http.StatusBadRequest, response)
	}
	// cek apakah kota ada
	city, errorCity := database.AccessGetCityById(&models.City{}, int(u.CityId))
	if errorCity != nil {
		response := helpers.ResponseFormater(http.StatusBadRequest, "error", "city not found", nil)
		return echo.NewHTTPError(http.StatusBadRequest, response)
	}
	//cek apakah email sudah ada
	cekUser, _ := database.FindUserByEmail(&models.Users{}, u.Email)
	if cekUser != nil {
		response := helpers.ResponseFormater(http.StatusBadRequest, "error", "email alredy used", nil)
		return echo.NewHTTPError(http.StatusBadRequest, response)
	}
	// cek apakah no hp ada
	//cek apakah email sudah ada
	cekUserByPhone, _ := database.FindUserByPhone(&models.Users{}, u.PhoneNumber)
	if cekUserByPhone != nil {
		response := helpers.ResponseFormater(http.StatusBadRequest, "error", "phone_number alredy used", nil)
		return echo.NewHTTPError(http.StatusBadRequest, response)
	}

	// fmt.Println(cekUser)
	hashed, err := bcrypt.GenerateFromPassword([]byte(u.Password), 14)
	if err != nil {
		response := helpers.ResponseFormater(http.StatusBadRequest, "error", err.Error(), helpers.ErrorFormater(err))
		return echo.NewHTTPError(http.StatusBadRequest, response)
	}
	//--- Extract Value without specifying Type

	// fmt.Println(user.(*User).UserEmail)
	formData := models.Users{
		Name:        u.Name,
		Email:       u.Email,
		Password:    string(hashed),
		ProvincesId: city.ProvincesId,
		Provinces:   city.Provinces,
		CityId:      city.ID,
		City:        *city,
		ZipCode:     u.ZipCode,
		Address:     u.Address,
		PhoneNumber: u.PhoneNumber,
	}
	user, err := database.StoreUser(&formData)
	if err != nil {
		response := helpers.ResponseFormater(http.StatusBadRequest, "error", err.Error(), nil)
		return echo.NewHTTPError(http.StatusBadRequest, response)
	}
	response := helpers.ResponseFormater(http.StatusOK, "success", "success register user", user)
	return c.JSON(http.StatusOK, response)
}

func LoginController(c echo.Context) error {
	u := new(models.UserLoginRequest)
	c.Bind(u)
	if err := c.Validate(u); err != nil {
		response := helpers.ResponseFormater(http.StatusBadRequest, "error", "invalid request", helpers.ErrorFormater(err))
		return echo.NewHTTPError(http.StatusBadRequest, response)
	}
	user, err := database.LoginUser(u.Email, u.Password)
	if err != nil {
		response := helpers.ResponseFormater(http.StatusBadRequest, "error", err.Error(), nil)
		return echo.NewHTTPError(http.StatusBadRequest, response)
	}

	response := helpers.ResponseFormater(http.StatusOK, "success", "success login user", user)
	return c.JSON(http.StatusOK, response)
}
