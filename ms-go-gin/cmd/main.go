package main

import (
	"ms-go-gin/config"
	"ms-go-gin/internal/router"
	"ms-go-gin/pkg/db"
)

func main() {
	//	r := gin.Default()
	cfg := config.New()

	database := db.ConnectDB(cfg).Debug()

	r := router.New(cfg, database)

	// Start the Gin server
	r.Run(":" + cfg.Port)
}
