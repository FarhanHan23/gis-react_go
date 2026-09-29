package seeders

import (
	"FarhanHan23/gis-go/database"
	"log"
)

func Seed() {
	db := database.DB
	log.Println("Running database seeders...")

	// Jalankan seeder secara berurutan
	SeedPermissions(db)
	SeedRoles(db)
	SeedUsers(db)
	SeedSetting(db)

	log.Println("Database seeding completed!")
}
