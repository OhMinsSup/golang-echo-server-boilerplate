package auth

import (
	"log"
	"net/http"

	"github.com/labstack/echo"
)

type SendEmail struct {
	Email string `json:"email" form:"email" query:"email"`
}

func sendAuthEmail(c echo.Context) error {
	u := new(SendEmail)
	if err := c.Bind(u); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Bad request Data: "+err.Error())
	}
	log.Printf("%s", u)

	return c.JSON(http.StatusOK, u)
}

func locaRegister(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
