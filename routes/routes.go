package routes

import (
	"github.com/labstack/echo/v4"

	"github.com/dewialvi/digital-channel-monitoring/config"
	"github.com/dewialvi/digital-channel-monitoring/handler"
	"github.com/dewialvi/digital-channel-monitoring/middleware"
)

func SetupRoutes(e *echo.Echo, authHandler *handler.AuthHandler, cfg *config.Config) {
	api := e.Group("/api/v1")

	// Public routes
	// Tidak membutuhkan login
	api.POST("/register", authHandler.Register)
	api.POST("/login", authHandler.Login)

	// Protected routes
	// Membutuhkan JWT
	protected := api.Group("")
	protected.Use(middleware.JWTMiddleware(cfg.JWTSecret))

	protected.POST("/logout", authHandler.Logout)

	// Admin routes
	// Membutuhkan JWT + role admin
	admin := protected.Group("/admin")
	admin.Use(middleware.RoleMiddleware("admin"))
}