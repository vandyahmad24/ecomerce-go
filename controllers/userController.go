package controllers

import (
	"github.com/labstack/echo"
	"go-ecommerce/lib"
	"go-ecommerce/lib/database"
	"go-ecommerce/models"
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

func AddUserControllers(c echo.Context) error{
	u := new(models.Users)
	if err := c.Bind(u); err != nil {
		response := lib.Message{
			"message":"error",
			"errors":err.Error(),
		}
		return c.JSON(http.StatusBadRequest, response)
	}
	if err := c.Validate(u); err != nil {
		response := lib.Message{
			"message":"error",
			"errors":err.Error(),
		}
		return c.JSON(http.StatusBadRequest, response)
	}

	return c.JSON(http.StatusOK, u)
}
