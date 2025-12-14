package controllers

import (
	"Foliogram-BE/models"
	"encoding/json"

	"github.com/beego/beego/v2/server/web"
	"github.com/google/uuid"
)

type ProjectController struct {
	web.Controller
}

// CreateProject creates a new project for a portfolio
func (c *ProjectController) CreateProject() {
	userID := c.Ctx.Input.GetData("user_id").(string)
	portfolioID := c.Ctx.Input.Param(":portfolioId")

	// Verify portfolio belongs to user
	var portfolio models.Portfolio
	if err := models.DB.Where("id = ? AND user_id = ?", portfolioID, userID).First(&portfolio).Error; err != nil {
		SendNotFound(&c.Controller, "Portfolio not found")
		return
	}

	var req struct {
		Title       string   `json:"title"`
		Description string   `json:"description"`
		Images      []string `json:"images"`
		TechStack   []string `json:"techStack"`
		DemoURL     string   `json:"demoUrl"`
		CodeURL     string   `json:"codeUrl"`
		Featured    bool     `json:"featured"`
		Order       int      `json:"order"`
	}

	if err := c.BindJSON(&req); err != nil {
		SendBadRequestWithMessage(&c.Controller, "Invalid request body")
		return
	}

	// Marshal arrays to JSON (always provide valid JSON, even if empty)
	galleryJSON, _ := json.Marshal(req.Images)
	techStackJSON, _ := json.Marshal(req.TechStack)

	// Empty array for tags (always valid JSON)
	emptyArrayJSON := "[]"

	// Use first image as main image
	imageURL := ""
	if len(req.Images) > 0 {
		imageURL = req.Images[0]
	}

	project := models.Project{
		ID:           uuid.New().String(),
		PortfolioID:  portfolioID,
		Title:        req.Title,
		Description:  req.Description,
		ImageURL:     imageURL,
		GalleryURLs:  string(galleryJSON),
		DemoURL:      req.DemoURL,
		GithubURL:    req.CodeURL,
		Tags:         emptyArrayJSON,
		Technologies: string(techStackJSON),
		IsFeatured:   req.Featured,
		SortOrder:    int64(req.Order),
	}

	if err := models.DB.Create(&project).Error; err != nil {
		SendInternalServerError(&c.Controller, "Failed to create project")
		return
	}

	SendCreated(&c.Controller, project)
}

// GetProjectsByPortfolio retrieves all projects for a portfolio
func (c *ProjectController) GetProjectsByPortfolio() {
	userID := c.Ctx.Input.GetData("user_id").(string)
	portfolioID := c.Ctx.Input.Param(":portfolioId")

	// Verify portfolio belongs to user
	var portfolio models.Portfolio
	if err := models.DB.Where("id = ? AND user_id = ?", portfolioID, userID).First(&portfolio).Error; err != nil {
		SendNotFound(&c.Controller, "Portfolio not found")
		return
	}

	var projects []models.Project
	if err := models.DB.Where("portfolio_id = ?", portfolioID).Order("sort_order ASC, created_at DESC").Find(&projects).Error; err != nil {
		SendInternalServerError(&c.Controller, "Failed to fetch projects")
		return
	}

	SendSuccess(&c.Controller, projects)
}

// GetProjectByID retrieves a specific project
func (c *ProjectController) GetProjectByID() {
	userID := c.Ctx.Input.GetData("user_id").(string)
	portfolioID := c.Ctx.Input.Param(":portfolioId")
	projectID := c.Ctx.Input.Param(":projectId")

	// Verify portfolio belongs to user
	var portfolio models.Portfolio
	if err := models.DB.Where("id = ? AND user_id = ?", portfolioID, userID).First(&portfolio).Error; err != nil {
		SendNotFound(&c.Controller, "Portfolio not found")
		return
	}

	var project models.Project
	if err := models.DB.Where("id = ? AND portfolio_id = ?", projectID, portfolioID).First(&project).Error; err != nil {
		SendNotFound(&c.Controller, "Project not found")
		return
	}

	SendSuccess(&c.Controller, project)
}

// UpdateProject updates an existing project
func (c *ProjectController) UpdateProject() {
	userID := c.Ctx.Input.GetData("user_id").(string)
	portfolioID := c.Ctx.Input.Param(":portfolioId")
	projectID := c.Ctx.Input.Param(":projectId")

	// Verify portfolio belongs to user
	var portfolio models.Portfolio
	if err := models.DB.Where("id = ? AND user_id = ?", portfolioID, userID).First(&portfolio).Error; err != nil {
		SendNotFound(&c.Controller, "Portfolio not found")
		return
	}

	var project models.Project
	if err := models.DB.Where("id = ? AND portfolio_id = ?", projectID, portfolioID).First(&project).Error; err != nil {
		SendNotFound(&c.Controller, "Project not found")
		return
	}

	var req struct {
		Title       string   `json:"title"`
		Description string   `json:"description"`
		Images      []string `json:"images"`
		TechStack   []string `json:"techStack"`
		DemoURL     string   `json:"demoUrl"`
		CodeURL     string   `json:"codeUrl"`
		Featured    bool     `json:"featured"`
		Order       int      `json:"order"`
	}

	if err := c.BindJSON(&req); err != nil {
		SendBadRequestWithMessage(&c.Controller, "Invalid request body")
		return
	}

	// Update fields
	if req.Title != "" {
		project.Title = req.Title
	}
	project.Description = req.Description

	if req.Images != nil {
		galleryJSON, _ := json.Marshal(req.Images)
		project.GalleryURLs = string(galleryJSON)
		if len(req.Images) > 0 {
			project.ImageURL = req.Images[0]
		}
	}

	if req.TechStack != nil {
		techStackJSON, _ := json.Marshal(req.TechStack)
		project.Technologies = string(techStackJSON)
	}

	project.DemoURL = req.DemoURL
	project.GithubURL = req.CodeURL
	project.IsFeatured = req.Featured
	project.SortOrder = int64(req.Order)

	if err := models.DB.Save(&project).Error; err != nil {
		SendInternalServerError(&c.Controller, "Failed to update project")
		return
	}

	SendSuccess(&c.Controller, project)
}

// DeleteProject deletes a project
func (c *ProjectController) DeleteProject() {
	userID := c.Ctx.Input.GetData("user_id").(string)
	portfolioID := c.Ctx.Input.Param(":portfolioId")
	projectID := c.Ctx.Input.Param(":projectId")

	// Verify portfolio belongs to user
	var portfolio models.Portfolio
	if err := models.DB.Where("id = ? AND user_id = ?", portfolioID, userID).First(&portfolio).Error; err != nil {
		SendNotFound(&c.Controller, "Portfolio not found")
		return
	}

	var project models.Project
	if err := models.DB.Where("id = ? AND portfolio_id = ?", projectID, portfolioID).First(&project).Error; err != nil {
		SendNotFound(&c.Controller, "Project not found")
		return
	}

	if err := models.DB.Delete(&project).Error; err != nil {
		SendInternalServerError(&c.Controller, "Failed to delete project")
		return
	}

	SendSuccess(&c.Controller, map[string]string{"message": "Project deleted successfully"})
}
