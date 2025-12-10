package controllers

import (
	"Foliogram-BE/models"
	"encoding/json"

	"github.com/beego/beego/v2/server/web"
	"github.com/google/uuid"
)

type PortfolioController struct {
	web.Controller
}

// CreatePortfolio creates a new portfolio for the authenticated user
func (c *PortfolioController) CreatePortfolio() {
	userID := c.Ctx.Input.GetData("user_id").(string)

	var req struct {
		TemplateID     string                 `json:"template_id"`
		Title          string                 `json:"title"`
		Headline       string                 `json:"headline"`
		Slug           string                 `json:"slug"`
		Description    string                 `json:"description"`
		ProfilePicture string                 `json:"profile_picture"`
		PrimaryColor   string                 `json:"primary_color"`
		SecondaryColor string                 `json:"secondary_color"`
		HeadingFont    string                 `json:"heading_font"`
		BodyFont       string                 `json:"body_font"`
		LayoutStyle    interface{}            `json:"layout_style"`
		SocialLinks    map[string]interface{} `json:"social_links"`
	}

	if err := c.BindJSON(&req); err != nil {
		SendBadRequestWithMessage(&c.Controller, "Invalid request body")
		return
	}

	// Store styling and metadata in config JSON
	configJSON := map[string]interface{}{
		"headline":       req.Headline,
		"profilePicture": req.ProfilePicture,
		"primaryColor":   req.PrimaryColor,
		"secondaryColor": req.SecondaryColor,
		"headingFont":    req.HeadingFont,
		"bodyFont":       req.BodyFont,
		"layoutStyle":    req.LayoutStyle,
		"socialLinks":    req.SocialLinks,
	}
	configStr, _ := json.Marshal(configJSON)

	// Use title or headline for portfolio title
	title := req.Title
	if title == "" {
		title = req.Headline
	}
	if title == "" {
		title = "My Portfolio"
	}

	// Validate slug is not empty
	slug := req.Slug
	if slug == "" {
		SendBadRequestWithMessage(&c.Controller, "Slug is required")
		return
	}

	portfolio := models.Portfolio{
		ID:          uuid.New().String(),
		UserID:      userID,
		TemplateID:  req.TemplateID,
		Title:       title,
		Slug:        slug,
		Description: req.Description,
		Config:      string(configStr),
		IsPublished: false,
	}

	if err := models.DB.Create(&portfolio).Error; err != nil {
		SendInternalServerError(&c.Controller, "Failed to create portfolio")
		return
	}

	SendCreated(&c.Controller, portfolio)
}

// GetMyPortfolios retrieves all portfolios for the authenticated user
func (c *PortfolioController) GetMyPortfolios() {
	userID := c.Ctx.Input.GetData("user_id").(string)

	var portfolios []models.Portfolio
	if err := models.DB.Where("user_id = ?", userID).Order("created_at DESC").Find(&portfolios).Error; err != nil {
		SendInternalServerError(&c.Controller, "Failed to fetch portfolios")
		return
	}

	SendSuccess(&c.Controller, portfolios)
}

