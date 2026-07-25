package routes

import (
	"github.com/labstack/echo/v4"

	"github.com/dewialvi/digital-channel-monitoring/config"
	"github.com/dewialvi/digital-channel-monitoring/handler"
	"github.com/dewialvi/digital-channel-monitoring/middleware"
	"github.com/dewialvi/digital-channel-monitoring/repository"
)

func SetupRoutes(
	e *echo.Echo,
	authHandler *handler.AuthHandler,
	bugHandler *handler.BugReportHandler,
	apiMonitorHandler *handler.APIMonitorHandler,
	trxHandler *handler.TransactionMonitorHandler,
	activityLogHandler *handler.ActivityLogHandler,
	apiMonitorRepo *repository.APIMonitorRepository,
	cfg *config.Config,
) {

	// =========================
	// API MONITORING MIDDLEWARE
	// =========================

	// Middleware dipasang secara global agar semua request
	// tercatat ke tabel api_monitors.
	e.Use(middleware.APIMonitorMiddleware(apiMonitorRepo))

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

	// =========================
	// AUTH
	// =========================

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
	// MONITORING API ROUTES
	// =========================

	// Get API Monitoring Logs
	protected.GET("/monitoring/api-logs", apiMonitorHandler.GetAll)

	// Get API Monitoring Statistics
	protected.GET("/monitoring/api-stats", apiMonitorHandler.GetStats)

	// =========================
	// TRANSACTION MONITORING
	// =========================

	// Create Transaction Monitoring
	protected.POST("/monitoring/transactions", trxHandler.Create)

	// Get Transaction Monitoring
	protected.GET("/monitoring/transactions", trxHandler.GetAll)

	protected.GET("/activity-logs", activityLogHandler.GetAll)

	// =========================
	// ADMIN ROUTES
	// =========================

	admin := protected.Group("/admin")
	admin.Use(middleware.RoleMiddleware("admin"))

	// Delete Bug Report
	admin.DELETE("/bug-reports/:id", bugHandler.Delete)
}