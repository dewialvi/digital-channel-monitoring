package migrations

import (
	"log"

	"gorm.io/gorm"

	"github.com/dewialvi/digital-channel-monitoring/models"
)

func RunMigrations(db *gorm.DB) {
	err := db.AutoMigrate(
		// User harus pertama karena menjadi parent
		// untuk beberapa tabel lainnya.
		&models.User{},

		// Tabel yang memiliki relasi dengan User
		&models.ActivityLog{},
		&models.BugReport{},
		&models.UserFeedback{},
		&models.Notification{},

		// Tabel monitoring yang berdiri sendiri
		&models.APIMonitor{},
		&models.TransactionMonitor{},
	)

	if err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}

	log.Println("Migrations completed successfully!")
}