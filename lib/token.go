package lib

import (
	"os"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

func GenerateToken(data JSON, exp time.Time, subject string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user": data,
		"exp":  exp.Unix(),
		"subject": subject 
	})

	// get jwt secret key value
	jwtsecret := os.Getenv("JWT")
	result, err := token.SignedString(jwtsecret)

	return result, err
}
