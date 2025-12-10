package controllers

import (
	"Foliogram-BE/models"

	beego "github.com/beego/beego/v2/server/web"
	"github.com/google/uuid"
)

// PublicationController handles publication CRUD operations
type PublicationController struct {
	beego.Controller
}

// CreatePublication creates a new publication entry for a portfolio
func (c *PublicationController) CreatePublication() {
	portfolioID := c.Ctx.Input.Param(":portfolioId")
	userID := c.Ctx.Input.GetData("user_id").(string)

	// Verify portfolio belongs to user
	var portfolio models.Portfolio
	if err := models.DB.Where("id = ? AND user_id = ?", portfolioID, userID).First(&portfolio).Error; err != nil {
		SendNotFound(&c.Controller, "Portfolio not found")
		return
	}

	var req struct {
		Title      string `json:"title"`
		Authors    string `json:"authors"`
		Journal    string `json:"journal"`
		Conference string `json:"conference"`
		Year       string `json:"year"`
		DOI        string `json:"doi"`
		URL        string `json:"url"`
		Abstract   string `json:"abstract"`
		Type       string `json:"type"`
		SortOrder  int    `json:"sort_order"`
	}

	if err := c.BindJSON(&req); err != nil {
		SendBadRequestWithMessage(&c.Controller, "Invalid request body")
		return
	}

	publication := models.Publication{
		ID:          uuid.New().String(),
		PortfolioID: portfolioID,
		Title:       req.Title,
		Authors:     req.Authors,
		Journal:     req.Journal,
		Conference:  req.Conference,
		Year:        req.Year,
		DOI:         req.DOI,
		URL:         req.URL,
		Abstract:    req.Abstract,
		Type:        req.Type,
		SortOrder:   req.SortOrder,
	}

	if err := models.DB.Create(&publication).Error; err != nil {
		SendInternalServerError(&c.Controller, "Failed to create publication entry")
		return
	}

	SendCreated(&c.Controller, publication)
}

// GetPublicationsByPortfolio retrieves all publication entries for a portfolio
func (c *PublicationController) GetPublicationsByPortfolio() {
	portfolioID := c.Ctx.Input.Param(":portfolioId")
	userID := c.Ctx.Input.GetData("user_id").(string)

	// Verify portfolio belongs to user
	var portfolio models.Portfolio
	if err := models.DB.Where("id = ? AND user_id = ?", portfolioID, userID).First(&portfolio).Error; err != nil {
		SendNotFound(&c.Controller, "Portfolio not found")
		return
	}

	var publications []models.Publication
	if err := models.DB.Where("portfolio_id = ?", portfolioID).Order("sort_order ASC, year DESC, created_at DESC").Find(&publications).Error; err != nil {
		SendInternalServerError(&c.Controller, "Failed to fetch publication entries")
		return
	}

	SendSuccess(&c.Controller, publications)
}

// GetPublicationByID retrieves a specific publication entry
func (c *PublicationController) GetPublicationByID() {
	portfolioID := c.Ctx.Input.Param(":portfolioId")
	publicationID := c.Ctx.Input.Param(":publicationId")
	userID := c.Ctx.Input.GetData("user_id").(string)

	// Verify portfolio belongs to user
	var portfolio models.Portfolio
	if err := models.DB.Where("id = ? AND user_id = ?", portfolioID, userID).First(&portfolio).Error; err != nil {
		SendNotFound(&c.Controller, "Portfolio not found")
		return
	}

	var publication models.Publication
	if err := models.DB.Where("id = ? AND portfolio_id = ?", publicationID, portfolioID).First(&publication).Error; err != nil {
		SendNotFound(&c.Controller, "Publication entry not found")
		return
	}

	SendSuccess(&c.Controller, publication)
}

// UpdatePublication updates a publication entry
func (c *PublicationController) UpdatePublication() {
	portfolioID := c.Ctx.Input.Param(":portfolioId")
	publicationID := c.Ctx.Input.Param(":publicationId")
	userID := c.Ctx.Input.GetData("user_id").(string)

	// Verify portfolio belongs to user
	var portfolio models.Portfolio
	if err := models.DB.Where("id = ? AND user_id = ?", portfolioID, userID).First(&portfolio).Error; err != nil {
		SendNotFound(&c.Controller, "Portfolio not found")
		return
	}

	var publication models.Publication
	if err := models.DB.Where("id = ? AND portfolio_id = ?", publicationID, portfolioID).First(&publication).Error; err != nil {
		SendNotFound(&c.Controller, "Publication entry not found")
		return
	}

	var req struct {
		Title      string `json:"title"`
		Authors    string `json:"authors"`
		Journal    string `json:"journal"`
		Conference string `json:"conference"`
		Year       string `json:"year"`
		DOI        string `json:"doi"`
		URL        string `json:"url"`
		Abstract   string `json:"abstract"`
		Type       string `json:"type"`
		SortOrder  int    `json:"sort_order"`
	}

	if err := c.BindJSON(&req); err != nil {
		SendBadRequestWithMessage(&c.Controller, "Invalid request body")
		return
	}

	publication.Title = req.Title
	publication.Authors = req.Authors
	publication.Journal = req.Journal
	publication.Conference = req.Conference
	publication.Year = req.Year
	publication.DOI = req.DOI
	publication.URL = req.URL
	publication.Abstract = req.Abstract
	publication.Type = req.Type
	publication.SortOrder = req.SortOrder

	if err := models.DB.Save(&publication).Error; err != nil {
		SendInternalServerError(&c.Controller, "Failed to update publication entry")
		return
	}

	SendSuccess(&c.Controller, publication)
}

// DeletePublication deletes a publication entry
func (c *PublicationController) DeletePublication() {
	portfolioID := c.Ctx.Input.Param(":portfolioId")
	publicationID := c.Ctx.Input.Param(":publicationId")
	userID := c.Ctx.Input.GetData("user_id").(string)

	// Verify portfolio belongs to user
	var portfolio models.Portfolio
	if err := models.DB.Where("id = ? AND user_id = ?", portfolioID, userID).First(&portfolio).Error; err != nil {
		SendNotFound(&c.Controller, "Portfolio not found")
		return
	}

	if err := models.DB.Where("id = ? AND portfolio_id = ?", publicationID, portfolioID).Delete(&models.Publication{}).Error; err != nil {
		SendInternalServerError(&c.Controller, "Failed to delete publication entry")
		return
	}

	SendSuccess(&c.Controller, map[string]string{"message": "Publication entry deleted successfully"})
}
