package v2

import (
	"github.com/OhMinsSup/pin-server/api/v2/auth"
	"github.com/labstack/echo"
)

func ApplyRoutes(e *echo.Group) {
	v2 := e.Group("/v2")
	auth.ApplyRoutes(v2)
}
