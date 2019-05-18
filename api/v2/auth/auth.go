package auth

import (
	"github.com/labstack/echo"
)

// ApplyRoutes 라우터
func ApplyRoutes(e *echo.Group) {
	auth := e.Group("/auth")

	auth.POST("/register/local", locaRegister)
	auth.POST("/login/local", login)
	auth.POST("/logout", logout)
}
