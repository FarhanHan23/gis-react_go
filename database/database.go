package database

import (
	"FarhanHan23/gis-go/config"
	"FarhanHan23/gis-go/models"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {

	// Load konfigurasi database dari .env
	dbUser := config.GetEnv("DB_USER", "root")
	dbPass := config.GetEnv("DB_PASS", "")
	dbHost := config.GetEnv("DB_HOST", "localhost")
	dbPort := config.GetEnv("DB_PORT", "3306") // Port default Postgres: 5432
	dbName := config.GetEnv("DB_NAME", "db_golang_gis")
	dbSSL := config.GetEnv("DB_SSLMODE", "disable") // Umumnya 'disable' untuk local development
	dbTZ := config.GetEnv("DB_TIMEZONE", "Asia/Jakarta")

	// 2. Format DSN khusus PostgreSQL
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		dbHost, dbUser, dbPass, dbName, dbPort, dbSSL, dbTZ,
	)

	// Koneksi ke database
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	fmt.Println("Database connected successfully!")

	// **Auto Migrate Models**
	err = DB.AutoMigrate(&models.User{}, &models.Role{}, &models.Permission{}, &models.Category{}, &models.Map{}, &models.Setting{}) // Tambahkan model lain jika perlu
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	fmt.Println("Database migrated successfully!")

}
