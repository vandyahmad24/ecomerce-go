package controllers

import (
	"go-ecommerce/helpers"
	"go-ecommerce/lib/database"
	"go-ecommerce/middlewares"
	"go-ecommerce/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

func GetListCartController(c echo.Context) error {
	userId := middlewares.ExtractTokenUserId(c)
	cart, e := database.GetListCart(userId)
	if e != nil {
		response := helpers.ResponseFormater(http.StatusBadRequest, "error", e.Error(), nil)
		return echo.NewHTTPError(http.StatusBadRequest, response)
	}
	response := helpers.ResponseFormater(http.StatusOK, "success", "list cart", cart)
	return c.JSON(http.StatusOK, response)

}

func DeleteCartController(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("cart_id"))
	cart, e := database.DeleteCartById(id)
	if e != nil {
		response := helpers.ResponseFormater(http.StatusBadRequest, "error", e.Error(), nil)
		return echo.NewHTTPError(http.StatusBadRequest, response)
	}
	response := helpers.ResponseFormater(http.StatusOK, "success", "Delete cart", cart)
	return c.JSON(http.StatusOK, response)

}

func AddCartByProdukProdukController(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("product_id"))

	u := new(models.CartsRequest)
	c.Bind(u)
	if err := c.Validate(u); err != nil {
		response := helpers.ResponseFormater(http.StatusBadRequest, "error", "invalid request", helpers.ErrorFormater(err))
		return echo.NewHTTPError(http.StatusBadRequest, response)
	}
	// cek apakah produk tersedia
	_, e := database.GetProductById(id)
	if e != nil {
		response := helpers.ResponseFormater(http.StatusBadRequest, "error", e.Error(), nil)
		return echo.NewHTTPError(http.StatusBadRequest, response)
	}

	idFromToken := middlewares.ExtractTokenUserId(c)
	cart := models.Carts{
		ProductsId: uint(id),
		UsersId:    uint(idFromToken),
		Qty:        u.Qty,
	}

	produk, e := database.AddCartByProdukID(&cart, id)
	if e != nil {
		response := helpers.ResponseFormater(http.StatusBadRequest, "error", e.Error(), nil)
		return echo.NewHTTPError(http.StatusBadRequest, response)
	}
	response := helpers.ResponseFormater(http.StatusOK, "success", "success add cart to produk", produk)
	return c.JSON(http.StatusOK, response)
}
