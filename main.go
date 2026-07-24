package main

import (
	"fmt"

	"github.com/dewialvi/digital-channel-monitoring/config"
	"github.com/dewialvi/digital-channel-monitoring/migrations"
)

func main() {
	// Load configuration dari .env
	cfg := config.LoadConfig()

	// Connect ke PostgreSQL
	db := config.ConnectDatabase(cfg)

	// Get database instance untuk menutup koneksi
	// saat aplikasi berhenti
	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}
	defer sqlDB.Close()

	// Jalankan database migration
	migrations.RunMigrations(db)

	// Informasi aplikasi
	fmt.Println("Digital Channel Monitoring System")
	fmt.Printf("Environment: %s\n", cfg.AppEnv)
	fmt.Printf("Server akan berjalan di port: %s\n", cfg.AppPort)
}