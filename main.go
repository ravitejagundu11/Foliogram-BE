package main

import (
	"log"

	"Foliogram-BE/models"
	_ "Foliogram-BE/routers"
	"Foliogram-BE/utils"

	beego "github.com/beego/beego/v2/server/web"
)

func main() {
	// Initialize database connection
	if err := models.InitDB(); err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	defer models.CloseDB()

	// Initialize Redis connection
	if err := utils.InitRedis(); err != nil {
		log.Printf("Warning: Failed to initialize Redis: %v", err)
		log.Println("Token blacklisting will not work without Redis")
	}
	defer utils.CloseRedis()

	log.Println("Starting Foliogram Backend Server...")
	beego.Run()
}
