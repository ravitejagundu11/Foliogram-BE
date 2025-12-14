package main

import (
	"log"
	"os"

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

	// Create uploads directory if it doesn't exist
	if err := os.MkdirAll("./uploads/images", 0755); err != nil {
		log.Printf("Warning: Failed to create uploads directory: %v", err)
	}

	// Enable static file serving for uploads
	beego.SetStaticPath("/uploads", "uploads")

	log.Println("Starting Foliogram Backend Server...")
	beego.Run()
}
