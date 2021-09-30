package controllers

import (
	"go-ecommerce/helpers"
	"go-ecommerce/lib/database"
	"go-ecommerce/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

type CityRequest struct {
	ID          uint
	Name        string `json:"name" form:"name" validate:"required" query:"name"`
	ProvincesId uint   `json:"provinces_id,string" form:"provinces_id" validate:"required"`
}

func GetCityByProvincesIdController(c echo.Context) error {
	prov_id, _ := strconv.Atoi(c.Param("id"))
	city, e := database.GetCityByProvincesId(prov_id)
	if e != nil {
		response := helpers.ResponseFormater(http.StatusBadRequest, "error", e.Error(), nil)
		return echo.NewHTTPError(http.StatusBadRequest, response)
	}
	response := helpers.ResponseFormater(http.StatusOK, "success", "success get city", city)
	return c.JSON(http.StatusOK, response)
}

func GetCityByIdController(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	city, e := database.GetCityById(&models.City{}, id)
	if e != nil {
		response := helpers.ResponseFormater(http.StatusBadRequest, "error", e.Error(), nil)
		return echo.NewHTTPError(http.StatusBadRequest, response)
	}
	response := helpers.ResponseFormater(http.StatusOK, "success", "success get city", city)
	return c.JSON(http.StatusOK, response)
}

func StoreCityController(c echo.Context) error {
	u := new(CityRequest)
	c.Bind(u)
	if err := c.Validate(u); err != nil {
		response := helpers.ResponseFormater(http.StatusBadRequest, "error", "invalid request", helpers.ErrorFormater(err))
		return echo.NewHTTPError(http.StatusBadRequest, response)
	}
	_, err := database.GetProvincesById(&models.Provinces{}, int(u.ProvincesId))
	if err != nil {
		response := helpers.ResponseFormater(http.StatusBadRequest, "error", "provinces not found", nil)
		return echo.NewHTTPError(http.StatusBadRequest, response)
	}
	kota := models.City{
		Name:        u.Name,
		ProvincesId: uint(u.ProvincesId),
	}

	city, e := database.Storecity(&kota)
	if e != nil {
		response := helpers.ResponseFormater(http.StatusBadRequest, "error", err.Error(), nil)
		return echo.NewHTTPError(http.StatusBadRequest, response)
	}
	return c.JSON(http.StatusOK, city)

}

func DeleteCityController(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	// cek apakah id nya ada
	_, e := database.GetCityById(&models.City{}, id)
	if e != nil {
		response := helpers.ResponseFormater(http.StatusBadRequest, "error", e.Error(), nil)
		return echo.NewHTTPError(http.StatusBadRequest, response)
	}
	city, e := database.DeleteCityById(id)
	if e != nil {
		response := helpers.ResponseFormater(http.StatusBadRequest, "error", e.Error(), nil)
		return echo.NewHTTPError(http.StatusBadRequest, response)
	}
	response := helpers.ResponseFormater(http.StatusOK, "success", "success delete City", city)
	return c.JSON(http.StatusOK, response)
}

func PutCityByIDController(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	u := new(CityRequest)
	c.Bind(u)
	if err := c.Validate(u); err != nil {
		response := helpers.ResponseFormater(http.StatusBadRequest, "error", "invalid request", helpers.ErrorFormater(err))
		return echo.NewHTTPError(http.StatusBadRequest, response)
	}

	// cek apakah IDnya ada
	_, errorCity := database.GetCityById(&models.City{}, id)
	if errorCity != nil {
		response := helpers.ResponseFormater(http.StatusBadRequest, "error", errorCity.Error(), nil)
		return echo.NewHTTPError(http.StatusBadRequest, response)
	}

	// cek apakah prov id ada
	_, err := database.GetProvincesById(&models.Provinces{}, int(u.ProvincesId))
	if err != nil {
		response := helpers.ResponseFormater(http.StatusBadRequest, "error", "provinces not found", nil)
		return echo.NewHTTPError(http.StatusBadRequest, response)
	}

	result, _ := database.PutCityById(&models.City{}, id, u.Name, u.ProvincesId)
	response := helpers.ResponseFormater(http.StatusOK, "success", "success update city", result)
	return c.JSON(http.StatusOK, response)
}
