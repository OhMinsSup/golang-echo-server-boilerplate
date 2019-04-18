package lib

import (
	"os"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

// GenerateToken 토큰을 생성하는 함수
func GenerateToken(data interface{}, exp time.Time, subject string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user":    data,
		"exp":     exp.Unix(),
		"subject": subject,
	})

	// get jwt secret key value
	jwtsecret := os.Getenv("JWT")
	result, err := token.SignedString(jwtsecret)

	return result, err
}
