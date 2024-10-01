package routes

import (
	"github.com/jidcode/go-commerce/internal/handlers"
	"github.com/labstack/echo/v4"
)

func RegisterRoute(e *echo.Echo, authHandler *handlers.AuthHandler) {
	e.POST("/register", authHandler.Register)
	e.POST("login", authHandler.Login)
}
