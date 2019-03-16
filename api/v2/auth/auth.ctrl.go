package auth

import (
	"net/http"

	"github.com/labstack/echo"
)

func locaRegister(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
