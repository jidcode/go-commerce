package middleware

import (
	"net/http"
	"strings"

	"github.com/jidcode/go-commerce/internal/models"
	"github.com/jidcode/go-commerce/internal/services/auth"
	"github.com/labstack/echo/v4"
)

func JWTMiddleware(authService *auth.AuthService) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			authHeader := c.Request().Header.Get("Authorization")
			if authHeader == "" {
				return echo.NewHTTPError(http.StatusUnauthorized, "missing auth token")
			}

			tokenParts := strings.Split(authHeader, " ")
			if len(tokenParts) != 2 || strings.ToLower(tokenParts[0]) != "bearer" {
				return echo.NewHTTPError(http.StatusUnauthorized, "invalid auth token format")
			}

			token := tokenParts[1]
			user, err := authService.GetUserFromToken(token)
			if err != nil {
				return echo.NewHTTPError(http.StatusUnauthorized, "invalid auth token")
			}

			c.Set("user", user)
			return next(c)
		}
	}
}

func RoleMiddleware(roles ...string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			user, ok := c.Get("user").(*models.User)
			if !ok {
				return echo.NewHTTPError(http.StatusUnauthorized, "user not found in context")
			}

			for _, role := range roles {
				if user.Role == role {
					return next(c)
				}
			}

			return echo.NewHTTPError(http.StatusForbidden, "insufficient permissions")
		}
	}
}

// Example Usage
// Here’s a hypothetical example of how routes might be defined within these groups:

// // Define a protected route
// protectedGroup.GET("/public", publicHandler) // Accessible by any authenticated user

// // Define an admin route
// adminGroup.GET("/dashboard", adminDashboardHandler) // Accessible only by users with the "admin" role
// In this structure:

// The /api/public route would require a valid JWT but would be accessible to any authenticated user.
// The /api/admin/dashboard route would require a valid JWT and that the user has the "admin" role to access it.
