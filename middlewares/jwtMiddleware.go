package middlewares

import (
	"go-ecommerce/constants"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

func CreateToken(userId int, isAdmin bool, isActive bool) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["userId"] = userId
	claims["isAdmin"] = isAdmin
	claims["isActive"] = isActive
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(constants.SECRET_JWT))

}

func ExtractTokenUserId(e echo.Context) int {
	user := e.Get("user").(*jwt.Token)
	if user.Valid {
		claims := user.Claims.(jwt.MapClaims)
		userId := claims["userId"].(float64)
		return int(userId)
	}
	return 0
}

func ExtractTokenIsAdmin(e echo.Context) bool {
	user := e.Get("user").(*jwt.Token)
	if user.Valid {
		claims := user.Claims.(jwt.MapClaims)
		isAdmin := claims["isAdmin"].(bool)
		return isAdmin
	}
	return false
}

func ExtractTokenIsActive(e echo.Context) bool {
	user := e.Get("user").(*jwt.Token)
	if user.Valid {
		claims := user.Claims.(jwt.MapClaims)
		isActive := claims["isActive"].(bool)
		return isActive
	}
	return false
}

// func ExtractTokenLevel(e echo.Context) bool {
// 	user := e.Get("user").(*jwt.Token)
// 	if user.Valid {
// 		claims := user.Claims.(jwt.MapClaims)
// 		// isAdmin := claims["isAdmin"].(bool)
// 		return false
// 	}
// 	return 0
// }
