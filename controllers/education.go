package controllers

import (
	"Foliogram-BE/models"

	beego "github.com/beego/beego/v2/server/web"
	"github.com/google/uuid"
)

// EducationController handles education CRUD operations
type EducationController struct {
	beego.Controller
}

// CreateEducation creates a new education entry for a portfolio
func (c *EducationController) CreateEducation() {
	portfolioID := c.Ctx.Input.Param(":portfolioId")
	userID := c.Ctx.Input.GetData("user_id").(string)

	// Verify portfolio belongs to user
	var portfolio models.Portfolio
	if err := models.DB.Where("id = ? AND user_id = ?", portfolioID, userID).First(&portfolio).Error; err != nil {
		SendNotFound(&c.Controller, "Portfolio not found")
		return
	}

	var req struct {
		Institution string `json:"institution"`
		Degree      string `json:"degree"`
		Field       string `json:"field"`
		StartDate   string `json:"start_date"`
		EndDate     string `json:"end_date"`
		Grade       string `json:"grade"`
		Description string `json:"description"`
		SortOrder   int    `json:"sort_order"`
	}

	if err := c.BindJSON(&req); err != nil {
		SendBadRequestWithMessage(&c.Controller, "Invalid request body")
		return
	}

	education := models.Education{
		ID:          uuid.New().String(),
		PortfolioID: portfolioID,
		Institution: req.Institution,
		Degree:      req.Degree,
		Field:       req.Field,
		StartDate:   req.StartDate,
		EndDate:     req.EndDate,
		Grade:       req.Grade,
		Description: req.Description,
		SortOrder:   req.SortOrder,
	}

	if err := models.DB.Create(&education).Error; err != nil {
		SendInternalServerError(&c.Controller, "Failed to create education entry")
		return
	}

	SendCreated(&c.Controller, education)
}

// GetEducationByPortfolio retrieves all education entries for a portfolio
func (c *EducationController) GetEducationByPortfolio() {
	portfolioID := c.Ctx.Input.Param(":portfolioId")
	userID := c.Ctx.Input.GetData("user_id").(string)

	// Verify portfolio belongs to user
	var portfolio models.Portfolio
	if err := models.DB.Where("id = ? AND user_id = ?", portfolioID, userID).First(&portfolio).Error; err != nil {
		SendNotFound(&c.Controller, "Portfolio not found")
		return
	}

	var education []models.Education
	if err := models.DB.Where("portfolio_id = ?", portfolioID).Order("sort_order ASC, created_at DESC").Find(&education).Error; err != nil {
		SendInternalServerError(&c.Controller, "Failed to fetch education entries")
		return
	}

	SendSuccess(&c.Controller, education)
}

// GetEducationByID retrieves a specific education entry
func (c *EducationController) GetEducationByID() {
	portfolioID := c.Ctx.Input.Param(":portfolioId")
	educationID := c.Ctx.Input.Param(":educationId")
	userID := c.Ctx.Input.GetData("user_id").(string)

	// Verify portfolio belongs to user
	var portfolio models.Portfolio
	if err := models.DB.Where("id = ? AND user_id = ?", portfolioID, userID).First(&portfolio).Error; err != nil {
		SendNotFound(&c.Controller, "Portfolio not found")
		return
	}

	var education models.Education
	if err := models.DB.Where("id = ? AND portfolio_id = ?", educationID, portfolioID).First(&education).Error; err != nil {
		SendNotFound(&c.Controller, "Education entry not found")
		return
	}

	SendSuccess(&c.Controller, education)
}

// UpdateEducation updates an education entry
func (c *EducationController) UpdateEducation() {
	portfolioID := c.Ctx.Input.Param(":portfolioId")
	educationID := c.Ctx.Input.Param(":educationId")
	userID := c.Ctx.Input.GetData("user_id").(string)

	// Verify portfolio belongs to user
	var portfolio models.Portfolio
	if err := models.DB.Where("id = ? AND user_id = ?", portfolioID, userID).First(&portfolio).Error; err != nil {
		SendNotFound(&c.Controller, "Portfolio not found")
		return
	}

	var education models.Education
	if err := models.DB.Where("id = ? AND portfolio_id = ?", educationID, portfolioID).First(&education).Error; err != nil {
		SendNotFound(&c.Controller, "Education entry not found")
		return
	}

	var req struct {
		Institution string `json:"institution"`
		Degree      string `json:"degree"`
		Field       string `json:"field"`
		StartDate   string `json:"start_date"`
		EndDate     string `json:"end_date"`
		Grade       string `json:"grade"`
		Description string `json:"description"`
		SortOrder   int    `json:"sort_order"`
	}

	if err := c.BindJSON(&req); err != nil {
		SendBadRequestWithMessage(&c.Controller, "Invalid request body")
		return
	}

	education.Institution = req.Institution
	education.Degree = req.Degree
	education.Field = req.Field
	education.StartDate = req.StartDate
	education.EndDate = req.EndDate
	education.Grade = req.Grade
	education.Description = req.Description
	education.SortOrder = req.SortOrder

	if err := models.DB.Save(&education).Error; err != nil {
		SendInternalServerError(&c.Controller, "Failed to update education entry")
		return
	}

	SendSuccess(&c.Controller, education)
}

// DeleteEducation deletes an education entry
func (c *EducationController) DeleteEducation() {
	portfolioID := c.Ctx.Input.Param(":portfolioId")
	educationID := c.Ctx.Input.Param(":educationId")
	userID := c.Ctx.Input.GetData("user_id").(string)

	// Verify portfolio belongs to user
	var portfolio models.Portfolio
	if err := models.DB.Where("id = ? AND user_id = ?", portfolioID, userID).First(&portfolio).Error; err != nil {
		SendNotFound(&c.Controller, "Portfolio not found")
		return
	}

	if err := models.DB.Where("id = ? AND portfolio_id = ?", educationID, portfolioID).Delete(&models.Education{}).Error; err != nil {
		SendInternalServerError(&c.Controller, "Failed to delete education entry")
		return
	}

	SendSuccess(&c.Controller, map[string]string{"message": "Education entry deleted successfully"})
}
