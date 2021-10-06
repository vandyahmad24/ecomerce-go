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

func CreateProdukController(c echo.Context) error {
	u := new(models.ProductsRequest)
	c.Bind(u)
	if err := c.Validate(u); err != nil {
		response := helpers.ResponseFormater(http.StatusBadRequest, "error", "invalid request", helpers.ErrorFormater(err))
		return echo.NewHTTPError(http.StatusBadRequest, response)
	}
	slug := strings.ReplaceAll(u.Name, " ", "_")
	seller := middlewares.ExtractTokenUserId(c)
	// cek apakah category ID ada
	_, e := database.GetCategoryById(&models.Categories{}, int(u.CategoriesId))
	if e != nil {
		response := helpers.ResponseFormater(http.StatusBadRequest, "error", e.Error(), nil)
		return echo.NewHTTPError(http.StatusBadRequest, response)
	}
	produk := models.Products{
		SellerId:     uint(seller),
		Name:         u.Name,
		Price:        u.Price,
		Description:  u.Description,
		Slug:         slug,
		CategoriesId: u.CategoriesId,
	}

	result, e := database.StoreProduct(&produk)
	if e != nil {
		response := helpers.ResponseFormater(http.StatusBadRequest, "error", e.Error(), nil)
		return echo.NewHTTPError(http.StatusBadRequest, response)
	}
	response := helpers.ResponseFormater(http.StatusOK, "success", "success create produk", result)
	return c.JSON(http.StatusOK, response)
}

func GetProdukByCategoryController(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	produk, e := database.GetProductByCategory(id)
	if e != nil {
		response := helpers.ResponseFormater(http.StatusBadRequest, "error", e.Error(), nil)
		return echo.NewHTTPError(http.StatusBadRequest, response)
	}
	response := helpers.ResponseFormater(http.StatusOK, "success", "success get produk", produk)
	return c.JSON(http.StatusOK, response)
}

func GetProdukByIdController(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	produk, e := database.GetProductById(id)
	if e != nil {
		response := helpers.ResponseFormater(http.StatusBadRequest, "error", e.Error(), nil)
		return echo.NewHTTPError(http.StatusBadRequest, response)
	}
	response := helpers.ResponseFormater(http.StatusOK, "success", "success get produk", produk)
	return c.JSON(http.StatusOK, response)
}
