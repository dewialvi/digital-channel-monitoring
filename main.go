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
	userRepo := repository.NewUserRepository(db)

	authService := service.NewAuthService(
		userRepo,
		cfg,
	)

	authHandler := handler.NewAuthHandler(
		authService,
	)

	// Membuat Echo server
	e := echo.New()

	// Middleware bawaan Echo
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Setup semua routes
	routes.SetupRoutes(
		e,
		authHandler,
		cfg,
	)

	fmt.Println("Digital Channel Monitoring System")
	fmt.Printf("Environment: %s\n", cfg.AppEnv)
	fmt.Printf("Server berjalan di http://localhost:%s\n", cfg.AppPort)

	// Menjalankan server
	e.Logger.Fatal(
		e.Start(":" + cfg.AppPort),
	)
}