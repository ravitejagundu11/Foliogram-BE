package controllers

import (
	"Foliogram-BE/models"

	"github.com/beego/beego/v2/server/web"
	"github.com/google/uuid"
)

type TestimonialController struct {
	web.Controller
}

// CreateTestimonial creates a new testimonial for a portfolio
func (c *TestimonialController) CreateTestimonial() {
	userID := c.Ctx.Input.GetData("user_id").(string)
	portfolioID := c.Ctx.Input.Param(":portfolioId")

	// Verify portfolio belongs to user
	var portfolio models.Portfolio
	if err := models.DB.Where("id = ? AND user_id = ?", portfolioID, userID).First(&portfolio).Error; err != nil {
		SendNotFound(&c.Controller, "Portfolio not found")
		return
	}

	var req struct {
		Name        string `json:"name"`
		Position    string `json:"position"`
		Company     string `json:"company"`
		Content     string `json:"content"`
		Rating      int64  `json:"rating"`
		AvatarURL   string `json:"avatar_url"`
		LinkedInURL string `json:"linked_in_url"`
		IsVisible   bool   `json:"is_visible"`
		SortOrder   int64  `json:"sort_order"`
	}

	if err := c.BindJSON(&req); err != nil {
		SendBadRequestWithMessage(&c.Controller, "Invalid request body")
		return
	}

	testimonial := models.Testimonial{
		ID:          uuid.New().String(),
		PortfolioID: portfolioID,
		Name:        req.Name,
		Position:    req.Position,
		Company:     req.Company,
		Content:     req.Content,
		Rating:      req.Rating,
		AvatarURL:   req.AvatarURL,
		LinkedInURL: req.LinkedInURL,
		IsVisible:   req.IsVisible,
		SortOrder:   req.SortOrder,
	}

	if err := models.DB.Create(&testimonial).Error; err != nil {
		SendInternalServerError(&c.Controller, "Failed to create testimonial")
		return
	}

	SendCreated(&c.Controller, testimonial)
}

// GetTestimonialsByPortfolio retrieves all testimonials for a portfolio
func (c *TestimonialController) GetTestimonialsByPortfolio() {
	userID := c.Ctx.Input.GetData("user_id").(string)
	portfolioID := c.Ctx.Input.Param(":portfolioId")

	// Verify portfolio belongs to user
	var portfolio models.Portfolio
	if err := models.DB.Where("id = ? AND user_id = ?", portfolioID, userID).First(&portfolio).Error; err != nil {
		SendNotFound(&c.Controller, "Portfolio not found")
		return
	}

	var testimonials []models.Testimonial
	if err := models.DB.Where("portfolio_id = ?", portfolioID).Order("sort_order ASC, created_at DESC").Find(&testimonials).Error; err != nil {
		SendInternalServerError(&c.Controller, "Failed to fetch testimonials")
		return
	}

	SendSuccess(&c.Controller, testimonials)
}

// GetTestimonialByID retrieves a specific testimonial
func (c *TestimonialController) GetTestimonialByID() {
	userID := c.Ctx.Input.GetData("user_id").(string)
	portfolioID := c.Ctx.Input.Param(":portfolioId")
	testimonialID := c.Ctx.Input.Param(":testimonialId")

	// Verify portfolio belongs to user
	var portfolio models.Portfolio
	if err := models.DB.Where("id = ? AND user_id = ?", portfolioID, userID).First(&portfolio).Error; err != nil {
		SendNotFound(&c.Controller, "Portfolio not found")
		return
	}

	var testimonial models.Testimonial
	if err := models.DB.Where("id = ? AND portfolio_id = ?", testimonialID, portfolioID).First(&testimonial).Error; err != nil {
		SendNotFound(&c.Controller, "Testimonial not found")
		return
	}

	SendSuccess(&c.Controller, testimonial)
}

// UpdateTestimonial updates an existing testimonial
func (c *TestimonialController) UpdateTestimonial() {
	userID := c.Ctx.Input.GetData("user_id").(string)
	portfolioID := c.Ctx.Input.Param(":portfolioId")
	testimonialID := c.Ctx.Input.Param(":testimonialId")

	// Verify portfolio belongs to user
	var portfolio models.Portfolio
	if err := models.DB.Where("id = ? AND user_id = ?", portfolioID, userID).First(&portfolio).Error; err != nil {
		SendNotFound(&c.Controller, "Portfolio not found")
		return
	}

	var testimonial models.Testimonial
	if err := models.DB.Where("id = ? AND portfolio_id = ?", testimonialID, portfolioID).First(&testimonial).Error; err != nil {
		SendNotFound(&c.Controller, "Testimonial not found")
		return
	}

	var req struct {
		Name        string `json:"name"`
		Position    string `json:"position"`
		Company     string `json:"company"`
		Content     string `json:"content"`
		Rating      int64  `json:"rating"`
		AvatarURL   string `json:"avatar_url"`
		LinkedInURL string `json:"linked_in_url"`
		IsVisible   bool   `json:"is_visible"`
		SortOrder   int64  `json:"sort_order"`
	}

	if err := c.BindJSON(&req); err != nil {
		SendBadRequestWithMessage(&c.Controller, "Invalid request body")
		return
	}

	// Update fields
	if req.Name != "" {
		testimonial.Name = req.Name
	}
	testimonial.Position = req.Position
	testimonial.Company = req.Company
	if req.Content != "" {
		testimonial.Content = req.Content
	}
	testimonial.Rating = req.Rating
	testimonial.AvatarURL = req.AvatarURL
	testimonial.LinkedInURL = req.LinkedInURL
	testimonial.IsVisible = req.IsVisible
	testimonial.SortOrder = req.SortOrder

	if err := models.DB.Save(&testimonial).Error; err != nil {
		SendInternalServerError(&c.Controller, "Failed to update testimonial")
		return
	}

	SendSuccess(&c.Controller, testimonial)
}

// DeleteTestimonial deletes a testimonial
func (c *TestimonialController) DeleteTestimonial() {
	userID := c.Ctx.Input.GetData("user_id").(string)
	portfolioID := c.Ctx.Input.Param(":portfolioId")
	testimonialID := c.Ctx.Input.Param(":testimonialId")

	// Verify portfolio belongs to user
	var portfolio models.Portfolio
	if err := models.DB.Where("id = ? AND user_id = ?", portfolioID, userID).First(&portfolio).Error; err != nil {
		SendNotFound(&c.Controller, "Portfolio not found")
		return
	}

	var testimonial models.Testimonial
	if err := models.DB.Where("id = ? AND portfolio_id = ?", testimonialID, portfolioID).First(&testimonial).Error; err != nil {
		SendNotFound(&c.Controller, "Testimonial not found")
		return
	}

	if err := models.DB.Delete(&testimonial).Error; err != nil {
		SendInternalServerError(&c.Controller, "Failed to delete testimonial")
		return
	}

	SendSuccess(&c.Controller, map[string]string{"message": "Testimonial deleted successfully"})
}
