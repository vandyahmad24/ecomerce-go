package controllers

import (
	"go-ecommerce/helpers"
	"go-ecommerce/lib/database"
	"go-ecommerce/middlewares"
	"go-ecommerce/models"
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo"
)

func GetAllCategoryActive(c echo.Context) error {
	category, e := database.GetAllCategoryActive()
	if e != nil {
		response := helpers.ResponseFormater(http.StatusBadRequest, "error", e.Error(), nil)
		return echo.NewHTTPError(http.StatusBadRequest, response)
	}
	response := helpers.ResponseFormater(http.StatusOK, "success", "success get category", category)
	return c.JSON(http.StatusOK, response)
}

func GetAllCategoryInActive(c echo.Context) error {
	category, e := database.GetAllCategoryInActive()
	if e != nil {
		response := helpers.ResponseFormater(http.StatusBadRequest, "error", e.Error(), nil)
		return echo.NewHTTPError(http.StatusBadRequest, response)
	}
	response := helpers.ResponseFormater(http.StatusOK, "success", "success get category", category)
	return c.JSON(http.StatusOK, response)
}

func InsertCategory(c echo.Context) error {
	u := new(models.CategoryRequest)
	c.Bind(u)
	if err := c.Validate(u); err != nil {
		response := helpers.ResponseFormater(http.StatusBadRequest, "error", "invalid request", helpers.ErrorFormater(err))
		return echo.NewHTTPError(http.StatusBadRequest, response)
	}
	slug := strings.ReplaceAll(u.Name, " ", "_")

	// cek apakah yg post admin atau user jika admin maka status categori true
	cekLevel := middlewares.ExtractTokenIsAdmin(c)
	category := models.Categories{
		Name:             u.Name,
		Slug:             slug,
		StatusCategories: cekLevel,
	}

	result, e := database.StoreCategory(&category)
	if e != nil {
		response := helpers.ResponseFormater(http.StatusBadRequest, "error", e.Error(), nil)
		return echo.NewHTTPError(http.StatusBadRequest, response)
	}
	response := helpers.ResponseFormater(http.StatusOK, "success", "success create category", result)
	return c.JSON(http.StatusOK, response)
}

func GetCategoryByID(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	category, e := database.GetCategoryById(&models.Categories{}, id)
	if e != nil {
		response := helpers.ResponseFormater(http.StatusBadRequest, "error", e.Error(), nil)
		return echo.NewHTTPError(http.StatusBadRequest, response)
	}

	response := helpers.ResponseFormater(http.StatusOK, "success", "success get category", category)
	return c.JSON(http.StatusOK, response)
}

func UpdateCategoryByID(c echo.Context) error {
	return c.JSON(http.StatusOK, "Update")
}