// GetPortfolioByID retrieves a specific portfolio by ID
func (c *PortfolioController) GetPortfolioByID() {
	portfolioID := c.Ctx.Input.Param(":id")
	userID := c.Ctx.Input.GetData("user_id").(string)

	var portfolio models.Portfolio
	if err := models.DB.Where("id = ? AND user_id = ?", portfolioID, userID).First(&portfolio).Error; err != nil {
		SendNotFound(&c.Controller, "Portfolio not found")
		return
	}

	// Fetch related data
	var projects []models.Project
	models.DB.Where("portfolio_id = ?", portfolio.ID).Order("sort_order ASC, created_at DESC").Find(&projects)

	var skills []models.Skill
	models.DB.Where("portfolio_id = ?", portfolio.ID).Order("sort_order ASC, created_at DESC").Find(&skills)

	var testimonials []models.Testimonial
	models.DB.Where("portfolio_id = ?", portfolio.ID).Order("sort_order ASC, created_at DESC").Find(&testimonials)

	var education []models.Education
	models.DB.Where("portfolio_id = ?", portfolio.ID).Order("sort_order ASC, created_at DESC").Find(&education)

	var experience []models.Experience
	models.DB.Where("portfolio_id = ?", portfolio.ID).Order("sort_order ASC, created_at DESC").Find(&experience)

	var publications []models.Publication
	models.DB.Where("portfolio_id = ?", portfolio.ID).Order("sort_order ASC, year DESC, created_at DESC").Find(&publications)

	var contact models.Contact
	models.DB.Where("portfolio_id = ?", portfolio.ID).First(&contact)

	// Return portfolio with related data
	response := map[string]interface{}{
		"id":               portfolio.ID,
		"user_id":          portfolio.UserID,
		"template_id":      portfolio.TemplateID,
		"title":            portfolio.Title,
		"slug":             portfolio.Slug,
		"description":      portfolio.Description,
		"is_published":     portfolio.IsPublished,
		"custom_domain":    portfolio.CustomDomain,
		"config":           portfolio.Config,
		"meta_title":       portfolio.MetaTitle,
		"meta_description": portfolio.MetaDescription,
		"meta_keywords":    portfolio.MetaKeywords,
		"view_count":       portfolio.ViewCount,
		"created_at":       portfolio.CreatedAt,
		"updated_at":       portfolio.UpdatedAt,
		"projects":         projects,
		"skills":           skills,
		"testimonials":     testimonials,
		"education":        education,
		"experience":       experience,
		"publications":     publications,
		"contact":          contact,
	}

	SendSuccess(&c.Controller, response)
}

// UpdatePortfolio updates an existing portfolio
func (c *PortfolioController) UpdatePortfolio() {
	portfolioID := c.Ctx.Input.Param(":id")
	userID := c.Ctx.Input.GetData("user_id").(string)

	var portfolio models.Portfolio
	if err := models.DB.Where("id = ? AND user_id = ?", portfolioID, userID).First(&portfolio).Error; err != nil {
		SendNotFound(&c.Controller, "Portfolio not found")
		return
	}

	var req struct {
		Title          string                 `json:"title"`
		Headline       string                 `json:"headline"`
		Description    string                 `json:"description"`
		Slug           string                 `json:"slug"`
		ProfilePicture string                 `json:"profile_picture"`
		PrimaryColor   string                 `json:"primary_color"`
		SecondaryColor string                 `json:"secondary_color"`
		HeadingFont    string                 `json:"heading_font"`
		BodyFont       string                 `json:"body_font"`
		LayoutStyle    string                 `json:"layout_style"`
		SocialLinks    map[string]interface{} `json:"social_links"`
	}

	if err := c.BindJSON(&req); err != nil {
		SendBadRequestWithMessage(&c.Controller, "Invalid request body")
		return
	}

	// Update config JSON with styling and metadata
	configData := map[string]interface{}{
		"headline":       req.Headline,
		"profilePicture": req.ProfilePicture,
		"primaryColor":   req.PrimaryColor,
		"secondaryColor": req.SecondaryColor,
		"headingFont":    req.HeadingFont,
		"bodyFont":       req.BodyFont,
		"layoutStyle":    req.LayoutStyle,
		"socialLinks":    req.SocialLinks,
	}
	configJSON, _ := json.Marshal(configData)

	// Update fields
	if req.Title != "" {
		portfolio.Title = req.Title
	} else if req.Headline != "" {
		portfolio.Title = req.Headline
	}
	if req.Slug != "" {
		portfolio.Slug = req.Slug
	}
	portfolio.Description = req.Description
	portfolio.Config = string(configJSON)

	if err := models.DB.Save(&portfolio).Error; err != nil {
		SendInternalServerError(&c.Controller, "Failed to update portfolio")
		return
	}

	SendSuccess(&c.Controller, portfolio)
}

