package controllers

import (
	"go-ecommerce/helpers"
	"go-ecommerce/lib/database"
	"go-ecommerce/middlewares"
	"go-ecommerce/models"
	"net/http"

	"github.com/labstack/echo"
)

func CheckoutController(c echo.Context) error {

	u := new(models.CheckoutReqeust)
	c.Bind(u)
	if err := c.Validate(u); err != nil {
		response := helpers.ResponseFormater(http.StatusBadRequest, "error", "invalid request", helpers.ErrorFormater(err))
		return echo.NewHTTPError(http.StatusBadRequest, response)
	}
	userId := middlewares.ExtractTokenUserId(c)
	// s := strings.Split(u.CartId, ",")
	checkout := models.Checkout{
		CartId:            u.CartId,
		ShippingPrice:     u.ShippingPrice,
		TransactionStatus: u.TransactionStatus,
		Resi:              u.Resi,
		PaymentId:         u.PaymentId,
		UserId:            userId,
	}

	cart, e := database.CheckoutChart(&checkout)
	if e != nil {
		response := helpers.ResponseFormater(http.StatusBadRequest, "error", e.Error(), nil)
		return echo.NewHTTPError(http.StatusBadRequest, response)
	}

	// fmt.Println()

	return c.JSON(http.StatusOK, cart)
}
