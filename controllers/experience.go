package controllers

import (
	"Foliogram-BE/models"

	beego "github.com/beego/beego/v2/server/web"
	"github.com/google/uuid"
)

// ExperienceController handles experience CRUD operations
type ExperienceController struct {
	beego.Controller
}

// CreateExperience creates a new experience entry for a portfolio
func (c *ExperienceController) CreateExperience() {
	portfolioID := c.Ctx.Input.Param(":portfolioId")
	userID := c.Ctx.Input.GetData("user_id").(string)

	// Verify portfolio belongs to user
	var portfolio models.Portfolio
	if err := models.DB.Where("id = ? AND user_id = ?", portfolioID, userID).First(&portfolio).Error; err != nil {
		SendNotFound(&c.Controller, "Portfolio not found")
		return
	}

	var req struct {
		Company     string `json:"company"`
		Position    string `json:"position"`
		Location    string `json:"location"`
		StartDate   string `json:"start_date"`
		EndDate     string `json:"end_date"`
		IsCurrent   bool   `json:"is_current"`
		Description string `json:"description"`
		SortOrder   int    `json:"sort_order"`
	}

	if err := c.BindJSON(&req); err != nil {
		SendBadRequestWithMessage(&c.Controller, "Invalid request body")
		return
	}

	experience := models.Experience{
		ID:          uuid.New().String(),
		PortfolioID: portfolioID,
		Company:     req.Company,
		Position:    req.Position,
		Location:    req.Location,
		StartDate:   req.StartDate,
		EndDate:     req.EndDate,
		IsCurrent:   req.IsCurrent,
		Description: req.Description,
		SortOrder:   req.SortOrder,
	}

	if err := models.DB.Create(&experience).Error; err != nil {
		SendInternalServerError(&c.Controller, "Failed to create experience entry")
		return
	}

	SendCreated(&c.Controller, experience)
}

// GetExperienceByPortfolio retrieves all experience entries for a portfolio
func (c *ExperienceController) GetExperienceByPortfolio() {
	portfolioID := c.Ctx.Input.Param(":portfolioId")
	userID := c.Ctx.Input.GetData("user_id").(string)

	// Verify portfolio belongs to user
	var portfolio models.Portfolio
	if err := models.DB.Where("id = ? AND user_id = ?", portfolioID, userID).First(&portfolio).Error; err != nil {
		SendNotFound(&c.Controller, "Portfolio not found")
		return
	}

	var experiences []models.Experience
	if err := models.DB.Where("portfolio_id = ?", portfolioID).Order("sort_order ASC, created_at DESC").Find(&experiences).Error; err != nil {
		SendInternalServerError(&c.Controller, "Failed to fetch experience entries")
		return
	}

	SendSuccess(&c.Controller, experiences)
}

// GetExperienceByID retrieves a specific experience entry
func (c *ExperienceController) GetExperienceByID() {
	portfolioID := c.Ctx.Input.Param(":portfolioId")
	experienceID := c.Ctx.Input.Param(":experienceId")
	userID := c.Ctx.Input.GetData("user_id").(string)

	// Verify portfolio belongs to user
	var portfolio models.Portfolio
	if err := models.DB.Where("id = ? AND user_id = ?", portfolioID, userID).First(&portfolio).Error; err != nil {
		SendNotFound(&c.Controller, "Portfolio not found")
		return
	}

	var experience models.Experience
	if err := models.DB.Where("id = ? AND portfolio_id = ?", experienceID, portfolioID).First(&experience).Error; err != nil {
		SendNotFound(&c.Controller, "Experience entry not found")
		return
	}

	SendSuccess(&c.Controller, experience)
}

// UpdateExperience updates an experience entry
func (c *ExperienceController) UpdateExperience() {
	portfolioID := c.Ctx.Input.Param(":portfolioId")
	experienceID := c.Ctx.Input.Param(":experienceId")
	userID := c.Ctx.Input.GetData("user_id").(string)

	// Verify portfolio belongs to user
	var portfolio models.Portfolio
	if err := models.DB.Where("id = ? AND user_id = ?", portfolioID, userID).First(&portfolio).Error; err != nil {
		SendNotFound(&c.Controller, "Portfolio not found")
		return
	}

	var experience models.Experience
	if err := models.DB.Where("id = ? AND portfolio_id = ?", experienceID, portfolioID).First(&experience).Error; err != nil {
		SendNotFound(&c.Controller, "Experience entry not found")
		return
	}

	var req struct {
		Company     string `json:"company"`
		Position    string `json:"position"`
		Location    string `json:"location"`
		StartDate   string `json:"start_date"`
		EndDate     string `json:"end_date"`
		IsCurrent   bool   `json:"is_current"`
		Description string `json:"description"`
		SortOrder   int    `json:"sort_order"`
	}

	if err := c.BindJSON(&req); err != nil {
		SendBadRequestWithMessage(&c.Controller, "Invalid request body")
		return
	}

	experience.Company = req.Company
	experience.Position = req.Position
	experience.Location = req.Location
	experience.StartDate = req.StartDate
	experience.EndDate = req.EndDate
	experience.IsCurrent = req.IsCurrent
	experience.Description = req.Description
	experience.SortOrder = req.SortOrder

	if err := models.DB.Save(&experience).Error; err != nil {
		SendInternalServerError(&c.Controller, "Failed to update experience entry")
		return
	}

	SendSuccess(&c.Controller, experience)
}

// DeleteExperience deletes an experience entry
func (c *ExperienceController) DeleteExperience() {
	portfolioID := c.Ctx.Input.Param(":portfolioId")
	experienceID := c.Ctx.Input.Param(":experienceId")
	userID := c.Ctx.Input.GetData("user_id").(string)

	// Verify portfolio belongs to user
	var portfolio models.Portfolio
	if err := models.DB.Where("id = ? AND user_id = ?", portfolioID, userID).First(&portfolio).Error; err != nil {
		SendNotFound(&c.Controller, "Portfolio not found")
		return
	}

	if err := models.DB.Where("id = ? AND portfolio_id = ?", experienceID, portfolioID).Delete(&models.Experience{}).Error; err != nil {
		SendInternalServerError(&c.Controller, "Failed to delete experience entry")
		return
	}

	SendSuccess(&c.Controller, map[string]string{"message": "Experience entry deleted successfully"})
}
