package controllers

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"Foliogram-BE/models"
	"Foliogram-BE/utils"

	beego "github.com/beego/beego/v2/server/web"
)

// PortfolioController handles portfolio operations
type PortfolioController struct {
	beego.Controller
}

// CreatePortfolioRequest represents the request to create a portfolio
type CreatePortfolioRequest struct {
	TemplateID  string `json:"template_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

// UpdatePortfolioRequest represents the request to update a portfolio
type UpdatePortfolioRequest struct {
	Title           string                 `json:"title"`
	Description     string                 `json:"description"`
	Config          map[string]interface{} `json:"config"`
	IsPublished     *bool                  `json:"is_published"`
	MetaTitle       string                 `json:"meta_title"`
	MetaDescription string                 `json:"meta_description"`
	MetaKeywords    string                 `json:"meta_keywords"`
}

// Create handles POST /api/foliogram/portfolios
func (c *PortfolioController) Create() {
	var req CreatePortfolioRequest

	// Parse JSON body
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &req); err != nil {
		SendBadRequest(&c.Controller, map[string]interface{}{
			"message": "Invalid JSON format",
			"error":   err.Error(),
		})
		return
	}

	// Get user ID from context (set by auth middleware)
	userID := c.Ctx.Input.GetData("user_id")
	if userID == nil {
		SendUnauthorized(&c.Controller, map[string]interface{}{
			"message": "User not authenticated",
		})
		return
	}

	// Validate required fields
	if req.TemplateID == "" || req.Title == "" {
		SendBadRequest(&c.Controller, map[string]interface{}{
			"message": "template_id and title are required",
		})
		return
	}

	// Verify template exists
	var template models.Template
	if err := models.DB.First(&template, "id = ?", req.TemplateID).Error; err != nil {
		SendNotFound(&c.Controller, map[string]interface{}{
			"message": "Template not found",
		})
		return
	}

	// Generate unique slug from title
	slug := generateSlug(req.Title)

	// Check if slug already exists, append number if it does
	originalSlug := slug
	counter := 1
	for {
		var existingPortfolio models.Portfolio
		err := models.DB.Where("slug = ?", slug).First(&existingPortfolio).Error
		if err != nil {
			// Slug is available
			break
		}
		slug = fmt.Sprintf("%s-%d", originalSlug, counter)
		counter++
	}

	// Create portfolio
	portfolio := models.Portfolio{
		UserID:      userID.(string),
		TemplateID:  req.TemplateID,
		Title:       req.Title,
		Slug:        slug,
		Description: req.Description,
		Config:      template.DefaultConfig, // Use template's default config
		IsPublished: false,
	}

	if err := models.DB.Create(&portfolio).Error; err != nil {
		SendInternalError(&c.Controller, map[string]interface{}{
			"message": "Failed to create portfolio",
			"error":   err.Error(),
		})
		return
	}

	// Preload template data
	models.DB.Preload("Template").First(&portfolio, "id = ?", portfolio.ID)

	SendCreated(&c.Controller, portfolio)
}

// GetByID handles GET /api/foliogram/portfolios/:id
func (c *PortfolioController) GetByID() {
	portfolioID := c.Ctx.Input.Param(":id")

	var portfolio models.Portfolio
	err := models.DB.
		Preload("Template").
		Preload("Projects").
		Preload("Skills").
		Preload("Testimonials").
		First(&portfolio, "id = ?", portfolioID).Error

	if err != nil {
		SendNotFound(&c.Controller, map[string]interface{}{
			"message": "Portfolio not found",
		})
		return
	}

	// Check if user has permission to view (owner or published)
	userID := c.Ctx.Input.GetData("user_id")
	if !portfolio.IsPublished && (userID == nil || userID.(string) != portfolio.UserID) {
		SendForbidden(&c.Controller, map[string]interface{}{
			"message": "You don't have permission to view this portfolio",
		})
		return
	}

	// Increment view count for published portfolios
	if portfolio.IsPublished {
		models.DB.Model(&portfolio).Update("view_count", portfolio.ViewCount+1)
	}

	SendSuccess(&c.Controller, portfolio)
}

// GetBySlug handles GET /api/foliogram/portfolios/slug/:slug
func (c *PortfolioController) GetBySlug() {
	slug := c.Ctx.Input.Param(":slug")

	var portfolio models.Portfolio
	err := models.DB.
		Preload("Template").
		Preload("Projects").
		Preload("Skills").
		Preload("Testimonials").
		Preload("User").
		First(&portfolio, "slug = ?", slug).Error

	if err != nil {
		SendNotFound(&c.Controller, map[string]interface{}{
			"message": "Portfolio not found",
		})
		return
	}

	// Check if user has permission to view (owner or published)
	userID := c.Ctx.Input.GetData("user_id")
	if !portfolio.IsPublished && (userID == nil || userID.(string) != portfolio.UserID) {
		SendForbidden(&c.Controller, map[string]interface{}{
			"message": "Portfolio not published",
		})
		return
	}

	// Sanitize user data before returning
	portfolio.User.Sanitize()

	// Increment view count for published portfolios
	if portfolio.IsPublished {
		models.DB.Model(&portfolio).Update("view_count", portfolio.ViewCount+1)
	}

	SendSuccess(&c.Controller, portfolio)
}

// GetUserPortfolios handles GET /api/foliogram/portfolios/my
func (c *PortfolioController) GetUserPortfolios() {
	// Get user ID from context
	userID := c.Ctx.Input.GetData("user_id")
	if userID == nil {
		SendUnauthorized(&c.Controller, map[string]interface{}{
			"message": "User not authenticated",
		})
		return
	}

	var portfolios []models.Portfolio
	err := models.DB.
		Preload("Template").
		Where("user_id = ?", userID.(string)).
		Order("created_at DESC").
		Find(&portfolios).Error

	if err != nil {
		SendInternalError(&c.Controller, map[string]interface{}{
			"message": "Failed to fetch portfolios",
			"error":   err.Error(),
		})
		return
	}

	SendSuccess(&c.Controller, portfolios)
}

// Update handles PUT /api/foliogram/portfolios/:id
func (c *PortfolioController) Update() {
	portfolioID := c.Ctx.Input.Param(":id")

	// Get user ID from context
	userID := c.Ctx.Input.GetData("user_id")
	if userID == nil {
		SendUnauthorized(&c.Controller, map[string]interface{}{
			"message": "User not authenticated",
		})
		return
	}

	// Find portfolio
	var portfolio models.Portfolio
	if err := models.DB.First(&portfolio, "id = ?", portfolioID).Error; err != nil {
		SendNotFound(&c.Controller, map[string]interface{}{
			"message": "Portfolio not found",
		})
		return
	}

	// Check ownership
	if portfolio.UserID != userID.(string) {
		SendForbidden(&c.Controller, map[string]interface{}{
			"message": "You don't have permission to update this portfolio",
		})
		return
	}

	// Parse update request
	var req UpdatePortfolioRequest
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &req); err != nil {
		SendBadRequest(&c.Controller, map[string]interface{}{
			"message": "Invalid JSON format",
			"error":   err.Error(),
		})
		return
	}

	// Update fields
	updates := make(map[string]interface{})

	if req.Title != "" {
		updates["title"] = req.Title
	}
	if req.Description != "" {
		updates["description"] = req.Description
	}
	if req.Config != nil {
		configJSON, _ := json.Marshal(req.Config)
		updates["config"] = string(configJSON)
	}
	if req.IsPublished != nil {
		updates["is_published"] = *req.IsPublished
	}
	if req.MetaTitle != "" {
		updates["meta_title"] = req.MetaTitle
	}
	if req.MetaDescription != "" {
		updates["meta_description"] = req.MetaDescription
	}
	if req.MetaKeywords != "" {
		updates["meta_keywords"] = req.MetaKeywords
	}

	// Perform update
	if err := models.DB.Model(&portfolio).Updates(updates).Error; err != nil {
		SendInternalError(&c.Controller, map[string]interface{}{
			"message": "Failed to update portfolio",
			"error":   err.Error(),
		})
		return
	}

	// Reload with associations
	models.DB.
		Preload("Template").
		Preload("Projects").
		Preload("Skills").
		Preload("Testimonials").
		First(&portfolio, "id = ?", portfolioID)

	SendSuccess(&c.Controller, portfolio)
}

// Delete handles DELETE /api/foliogram/portfolios/:id
func (c *PortfolioController) Delete() {
	portfolioID := c.Ctx.Input.Param(":id")

	// Get user ID from context
	userID := c.Ctx.Input.GetData("user_id")
	if userID == nil {
		SendUnauthorized(&c.Controller, map[string]interface{}{
			"message": "User not authenticated",
		})
		return
	}

	// Find portfolio
	var portfolio models.Portfolio
	if err := models.DB.First(&portfolio, "id = ?", portfolioID).Error; err != nil {
		SendNotFound(&c.Controller, map[string]interface{}{
			"message": "Portfolio not found",
		})
		return
	}

	// Check ownership
	if portfolio.UserID != userID.(string) {
		SendForbidden(&c.Controller, map[string]interface{}{
			"message": "You don't have permission to delete this portfolio",
		})
		return
	}

	// Soft delete
	if err := models.DB.Delete(&portfolio).Error; err != nil {
		SendInternalError(&c.Controller, map[string]interface{}{
			"message": "Failed to delete portfolio",
			"error":   err.Error(),
		})
		return
	}

	SendSuccess(&c.Controller, map[string]interface{}{
		"message": "Portfolio deleted successfully",
	})
}

// GetPublicByID handles GET /api/public/portfolios/:id (no auth required)
func (c *PortfolioController) GetPublicByID() {
	portfolioID := c.Ctx.Input.Param(":id")

	var portfolio models.Portfolio
	err := models.DB.
		Preload("Template").
		Preload("Projects").
		Preload("Skills").
		Preload("Testimonials").
		Preload("User").
		First(&portfolio, "id = ? AND is_published = ?", portfolioID, true).Error

	if err != nil {
		SendNotFound(&c.Controller, map[string]interface{}{
			"message": "Portfolio not found or not published",
		})
		return
	}

	// Sanitize user data before returning
	portfolio.User.Sanitize()

	// Track analytics
	trackPortfolioView(&c.Controller, portfolio.ID)

	// Increment view count
	models.DB.Model(&portfolio).UpdateColumn("view_count", portfolio.ViewCount+1)

	SendSuccess(&c.Controller, portfolio)
}

// GetPublicBySlug handles GET /api/public/portfolios/slug/:slug (no auth required)
func (c *PortfolioController) GetPublicBySlug() {
	slug := c.Ctx.Input.Param(":slug")

	var portfolio models.Portfolio
	err := models.DB.
		Preload("Template").
		Preload("Projects").
		Preload("Skills").
		Preload("Testimonials").
		Preload("User").
		First(&portfolio, "slug = ? AND is_published = ?", slug, true).Error

	if err != nil {
		SendNotFound(&c.Controller, map[string]interface{}{
			"message": "Portfolio not found or not published",
		})
		return
	}

	// Sanitize user data before returning
	portfolio.User.Sanitize()

	// Track analytics
	trackPortfolioView(&c.Controller, portfolio.ID)

	// Increment view count
	models.DB.Model(&portfolio).UpdateColumn("view_count", portfolio.ViewCount+1)

	SendSuccess(&c.Controller, portfolio)
}

// GetPublicPortfolios handles GET /api/public/portfolios (no auth required)
// Returns a list of all published portfolios
func (c *PortfolioController) GetPublicPortfolios() {
	// Get query parameters for pagination
	page, _ := c.GetInt("page", 1)
	limit, _ := c.GetInt("limit", 20)
	category := c.GetString("category")
	search := c.GetString("search")

	// Validate pagination
	if page < 1 {
		page = 1
	}
	if limit < 1 || limit > 100 {
		limit = 20
	}

	offset := (page - 1) * limit

	query := models.DB.
		Preload("Template").
		Preload("User").
		Where("is_published = ?", true)

	// Filter by category if provided
	if category != "" {
		query = query.Joins("JOIN templates ON templates.id = portfolios.template_id").
			Where("templates.category = ?", category)
	}

	// Search by title or description
	if search != "" {
		searchPattern := "%" + search + "%"
		query = query.Where("title LIKE ? OR description LIKE ?", searchPattern, searchPattern)
	}

	var portfolios []models.Portfolio
	var total int64

	// Get total count
	query.Model(&models.Portfolio{}).Count(&total)

	// Get paginated results
	err := query.
		Order("created_at DESC").
		Limit(limit).
		Offset(offset).
		Find(&portfolios).Error

	if err != nil {
		SendInternalError(&c.Controller, map[string]interface{}{
			"message": "Failed to fetch portfolios",
			"error":   err.Error(),
		})
		return
	}

	// Sanitize user data
	for i := range portfolios {
		portfolios[i].User.Sanitize()
	}

	SendSuccess(&c.Controller, map[string]interface{}{
		"portfolios": portfolios,
		"pagination": map[string]interface{}{
			"page":       page,
			"limit":      limit,
			"total":      total,
			"totalPages": (total + int64(limit) - 1) / int64(limit),
		},
	})
}

// Helper function to track portfolio views
func trackPortfolioView(c *beego.Controller, portfolioID string) {
	// Get viewer information
	ip := c.Ctx.Input.IP()
	userAgent := c.Ctx.Input.Header("User-Agent")
	referrer := c.Ctx.Input.Header("Referer")

	// Generate or get session ID from cookie
	sessionID := c.Ctx.GetCookie("session_id")
	if sessionID == "" {
		sessionID = fmt.Sprintf("%d-%s", time.Now().Unix(), ip)
		c.Ctx.SetCookie("session_id", sessionID, 86400) // 24 hours
	}

	// Create analytics record
	analyticsView := models.AnalyticsView{
		PortfolioID: portfolioID,
		ViewerIP:    ip,
		SessionID:   sessionID,
		UserAgent:   userAgent,
		Referrer:    referrer,
	}

	// Save asynchronously to avoid blocking the response
	go func() {
		if err := models.DB.Create(&analyticsView).Error; err != nil {
			fmt.Printf("Failed to log analytics view: %v\n", err)
		}
	}()
}

// Helper function to generate URL-friendly slug
func generateSlug(title string) string {
	// Convert to lowercase
	slug := strings.ToLower(title)

	// Replace spaces with hyphens
	slug = strings.ReplaceAll(slug, " ", "-")

	// Remove special characters (keep only alphanumeric and hyphens)
	slug = utils.SanitizeInput(slug)
	var result []rune
	for _, char := range slug {
		if (char >= 'a' && char <= 'z') || (char >= '0' && char <= '9') || char == '-' {
			result = append(result, char)
		}
	}
	slug = string(result)

	// Remove consecutive hyphens
	for strings.Contains(slug, "--") {
		slug = strings.ReplaceAll(slug, "--", "-")
	}

	// Trim hyphens from start and end
	slug = strings.Trim(slug, "-")

	return slug
}
