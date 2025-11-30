package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
)

type HealthController struct {
	beego.Controller
}

// Get handles GET requests to /health
func (c *HealthController) Get() {
	SendSuccess(&c.Controller, map[string]interface{}{
		"status":  "healthy",
		"service": "foliogram-backend",
		"message": "Service is running",
	})
}
