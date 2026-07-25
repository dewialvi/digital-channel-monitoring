package routes

import (
	"github.com/labstack/echo/v4"

	"github.com/dewialvi/digital-channel-monitoring/config"
	"github.com/dewialvi/digital-channel-monitoring/handler"
	"github.com/dewialvi/digital-channel-monitoring/middleware"
)

func SetupRoutes(
	e *echo.Echo,
	authHandler *handler.AuthHandler,
	bugHandler *handler.BugReportHandler,
	cfg *config.Config,
) {
	api := e.Group("/api/v1")

	// =========================
	// PUBLIC ROUTES
	// =========================

	api.POST("/register", authHandler.Register)
	api.POST("/login", authHandler.Login)

	// =========================
	// PROTECTED ROUTES
	// =========================

	protected := api.Group("")
	protected.Use(middleware.JWTMiddleware(cfg.JWTSecret))

	// Auth
	protected.POST("/logout", authHandler.Logout)

	// =========================
	// BUG REPORT ROUTES
	// =========================

	// Create Bug Report
	protected.POST("/bug-reports", bugHandler.Create)

	// Get All Bug Reports
	protected.GET("/bug-reports", bugHandler.GetAll)

	// Get Bug Report By ID
	protected.GET("/bug-reports/:id", bugHandler.GetByID)

	// Update Bug Report Status
	protected.PATCH("/bug-reports/:id/status", bugHandler.UpdateStatus)

	// =========================
	// ADMIN ROUTES
	// =========================

	admin := protected.Group("/admin")
	admin.Use(middleware.RoleMiddleware("admin"))

	// Delete Bug Report
	admin.DELETE("/bug-reports/:id", bugHandler.Delete)
}
