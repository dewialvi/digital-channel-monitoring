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

	// Dependency Injection: Repository -> Service -> Handler

	// Authentication
	userRepo := repository.NewUserRepository(db)
	authService := service.NewAuthService(userRepo, cfg)
	authHandler := handler.NewAuthHandler(authService)

	// Bug Report
	bugReportRepo := repository.NewBugReportRepository(db)
	bugReportService := service.NewBugReportService(bugReportRepo)
	bugHandler := handler.NewBugReportHandler(bugReportService)

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	routes.SetupRoutes(e, authHandler, bugHandler, cfg)

	fmt.Println("Digital Channel Monitoring System")
	fmt.Printf("Environment: %s\n", cfg.AppEnv)
	fmt.Printf("Server berjalan di http://localhost:%s\n", cfg.AppPort)

	e.Logger.Fatal(e.Start(":" + cfg.AppPort))
}
