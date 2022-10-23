package middleware

import (
	"fmt"
	"github.com/arvinpaundra/clean-architecture/config"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"time"
)

func CreateToken(userId uint) (string, error) {
	claims := jwt.MapClaims{}
	claims["user_id"] = userId
	claims["exp"] = time.Now().Add(1 * time.Hour).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(config.Cfg.JWTSecret))
}

func ExtractToken(c echo.Context) (uint, error) {
	tokenString := c.Request().Header.Get("Authorization")

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(config.Cfg.JWTSecret), nil
	})

	if err != nil {
		return 0, err
	}

	if token.Valid {
		claims := token.Claims.(jwt.MapClaims)
		userId := claims["user_id"].(uint)
		return userId, nil
	}

	return 0, err
}
