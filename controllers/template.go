package controllers

import (
	"Foliogram-BE/models"

	beego "github.com/beego/beego/v2/server/web"
)

// TemplateController handles template operations
type TemplateController struct {
	beego.Controller
}

// GetAll handles GET /api/foliogram/templates
func (c *TemplateController) GetAll() {
	// Get query parameters
	category := c.GetString("category")
	premiumOnly := c.GetString("premium") == "true"

	query := models.DB.Where("is_active = ?", true)

	// Filter by category if provided
	if category != "" {
		query = query.Where("category = ?", category)
	}

	// Filter by premium if requested
	if premiumOnly {
		query = query.Where("is_premium = ?", true)
	}

	var templates []models.Template
	if err := query.Order("created_at DESC").Find(&templates).Error; err != nil {
		SendInternalError(&c.Controller, map[string]interface{}{
			"message": "Failed to fetch templates",
			"error":   err.Error(),
		})
		return
	}

	SendSuccess(&c.Controller, templates)
}

// GetByID handles GET /api/foliogram/templates/:id
func (c *TemplateController) GetByID() {
	templateID := c.Ctx.Input.Param(":id")

	var template models.Template
	if err := models.DB.First(&template, "id = ? AND is_active = ?", templateID, true).Error; err != nil {
		SendNotFound(&c.Controller, map[string]interface{}{
			"message": "Template not found",
		})
		return
	}

	SendSuccess(&c.Controller, template)
}

// GetCategories handles GET /api/foliogram/templates/categories
func (c *TemplateController) GetCategories() {
	var categories []string

	err := models.DB.Model(&models.Template{}).
		Where("is_active = ?", true).
		Distinct("category").
		Pluck("category", &categories).Error

	if err != nil {
		SendInternalError(&c.Controller, map[string]interface{}{
			"message": "Failed to fetch categories",
			"error":   err.Error(),
		})
		return
	}

	SendSuccess(&c.Controller, map[string]interface{}{
		"categories": categories,
	})
}
