package auth

import (
	"github.com/labstack/echo"
)

func ApplyRoutes(e *echo.Group) {
	auth := e.Group("/auth")
	
	auth.POST("/register/local", locaRegister)
}
