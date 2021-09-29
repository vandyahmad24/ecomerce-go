package controllers

import (
	"go-ecommerce/helpers"
	"go-ecommerce/lib/database"
	"go-ecommerce/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

type ProvincesRequest struct {
	Name string `json:"name" form:"name" validate:"required" query:"name"`
}

func StoreProvincesController(c echo.Context) error {
	u := new(ProvincesRequest)
	c.Bind(u)

	provinces := models.Provinces{
		Name: u.Name,
	}
	if err := c.Validate(u); err != nil {
		response := helpers.ResponseFormater(http.StatusBadRequest, "error", "invalid request", helpers.ErrorFormater(err))
		return echo.NewHTTPError(http.StatusBadRequest, response)
	}

	provnices, err := database.StoreProvinces(&provinces)
	if err != nil {
		response := helpers.ResponseFormater(http.StatusBadRequest, "error", err.Error(), nil)
		return echo.NewHTTPError(http.StatusBadRequest, response)
	}
	response := helpers.ResponseFormater(http.StatusOK, "success", "success store provinces", provnices)
	return c.JSON(http.StatusOK, response)
}

func GetAllProvincesController(c echo.Context) error {
	provnices, e := database.GetProvinces()
	if e != nil {
		response := helpers.ResponseFormater(http.StatusBadRequest, "error", e.Error(), nil)
		return echo.NewHTTPError(http.StatusBadRequest, response)
	}
	response := helpers.ResponseFormater(http.StatusOK, "success", "success get provinces", provnices)
	return c.JSON(http.StatusOK, response)
}

func GetProvincesByIDController(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	provnices, e := database.GetProvincesById(&models.Provinces{}, id)
	if e != nil {
		response := helpers.ResponseFormater(http.StatusBadRequest, "error", e.Error(), nil)
		return echo.NewHTTPError(http.StatusBadRequest, response)
	}

	response := helpers.ResponseFormater(http.StatusOK, "success", "success get provinces", provnices)
	return c.JSON(http.StatusOK, response)
}

func DeleteProvincesByIDController(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	provnices, e := database.DeleteProvincesById(id)
	if e != nil {
		response := helpers.ResponseFormater(http.StatusBadRequest, "error", e.Error(), nil)
		return echo.NewHTTPError(http.StatusBadRequest, response)
	}
	response := helpers.ResponseFormater(http.StatusOK, "success", "success delete provinces", provnices)
	return c.JSON(http.StatusOK, response)

}

func PutProvincesByIDController(c echo.Context) error {

	id, _ := strconv.Atoi(c.Param("id"))
	u := new(ProvincesRequest)
	c.Bind(u)
	if err := c.Validate(u); err != nil {
		response := helpers.ResponseFormater(http.StatusBadRequest, "error", "invalid request", helpers.ErrorFormater(err))
		return echo.NewHTTPError(http.StatusBadRequest, response)
	}

	formData := models.Provinces{
		Name: u.Name,
	}
	result, _ := database.PutProvincesByID(&models.Provinces{}, id, formData.Name)
	response := helpers.ResponseFormater(http.StatusOK, "success", "success update provinces", result)
	return c.JSON(http.StatusOK, response)

}
