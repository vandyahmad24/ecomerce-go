package controllers

import (
	"go-ecommerce/helpers"

	"github.com/labstack/echo"

	//"go-ecommerce/lib"
	"go-ecommerce/lib/database"
	"go-ecommerce/models"
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
	u := new(models.Users)
	if err := c.Bind(u); err != nil {
		response := helpers.ResponseFormater(http.StatusBadRequest, "error", "invalid request", helpers.ErrorFormater(err))
		return echo.NewHTTPError(http.StatusBadRequest, response)
		//return c.JSON(http.StatusBadRequest, response)
	}
	if err := c.Validate(u); err != nil {
		response := helpers.ResponseFormater(http.StatusBadRequest, "error", "invalid request", helpers.ErrorFormater(err))
		return echo.NewHTTPError(http.StatusBadRequest, response)
	}

	return c.JSON(http.StatusOK, u)
}
