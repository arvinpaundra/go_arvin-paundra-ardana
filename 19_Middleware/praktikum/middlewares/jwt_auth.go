package middlewares

import (
	"fmt"
	"time"

	"github.com/arvinpaundra/rest-orm/constants"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

func CreateToken(userId uint) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["user_id"] = userId
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(constants.SECRET_JWT))
}

func ExtractToken(e echo.Context) uint {
	tokenString := e.Request().Header.Get("Authorization")

	// parse token to jwt
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return constants.SECRET_JWT, nil
	})

	if err != nil {
		return 0
	}

	if token.Valid {
		claims := token.Claims.(jwt.MapClaims)
		userId := claims["user_id"].(uint)
		return userId
	}

	return 0
}
