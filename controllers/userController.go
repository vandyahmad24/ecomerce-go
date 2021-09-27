package controllers

import (
	"github.com/labstack/echo"
	"go-ecommerce/lib/database"
	"net/http"
)

func GetUserControllers(c echo.Context) error{
	users, e := database.GetUsers()
	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":"success",
		"users":users,
	})
}
