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

func Router(
	authService *auth.AuthService,
	authRouter *handlers.AuthHandler,
	storeRouter *handlers.StoreHandler,
	categoryRouter *handlers.CategoryHandler,
	productRouter *handlers.ProductHandler) *echo.Echo {

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

	// Register routes
	RegisterAuthRoutes(e, authRouter)
	RegisterStoreRoutes(e, authService, storeRouter)
	RegisterCategoryRoutes(e, authService, categoryRouter)
	RegisterProductRoutes(e, authService, productRouter)

	return e
}

// RegisterAuthRoutes handles authentication-related routes
func RegisterAuthRoutes(e *echo.Echo, authRouter *handlers.AuthHandler) {
	e.POST("/register", authRouter.Register)
	e.POST("/login", authRouter.Login)
}

// RegisterStoreRoutes handles store-related routes
func RegisterStoreRoutes(e *echo.Echo, authService *auth.AuthService, storeRouter *handlers.StoreHandler) {
	api := e.Group("/api")
	api.Use(customMiddleware.JWTMiddleware(authService))

	api.GET("/stores", storeRouter.GetStores)
	api.POST("/stores", storeRouter.CreateStore)
	api.GET("/stores/:id", storeRouter.GetStoreByID)
	api.PUT("/stores/:id", storeRouter.UpdateStore)
	api.DELETE("/stores/:id", storeRouter.DeleteStore)
}

// RegisterCategoryRoutes handles category-related routes
func RegisterCategoryRoutes(e *echo.Echo, authService *auth.AuthService, categoryRouter *handlers.CategoryHandler) {
	api := e.Group("/api")
	api.Use(customMiddleware.JWTMiddleware(authService))

	api.GET("/categories", categoryRouter.GetCategories)
	api.POST("/categories", categoryRouter.CreateCategory)
	api.GET("/categories/:id", categoryRouter.GetCategoryByID)
	api.PUT("/categories/:id", categoryRouter.UpdateCategory)
	api.DELETE("/categories/:id", categoryRouter.DeleteCategory)
}

// RegisterCategoryRoutes handles category-related routes
func RegisterProductRoutes(e *echo.Echo, authService *auth.AuthService, productRouter *handlers.ProductHandler) {
	api := e.Group("/api")
	api.Use(customMiddleware.JWTMiddleware(authService))

	api.GET("/products", productRouter.GetProducts)
	api.POST("/products", productRouter.CreateProduct)
	api.GET("/products/:id", productRouter.GetProductByID)
	api.PUT("/products/:id", productRouter.UpdateProduct)
	api.DELETE("/products/:id", productRouter.DeleteProduct)
}

// func RegisterAdminRoutes(e *echo.Echo, categoryHandler *handlers.CategoryHandler) {
// 	// Protect category management routes with admin-only middleware
// 	adminGroup := e.Group("/admin", middleware.RoleMiddleware("admin"))

// 	adminGroup.POST("/categories", categoryHandler.CreateCategory)
// 	adminGroup.PUT("/categories/:id", categoryHandler.UpdateCategory)
// 	adminGroup.DELETE("/categories/:id", categoryHandler.DeleteCategory)

// 	// Public access to view categories
// 	e.GET("/categories", categoryHandler.GetCategories)
// 	e.GET("/categories/:id", categoryHandler.GetCategoryByID)
// }
