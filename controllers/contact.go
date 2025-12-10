package controllers

import (
	"Foliogram-BE/models"

	beego "github.com/beego/beego/v2/server/web"
	"github.com/google/uuid"
)

// ContactController handles contact CRUD operations
type ContactController struct {
	beego.Controller
}

// CreateOrUpdateContact creates or updates contact info for a portfolio
func (c *ContactController) CreateOrUpdateContact() {
	portfolioID := c.Ctx.Input.Param(":portfolioId")
	userID := c.Ctx.Input.GetData("user_id").(string)

	// Verify portfolio belongs to user
	var portfolio models.Portfolio
	if err := models.DB.Where("id = ? AND user_id = ?", portfolioID, userID).First(&portfolio).Error; err != nil {
		SendNotFound(&c.Controller, "Portfolio not found")
		return
	}

	var req struct {
		Email   string `json:"email"`
		Phone   string `json:"phone"`
		Address string `json:"address"`
		City    string `json:"city"`
		Country string `json:"country"`
		MapURL  string `json:"map_url"`
	}

	if err := c.BindJSON(&req); err != nil {
		SendBadRequestWithMessage(&c.Controller, "Invalid request body")
		return
	}

	// Check if contact already exists for this portfolio
	var contact models.Contact
	result := models.DB.Where("portfolio_id = ?", portfolioID).First(&contact)

	if result.Error != nil {
		// Create new contact
		contact = models.Contact{
			ID:          uuid.New().String(),
			PortfolioID: portfolioID,
			Email:       req.Email,
			Phone:       req.Phone,
			Address:     req.Address,
			City:        req.City,
			Country:     req.Country,
			MapURL:      req.MapURL,
		}

		if err := models.DB.Create(&contact).Error; err != nil {
			SendInternalServerError(&c.Controller, "Failed to create contact info")
			return
		}

		SendCreated(&c.Controller, contact)
	} else {
		// Update existing contact
		contact.Email = req.Email
		contact.Phone = req.Phone
		contact.Address = req.Address
		contact.City = req.City
		contact.Country = req.Country
		contact.MapURL = req.MapURL

		if err := models.DB.Save(&contact).Error; err != nil {
			SendInternalServerError(&c.Controller, "Failed to update contact info")
			return
		}

		SendSuccess(&c.Controller, contact)
	}
}

// GetContactByPortfolio retrieves contact info for a portfolio
func (c *ContactController) GetContactByPortfolio() {
	portfolioID := c.Ctx.Input.Param(":portfolioId")
	userID := c.Ctx.Input.GetData("user_id").(string)

	// Verify portfolio belongs to user
	var portfolio models.Portfolio
	if err := models.DB.Where("id = ? AND user_id = ?", portfolioID, userID).First(&portfolio).Error; err != nil {
		SendNotFound(&c.Controller, "Portfolio not found")
		return
	}

	var contact models.Contact
	if err := models.DB.Where("portfolio_id = ?", portfolioID).First(&contact).Error; err != nil {
		// Return empty contact instead of error
		SendSuccess(&c.Controller, models.Contact{PortfolioID: portfolioID})
		return
	}

	SendSuccess(&c.Controller, contact)
}

// DeleteContact deletes contact info for a portfolio
func (c *ContactController) DeleteContact() {
	portfolioID := c.Ctx.Input.Param(":portfolioId")
	userID := c.Ctx.Input.GetData("user_id").(string)

	// Verify portfolio belongs to user
	var portfolio models.Portfolio
	if err := models.DB.Where("id = ? AND user_id = ?", portfolioID, userID).First(&portfolio).Error; err != nil {
		SendNotFound(&c.Controller, "Portfolio not found")
		return
	}

	if err := models.DB.Where("portfolio_id = ?", portfolioID).Delete(&models.Contact{}).Error; err != nil {
		SendInternalServerError(&c.Controller, "Failed to delete contact info")
		return
	}

	SendSuccess(&c.Controller, map[string]string{"message": "Contact info deleted successfully"})
}
