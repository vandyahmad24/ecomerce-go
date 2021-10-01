package middlewares

import (
	"github.com/labstack/echo"
)

func IsAdmin(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// user := c.Get("user").(*jwt.Token)
		isAdmin := ExtractTokenIsAdmin(c)
		if !isAdmin {
			return echo.ErrUnauthorized
		}
		return next(c)

	}
}

func IsActive(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// user := c.Get("user").(*jwt.Token)
		isAdmin := ExtractTokenIsActive(c)
		if !isAdmin {
			return echo.ErrUnauthorized
		}
		return next(c)

	}
}
