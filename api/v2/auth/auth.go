package auth

import (
	"github.com/labstack/echo"
)

func ApplyRoutes(e *echo.Group) {
	auth := e.Group("/auth")
	auth.GET("/code/:code", code)

	auth.POST("/sendmail", sendAuthEmail)
	auth.POST("/register/local", locaRegister)
}
