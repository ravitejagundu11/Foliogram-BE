package controllers

import (
	"encoding/json"
	"fmt"

	"Foliogram-BE/models"
	"Foliogram-BE/utils"

	beego "github.com/beego/beego/v2/server/web"
)

// AuthController handles authentication operations
type AuthController struct {
	beego.Controller
}

// RegisterRequest represents the registration request body
type RegisterRequest struct {
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Mobile   string `json:"mobile,omitempty"`
	Location string `json:"location,omitempty"`
}

// Register handles POST /api/register
func (c *AuthController) Register() {
	var req RegisterRequest
	fmt.Print(string(c.Ctx.Input.RequestBody) + "request body\n")

	// Read and parse JSON body
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &req); err != nil {
		SendBadRequest(&c.Controller, map[string]interface{}{
			"message": "Invalid JSON format",
			"error":   err.Error(),
		})
		return
	}

	fmt.Printf("Received request: %+v\n", req)

	// Sanitize input
	req.FullName = utils.SanitizeInput(req.FullName)
	req.Email = utils.SanitizeEmail(req.Email)
	req.Mobile = utils.SanitizeInput(req.Mobile)
	req.Location = utils.SanitizeInput(req.Location)

	// Validate required fields
	validationErrors := make(map[string]interface{})

	if req.FullName == "" {
		validationErrors["full_name"] = "Full name is required"
	}

	if req.Email == "" {
		validationErrors["email"] = "Email is required"
	} else if !utils.ValidateEmail(req.Email) {
		validationErrors["email"] = "Invalid email format"
	}

	if req.Password == "" {
		validationErrors["password"] = "Password is required"
	} else if valid, message := utils.ValidatePassword(req.Password); !valid {
		validationErrors["password"] = message
	}

	if req.Mobile != "" && !utils.ValidateMobile(req.Mobile) {
		validationErrors["mobile"] = "Invalid mobile number format"
	}

	// If there are validation errors, return them
	if len(validationErrors) > 0 {
		SendValidationError(&c.Controller, validationErrors)
		return
	}

	// Check if email already exists
	var existingUser models.User
	result := models.DB.Where("email = ?", req.Email).First(&existingUser)
	if result.RowsAffected > 0 {
		SendConflict(&c.Controller, "Email already registered")
		return
	}

	// Create new user
	user := models.User{
		FullName: req.FullName,
		Email:    req.Email,
		Password: req.Password, // Will be hashed in BeforeCreate hook
		Mobile:   req.Mobile,
		Location: req.Location,
		Role:     "user",
	}

	// Save user to database
	if err := models.DB.Create(&user).Error; err != nil {
		SendInternalServerError(&c.Controller, "Failed to create user account")
		return
	}

	// Generate JWT token
	token, err := utils.GenerateJWT(user.ID, user.Email, user.Role)
	if err != nil {
		SendInternalServerError(&c.Controller, "Failed to generate authentication token")
		return
	}

	// Sanitize user data before returning
	user.Sanitize()

	// Return success response
	SendCreated(&c.Controller, map[string]interface{}{
		"user":  user,
		"token": token,
	})
}
