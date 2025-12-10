package controllers

import (
	"Foliogram-BE/models"

	"github.com/beego/beego/v2/server/web"
	"github.com/google/uuid"
)

type SkillController struct {
	web.Controller
}

// CreateSkill creates a new skill for a portfolio
func (c *SkillController) CreateSkill() {
	userID := c.Ctx.Input.GetData("user_id").(string)
	portfolioID := c.Ctx.Input.Param(":portfolioId")

	// Verify portfolio belongs to user
	var portfolio models.Portfolio
	if err := models.DB.Where("id = ? AND user_id = ?", portfolioID, userID).First(&portfolio).Error; err != nil {
		SendNotFound(&c.Controller, "Portfolio not found")
		return
	}

	var req struct {
		Name       string `json:"name"`
		Category   string `json:"category"`
		Level      string `json:"level"`
		Percentage int64  `json:"percentage"`
		IconURL    string `json:"icon_url"`
		SortOrder  int64  `json:"sort_order"`
	}

	if err := c.BindJSON(&req); err != nil {
		SendBadRequestWithMessage(&c.Controller, "Invalid request body")
		return
	}

	skill := models.Skill{
		ID:          uuid.New().String(),
		PortfolioID: portfolioID,
		Name:        req.Name,
		Category:    req.Category,
		Level:       req.Level,
		Percentage:  req.Percentage,
		IconURL:     req.IconURL,
		SortOrder:   req.SortOrder,
	}

	if err := models.DB.Create(&skill).Error; err != nil {
		SendInternalServerError(&c.Controller, "Failed to create skill")
		return
	}

	SendCreated(&c.Controller, skill)
}

// GetSkillsByPortfolio retrieves all skills for a portfolio
func (c *SkillController) GetSkillsByPortfolio() {
	userID := c.Ctx.Input.GetData("user_id").(string)
	portfolioID := c.Ctx.Input.Param(":portfolioId")

	// Verify portfolio belongs to user
	var portfolio models.Portfolio
	if err := models.DB.Where("id = ? AND user_id = ?", portfolioID, userID).First(&portfolio).Error; err != nil {
		SendNotFound(&c.Controller, "Portfolio not found")
		return
	}

	var skills []models.Skill
	if err := models.DB.Where("portfolio_id = ?", portfolioID).Order("sort_order ASC, created_at DESC").Find(&skills).Error; err != nil {
		SendInternalServerError(&c.Controller, "Failed to fetch skills")
		return
	}

	SendSuccess(&c.Controller, skills)
}

// GetSkillByID retrieves a specific skill
func (c *SkillController) GetSkillByID() {
	userID := c.Ctx.Input.GetData("user_id").(string)
	portfolioID := c.Ctx.Input.Param(":portfolioId")
	skillID := c.Ctx.Input.Param(":skillId")

	// Verify portfolio belongs to user
	var portfolio models.Portfolio
	if err := models.DB.Where("id = ? AND user_id = ?", portfolioID, userID).First(&portfolio).Error; err != nil {
		SendNotFound(&c.Controller, "Portfolio not found")
		return
	}

	var skill models.Skill
	if err := models.DB.Where("id = ? AND portfolio_id = ?", skillID, portfolioID).First(&skill).Error; err != nil {
		SendNotFound(&c.Controller, "Skill not found")
		return
	}

	SendSuccess(&c.Controller, skill)
}

// UpdateSkill updates an existing skill
func (c *SkillController) UpdateSkill() {
	userID := c.Ctx.Input.GetData("user_id").(string)
	portfolioID := c.Ctx.Input.Param(":portfolioId")
	skillID := c.Ctx.Input.Param(":skillId")

	// Verify portfolio belongs to user
	var portfolio models.Portfolio
	if err := models.DB.Where("id = ? AND user_id = ?", portfolioID, userID).First(&portfolio).Error; err != nil {
		SendNotFound(&c.Controller, "Portfolio not found")
		return
	}

	var skill models.Skill
	if err := models.DB.Where("id = ? AND portfolio_id = ?", skillID, portfolioID).First(&skill).Error; err != nil {
		SendNotFound(&c.Controller, "Skill not found")
		return
	}

	var req struct {
		Name       string `json:"name"`
		Category   string `json:"category"`
		Level      string `json:"level"`
		Percentage int64  `json:"percentage"`
		IconURL    string `json:"icon_url"`
		SortOrder  int64  `json:"sort_order"`
	}

	if err := c.BindJSON(&req); err != nil {
		SendBadRequestWithMessage(&c.Controller, "Invalid request body")
		return
	}

	// Update fields
	if req.Name != "" {
		skill.Name = req.Name
	}
	skill.Category = req.Category
	skill.Level = req.Level
	skill.Percentage = req.Percentage
	skill.IconURL = req.IconURL
	skill.SortOrder = req.SortOrder

	if err := models.DB.Save(&skill).Error; err != nil {
		SendInternalServerError(&c.Controller, "Failed to update skill")
		return
	}

	SendSuccess(&c.Controller, skill)
}

// DeleteSkill deletes a skill
func (c *SkillController) DeleteSkill() {
	userID := c.Ctx.Input.GetData("user_id").(string)
	portfolioID := c.Ctx.Input.Param(":portfolioId")
	skillID := c.Ctx.Input.Param(":skillId")

	// Verify portfolio belongs to user
	var portfolio models.Portfolio
	if err := models.DB.Where("id = ? AND user_id = ?", portfolioID, userID).First(&portfolio).Error; err != nil {
		SendNotFound(&c.Controller, "Portfolio not found")
		return
	}

	var skill models.Skill
	if err := models.DB.Where("id = ? AND portfolio_id = ?", skillID, portfolioID).First(&skill).Error; err != nil {
		SendNotFound(&c.Controller, "Skill not found")
		return
	}

	if err := models.DB.Delete(&skill).Error; err != nil {
		SendInternalServerError(&c.Controller, "Failed to delete skill")
		return
	}

	SendSuccess(&c.Controller, map[string]string{"message": "Skill deleted successfully"})
}
