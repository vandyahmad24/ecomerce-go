package controllers

import (
	"github.com/labstack/echo"
	"go-ecommerce/helpers"
	"go-ecommerce/lib/database"
	"go-ecommerce/models"
	"net/http"
	"strconv"
)

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

func StoreProvincesController(c echo.Context) error {
	provinces := models.Provinces{
		Name: c.FormValue("name"),
	}
	if err := c.Validate(provinces); err != nil {
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

func PutProvincesByIDController(c echo.Context) error {

	id, _ := strconv.Atoi(c.Param("id"))
	formData := models.Provinces{
		Name: c.FormValue("name"),
	}

	if err := c.Validate(formData); err != nil {
		response := helpers.ResponseFormater(http.StatusBadRequest, "error", "invalid request", helpers.ErrorFormater(err))
		return echo.NewHTTPError(http.StatusBadRequest, response)
	}

	result, _ := database.PutProvincesByID(&models.Provinces{}, id, formData.Name)
	response := helpers.ResponseFormater(http.StatusOK, "success", "success get provinces", result)
	return c.JSON(http.StatusOK, response)

}
