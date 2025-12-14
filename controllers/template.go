package controllers

import (
	"Foliogram-BE/models"

	beego "github.com/beego/beego/v2/server/web"
)

// TemplateController handles template operations
type TemplateController struct {
	beego.Controller
}

// GetAll handles GET /api/templates
func (c *TemplateController) GetAll() {
	var templates []models.Template

	// Build query
	query := models.DB.Where("is_active = ?", true)

	// Filter by category if provided
	category := c.GetString("category")
	if category != "" {
		query = query.Where("category = ?", category)
	}

	// Filter by premium status if provided
	if premium := c.GetString("premium"); premium != "" {
		if premium == "true" {
			query = query.Where("is_premium = ?", true)
		} else if premium == "false" {
			query = query.Where("is_premium = ?", false)
		}
	}

	// Execute query
	if err := query.Order("category ASC, created_at DESC").Find(&templates).Error; err != nil {
		SendInternalServerError(&c.Controller, "Failed to fetch templates")
		return
	}

	// Return success response
	SendSuccess(&c.Controller, templates)
}

// GetByID handles GET /api/templates/:id
func (c *TemplateController) GetByID() {
	id := c.Ctx.Input.Param(":id")

	var template models.Template
	if err := models.DB.Where("id = ? AND is_active = ?", id, true).First(&template).Error; err != nil {
		SendNotFound(&c.Controller, "Template not found")
		return
	}

	// Return success response
	SendSuccess(&c.Controller, template)
}

// GetCategories handles GET /api/templates/categories
func (c *TemplateController) GetCategories() {
	var categories []string

	if err := models.DB.Model(&models.Template{}).
		Where("is_active = ?", true).
		Distinct("category").
		Pluck("category", &categories).Error; err != nil {
		SendInternalServerError(&c.Controller, "Failed to fetch categories")
		return
	}

	// Return success response
	SendSuccess(&c.Controller, map[string]interface{}{
		"categories": categories,
	})
}
