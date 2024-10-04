package routes

import (
	"github.com/go-playground/validator/v10"
	"github.com/jidcode/go-commerce/internal/api/handlers"
	customMiddleware "github.com/jidcode/go-commerce/internal/api/middleware"
	"github.com/jidcode/go-commerce/internal/services/auth"
	customvalidator "github.com/jidcode/go-commerce/internal/validator"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Router(authService *auth.AuthService, authRouter *handlers.AuthHandler, storeRouter *handlers.StoreHandler, categoryRouter *handlers.CategoryHandler) *echo.Echo {
	// New router
	e := echo.New()

	// Register the custom validator
	e.Validator = &customvalidator.CustomValidator{Validator: validator.New()}

	// Middleware
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())

	// Health Check
	e.GET("/health", func(ctx echo.Context) error {
		return ctx.JSON(200, map[string]string{"status": "OK"})
	})

	// Authentication Routes
	e.POST("/register", authRouter.Register)
	e.POST("/login", authRouter.Login)

	// Protected routes
	protectedGroup := e.Group("/api")
	protectedGroup.Use(customMiddleware.JWTMiddleware(authService))

	// Store routes
	protectedGroup.GET("/stores", storeRouter.GetStores)
	protectedGroup.POST("/stores", storeRouter.CreateStore)
	protectedGroup.GET("/stores/:id", storeRouter.GetStoreByID)
	protectedGroup.PUT("/stores/:id", storeRouter.UpdateStore)
	protectedGroup.DELETE("/stores/:id", storeRouter.DeleteStore)

	//category routes
	protectedGroup.GET("/categories", categoryRouter.GetCategories)
	protectedGroup.POST("/categories", categoryRouter.CreateCategory)
	protectedGroup.GET("/categories/:id", categoryRouter.GetCategoryByID)
	protectedGroup.PUT("/categories/:id", categoryRouter.UpdateCategory)
	protectedGroup.DELETE("/categories/:id", categoryRouter.DeleteCategory)

	// Admin routes
	adminGroup := protectedGroup.Group("/admin")
	adminGroup.Use(customMiddleware.RoleMiddleware("admin"))

	return e
}
