package main

import (
	"log"

	"Foliogram-BE/models"
	_ "Foliogram-BE/routers"
	beego "github.com/beego/beego/v2/server/web"
)

func main() {
	// Initialize database connection
	if err := models.InitDB(); err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	defer models.CloseDB()

	log.Println("Starting Foliogram Backend Server...")
	beego.Run()
}

