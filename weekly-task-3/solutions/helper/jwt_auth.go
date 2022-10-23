package helper

import (
	"github.com/arvinpaundra/golang-blog/config"
	jwt "github.com/dgrijalva/jwt-go"
	"time"
)

func CreateToken(userId string) (string, error) {
	claims := jwt.MapClaims{}
	claims["user_id"] = userId
	claims["iat"] = time.Now().Unix()
	claims["exp"] = time.Now().Add(1 * time.Hour).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(config.Cfg.JWTSecret))
}
