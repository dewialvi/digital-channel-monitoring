package main

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/dewialvi/digital-channel-monitoring/config"
	"github.com/dewialvi/digital-channel-monitoring/handler"
	"github.com/dewialvi/digital-channel-monitoring/migrations"
	"github.com/dewialvi/digital-channel-monitoring/repository"
	"github.com/dewialvi/digital-channel-monitoring/routes"
	"github.com/dewialvi/digital-channel-monitoring/service"
)

func main() {
	cfg := config.LoadConfig()
	db := config.ConnectDatabase(cfg)

	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	migrations.RunMigrations(db)

	// Dependency Injection:
	// Repository -> Service -> Handler

	// =========================
	// AUTHENTICATION
	// =========================

	userRepo := repository.NewUserRepository(db)

	authService := service.NewAuthService(
		userRepo,
		cfg,
	)

	authHandler := handler.NewAuthHandler(
		authService,
	)

	// Activity Log
activityLogRepo := repository.NewActivityLogRepository(db)
activityLogService := service.NewActivityLogService(activityLogRepo)
activityLogHandler := handler.NewActivityLogHandler(activityLogService)

// Notification
notificationService := service.NewNotificationService()

// Bug Report
bugReportRepo := repository.NewBugReportRepository(db)
bugReportService := service.NewBugReportService(
	bugReportRepo,
	activityLogService,
	notificationService,
)
bugHandler := handler.NewBugReportHandler(bugReportService)

	// =========================
	// API MONITORING
	// =========================

	apiMonitorRepo := repository.NewAPIMonitorRepository(db)

	apiMonitorHandler := handler.NewAPIMonitorHandler(
		apiMonitorRepo,
	)

	// =========================
	// TRANSACTION MONITORING
	// =========================

	transactionMonitorRepo := repository.NewTransactionMonitorRepository(db)

	transactionMonitorHandler := handler.NewTransactionMonitorHandler(
		transactionMonitorRepo,
	)

	// =========================
	// ECHO SERVER
	// =========================

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// =========================
	// SETUP ROUTES
	// =========================

	routes.SetupRoutes(
		e,
		authHandler,
		bugHandler,
		apiMonitorHandler,
		transactionMonitorHandler,
		activityLogHandler,
		apiMonitorRepo,
		cfg,
	)

	// =========================
	// FRONTEND
	// =========================

	fmt.Println("Digital Channel Monitoring System")
	fmt.Printf("Environment: %s\n", cfg.AppEnv)
	fmt.Printf(
		"Server berjalan di http://localhost:%s\n",
		cfg.AppPort,
	)

	e.Static("/static", "static")

	e.File("/login", "templates/login.html")

	e.File("/dashboard", "templates/dashboard.html")

	e.File("/bug-reports", "templates/bug-reports.html")

	e.Logger.Fatal(
		e.Start(":" + cfg.AppPort),
	)
}