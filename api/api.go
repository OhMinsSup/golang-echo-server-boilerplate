package api

import (
	"github.com/OhMinsSup/pin-server/api/v2"
	"github.com/labstack/echo"
)

func ApplyRoutes(e *echo.Echo) {
	api := e.Group("/api")
	v2.ApplyRoutes(api)
}
