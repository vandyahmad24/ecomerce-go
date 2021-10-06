package controllers

import (
	"net/http"

	"github.com/labstack/echo"
)

func CobaSegalaController(c echo.Context) error {
	// utils.Test()

	return c.JSON(http.StatusOK, "halo")

}
