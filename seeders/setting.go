package seeders

import (
	"FarhanHan23/gis-go/models"
	"log"

	"gorm.io/gorm"
)

func SeedSetting(db *gorm.DB) {
	// default values
	defaults := models.Setting{
		Title:           "GIS Rejowinangun",
		Description:     "Eksplorasi Desa Rejowinangun secara interaktif melalui peta GIS.",
		MapCenterLat:    "-7.592589928951457",
		MapCenterLng:    "112.26113954274147",
		MapZoom:         16,
		VillageBoundary: "[]",
	}

	var out models.Setting
	if err := db.
		Where(&models.Setting{Id: 1}).
		Attrs(defaults).
		FirstOrCreate(&out).Error; err != nil {
		log.Fatalf("failed seeding setting: %v", err)
	}
}
