package main

import (
	"log"

	"github.com/jidcode/go-commerce/internal/config"
	"github.com/jidcode/go-commerce/internal/db"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	cfg := config.LoadEnv()

	// Connect to the database
	db.ConnectDB(*cfg)

	//New router
	e := echo.New()

	// Middleware
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())

	// Health Check route
	e.GET("/health", func(ctx echo.Context) error {
		return ctx.JSON(200, map[string]string{"status": "OK!"})
	})

	//start server
	log.Fatal(e.Start(":" + cfg.Port))
}
