package main

import (
	"log"

	"github.com/jidcode/go-commerce/internal/api/handlers"
	"github.com/jidcode/go-commerce/internal/api/repository"
	"github.com/jidcode/go-commerce/internal/api/routes"
	"github.com/jidcode/go-commerce/internal/config"
	"github.com/jidcode/go-commerce/internal/db"
	"github.com/jidcode/go-commerce/internal/services/auth"
)

func main() {
	cfg := config.LoadEnv()

	// Connect to the database
	db.ConnectDB(*cfg)

	// Initialize repositories
	userRepo := repository.NewUserRepository(db.DB)
	storeRepo := repository.NewStoreRepository(db.DB)
	categoryRepo := repository.NewCategoryRepository(db.DB)

	// Initialize services
	authService := auth.NewAuthService(userRepo, cfg)

	// Initialize handlers
	authRouter := handlers.NewAuthHandler(authService)
	storeRouter := handlers.NewStoreHandler(storeRepo)
	categoryRouter := handlers.NewCategoryHandler(categoryRepo)

	// Set up routes
	router := routes.Router(authService, authRouter, storeRouter, categoryRouter)

	// Start server
	log.Println("Starting server on :5000...")
	log.Fatal(router.Start(":" + cfg.Port))
}