// PublishPortfolio publishes a portfolio
func (c *PortfolioController) PublishPortfolio() {
	portfolioID := c.Ctx.Input.Param(":id")
	userID := c.Ctx.Input.GetData("user_id").(string)

	var portfolio models.Portfolio
	if err := models.DB.Where("id = ? AND user_id = ?", portfolioID, userID).First(&portfolio).Error; err != nil {
		SendNotFound(&c.Controller, "Portfolio not found")
		return
	}

	portfolio.IsPublished = true
	if err := models.DB.Save(&portfolio).Error; err != nil {
		SendInternalServerError(&c.Controller, "Failed to publish portfolio")
		return
	}

	SendSuccess(&c.Controller, portfolio)
}

// UnpublishPortfolio unpublishes a portfolio
func (c *PortfolioController) UnpublishPortfolio() {
	portfolioID := c.Ctx.Input.Param(":id")
	userID := c.Ctx.Input.GetData("user_id").(string)

	var portfolio models.Portfolio
	if err := models.DB.Where("id = ? AND user_id = ?", portfolioID, userID).First(&portfolio).Error; err != nil {
		SendNotFound(&c.Controller, "Portfolio not found")
		return
	}

	portfolio.IsPublished = false
	if err := models.DB.Save(&portfolio).Error; err != nil {
		SendInternalServerError(&c.Controller, "Failed to unpublish portfolio")
		return
	}

	SendSuccess(&c.Controller, portfolio)
}

// DeletePortfolio deletes a portfolio
func (c *PortfolioController) DeletePortfolio() {
	portfolioID := c.Ctx.Input.Param(":id")
	userID := c.Ctx.Input.GetData("user_id").(string)

	var portfolio models.Portfolio
	if err := models.DB.Where("id = ? AND user_id = ?", portfolioID, userID).First(&portfolio).Error; err != nil {
		SendNotFound(&c.Controller, "Portfolio not found")
		return
	}

	if err := models.DB.Delete(&portfolio).Error; err != nil {
		SendInternalServerError(&c.Controller, "Failed to delete portfolio")
		return
	}

	SendSuccess(&c.Controller, map[string]string{"message": "Portfolio deleted successfully"})
}

// GetPublicPortfolioBySlug retrieves a published portfolio by slug (public access)
func (c *PortfolioController) GetPublicPortfolioBySlug() {
	slug := c.Ctx.Input.Param(":slug")

	var portfolio models.Portfolio
	if err := models.DB.Where("slug = ? AND is_published = ?", slug, true).First(&portfolio).Error; err != nil {
		SendNotFound(&c.Controller, "Portfolio not found")
		return
	}

	// Fetch related data
	var projects []models.Project
	models.DB.Where("portfolio_id = ?", portfolio.ID).Order("sort_order ASC, created_at DESC").Find(&projects)

	var skills []models.Skill
	models.DB.Where("portfolio_id = ?", portfolio.ID).Order("sort_order ASC, created_at DESC").Find(&skills)

	var testimonials []models.Testimonial
	models.DB.Where("portfolio_id = ?", portfolio.ID).Order("sort_order ASC, created_at DESC").Find(&testimonials)

	var education []models.Education
	models.DB.Where("portfolio_id = ?", portfolio.ID).Order("sort_order ASC, created_at DESC").Find(&education)

	var experience []models.Experience
	models.DB.Where("portfolio_id = ?", portfolio.ID).Order("sort_order ASC, created_at DESC").Find(&experience)

	var publications []models.Publication
	models.DB.Where("portfolio_id = ?", portfolio.ID).Order("sort_order ASC, year DESC, created_at DESC").Find(&publications)

	var contact models.Contact
	models.DB.Where("portfolio_id = ?", portfolio.ID).First(&contact)

	// Track view in analytics
	viewerIP := c.Ctx.Input.IP()
	sessionID := c.Ctx.Input.Header("User-Agent") // Simple session tracking

	analytics := struct {
		ID          string `gorm:"type:char(36);primaryKey"`
		PortfolioID string `gorm:"type:char(36);not null"`
		ViewerIP    string `gorm:"type:varchar(45)"`
		SessionID   string `gorm:"type:varchar(255)"`
	}{
		ID:          uuid.New().String(),
		PortfolioID: portfolio.ID,
		ViewerIP:    viewerIP,
		SessionID:   sessionID,
	}
	models.DB.Table("analytics_views").Create(&analytics)

	// Return portfolio with related data
	response := map[string]interface{}{
		"id":               portfolio.ID,
		"user_id":          portfolio.UserID,
		"template_id":      portfolio.TemplateID,
		"title":            portfolio.Title,
		"slug":             portfolio.Slug,
		"description":      portfolio.Description,
		"is_published":     portfolio.IsPublished,
		"custom_domain":    portfolio.CustomDomain,
		"config":           portfolio.Config,
		"meta_title":       portfolio.MetaTitle,
		"meta_description": portfolio.MetaDescription,
		"meta_keywords":    portfolio.MetaKeywords,
		"view_count":       portfolio.ViewCount,
		"created_at":       portfolio.CreatedAt,
		"updated_at":       portfolio.UpdatedAt,
		"projects":         projects,
		"skills":           skills,
		"testimonials":     testimonials,
		"education":        education,
		"experience":       experience,
		"publications":     publications,
		"contact":          contact,
	}

	SendSuccess(&c.Controller, response)
}

