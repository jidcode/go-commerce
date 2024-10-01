package main

import (
	"log"

	"github.com/go-playground/validator/v10"
	"github.com/jidcode/go-commerce/internal/api/handlers"
	customMiddleware "github.com/jidcode/go-commerce/internal/api/middleware"
	"github.com/jidcode/go-commerce/internal/api/repository"
	"github.com/jidcode/go-commerce/internal/config"
	"github.com/jidcode/go-commerce/internal/db"
	"github.com/jidcode/go-commerce/internal/services/auth"
	customvalidator "github.com/jidcode/go-commerce/internal/validator"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	cfg := config.LoadEnv()

	// Connect to the database
	db.ConnectDB(*cfg)

	userRepo := repository.NewUserRepository(db.DB)
	authService := auth.NewAuthService(userRepo, cfg)
	authHandler := handlers.NewAuthHandler(authService)

	// New router
	e := echo.New()

	// Register the custom validator
	e.Validator = &customvalidator.CustomValidator{Validator: validator.New()}

	// Middleware
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())

	// Health Check route
	e.GET("/health", func(ctx echo.Context) error {
		return ctx.JSON(200, map[string]string{"status": "OK..."})
	})

	// Routes
	e.POST("/register", authHandler.Register)
	e.POST("/login", authHandler.Login)

	// Protected routes
	protectedGroup := e.Group("/api")
	protectedGroup.Use(customMiddleware.JWTMiddleware(authService))

	// Admin routes
	adminGroup := protectedGroup.Group("/admin")
	adminGroup.Use(customMiddleware.RoleMiddleware("admin"))

	// Start server
	log.Fatal(e.Start(":" + cfg.Port))
}
