package api

import (
	v2 "github.com/OhMinsSup/pin-server/api/v2"
	"github.com/labstack/echo"
)

// ApplyRoutes 라우터
func ApplyRoutes(e *echo.Echo) {
	api := e.Group("/api")
	v2.ApplyRoutes(api)
}