// GetPublicPortfolio retrieves a published portfolio by ID (public access)
func (c *PortfolioController) GetPublicPortfolio() {
	portfolioID := c.Ctx.Input.Param(":id")

	var portfolio models.Portfolio
	if err := models.DB.Where("id = ? AND is_published = ?", portfolioID, true).First(&portfolio).Error; err != nil {
		SendNotFound(&c.Controller, "Portfolio not found")
		return
	}

	// Fetch related data
	var projects []models.Project
	models.DB.Where("portfolio_id = ?", portfolio.ID).Order("sort_order ASC, created_at DESC").Find(&projects)

	var skills []models.Skill
	models.DB.Where("portfolio_id = ?", portfolio.ID).Order("sort_order ASC, created_at DESC").Find(&skills)

	var testimonials []models.Testimonial
	models.DB.Where("portfolio_id = ?", portfolio.ID).Order("sort_order ASC, created_at DESC").Find(&testimonials)

	// Track view in analytics
	viewerIP := c.Ctx.Input.IP()
	sessionID := c.Ctx.Input.Header("User-Agent") // Simple session tracking

	analytics := struct {
		ID          string `gorm:"type:char(36);primaryKey"`
		PortfolioID string `gorm:"type:char(36);not null"`
		ViewerIP    string `gorm:"type:varchar(45)"`
		SessionID   string `gorm:"type:varchar(255)"`
	}{
		ID:          uuid.New().String(),
		PortfolioID: portfolioID,
		ViewerIP:    viewerIP,
		SessionID:   sessionID,
	}
	models.DB.Table("analytics_views").Create(&analytics)

	// Return portfolio with related data
	response := map[string]interface{}{
		"id":               portfolio.ID,
		"user_id":          portfolio.UserID,
		"template_id":      portfolio.TemplateID,
		"title":            portfolio.Title,
		"slug":             portfolio.Slug,
		"description":      portfolio.Description,
		"is_published":     portfolio.IsPublished,
		"custom_domain":    portfolio.CustomDomain,
		"config":           portfolio.Config,
		"meta_title":       portfolio.MetaTitle,
		"meta_description": portfolio.MetaDescription,
		"meta_keywords":    portfolio.MetaKeywords,
		"view_count":       portfolio.ViewCount,
		"created_at":       portfolio.CreatedAt,
		"updated_at":       portfolio.UpdatedAt,
		"projects":         projects,
		"skills":           skills,
		"testimonials":     testimonials,
	}

	SendSuccess(&c.Controller, response)
}
