package main

import (
	"os"

	"github.com/OhMinsSup/pin-server/api"
	"github.com/OhMinsSup/pin-server/database"
	"github.com/OhMinsSup/pin-server/lib"

	"github.com/joho/godotenv"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	// load .env environment variables
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	e := echo.New()
	port := os.Getenv("PORT")
	db, _ := database.Initialize()

	//validate
	e.Validator = lib.NewValidator()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(database.Inject(db))

	//CORS
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))
	api.ApplyRoutes(e)

	e.Logger.Fatal(e.Start(":" + port))
}
