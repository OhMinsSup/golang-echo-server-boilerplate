package database

import (
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
)

func Inject(db *gorm.DB) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set("db", db)
			return next(c)
		}
	}
}
