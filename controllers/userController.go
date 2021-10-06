package controllers

import (
	"go-ecommerce/helpers"
	"go-ecommerce/lib/database"
	"go-ecommerce/lib/utils"
	"go-ecommerce/middlewares"
	"go-ecommerce/models"
	"net/http"

	"github.com/labstack/echo"
)

func GetUserByTokenControllers(c echo.Context) error {
	idFromToken := middlewares.ExtractTokenUserId(c)

	user, e := database.FindUserById(idFromToken)
	if e != nil {
		response := helpers.ResponseFormater(http.StatusBadRequest, "error", e.Error(), nil)
		return echo.NewHTTPError(http.StatusBadRequest, response)
	}

	response := helpers.ResponseFormater(http.StatusOK, "success", "success get user", user)
	return c.JSON(http.StatusOK, response)
}

func ResendOTPController(c echo.Context) error {
	idFromToken := middlewares.ExtractTokenUserId(c)
	kodeOTP := utils.EncodeToString(4)
	text := "Anda mengajukan permohonan OTP kode OTP anda " + kodeOTP
	user, e := database.ReCreateOTP(idFromToken, kodeOTP, text)
	if e != nil {
		response := helpers.ResponseFormater(http.StatusBadRequest, "error", e.Error(), nil)
		return echo.NewHTTPError(http.StatusBadRequest, response)
	}

	response := helpers.ResponseFormater(http.StatusOK, "success", "success change OTP", user)
	return c.JSON(http.StatusOK, response)
}

func VerifOTPController(c echo.Context) error {
	idFromToken := middlewares.ExtractTokenUserId(c)
	u := new(models.UserVerifOTPRequest)
	c.Bind(u)
	if err := c.Validate(u); err != nil {
		response := helpers.ResponseFormater(http.StatusBadRequest, "error", "invalid request", helpers.ErrorFormater(err))
		return echo.NewHTTPError(http.StatusBadRequest, response)
	}
	user, e := database.VerifOTP(idFromToken, u.OtpCode)
	if e != nil {
		response := helpers.ResponseFormater(http.StatusBadRequest, "error", e.Error(), nil)
		return echo.NewHTTPError(http.StatusBadRequest, response)
	}

	response := helpers.ResponseFormater(http.StatusOK, "success", "success change OTP", user)
	return c.JSON(http.StatusOK, response)

}
