package main

import (
	"FarhanHan23/gis-go/config"
	"FarhanHan23/gis-go/database"
	"FarhanHan23/gis-go/seeders"

	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadEnv()

	// inisialisasi database
	database.InitDB()

	seeders.Seed()
	router := gin.Default()

	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Hello Gin",
		})
	})

	router.Run(":3000")
}
