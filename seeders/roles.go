package seeders

import (
	"FarhanHan23/gis-go/models"

	"gorm.io/gorm"
)

func SeedRoles(db *gorm.DB) {

	// Data role dasar dengan nama
	roles := []models.Role{
		{Name: "admin"},
		{Name: "user"},
	}
	// Ambil semua permission dari database
	var allPermissions []models.Permission
	db.Find(&allPermissions)

	// Loop dan assign permissions sesuai role
	for i := range roles {
		role := &roles[i] // 2. Fix bug pointer loop
		db.FirstOrCreate(role, models.Role{Name: role.Name})
		switch role.Name {
		case "admin":
			db.Model(role).Association("Permissions").Replace(allPermissions)
		case "user":
			var userPerms []models.Permission
			db.Where("name IN?", []string{
				"categories-index", "categories-show",
				"maps-index", "maps-show",
			}).Find(&userPerms)
			db.Model(role).Association("Permissions").Replace(userPerms)
		}
	}
}
