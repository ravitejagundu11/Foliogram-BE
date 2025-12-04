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
		SendConflict(&c.Controller, map[string]interface{}{
			"message": "Email already registered",
		})
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
		SendInternalError(&c.Controller, map[string]interface{}{
			"message": "Failed to create user account",
			"error":   err.Error(),
		})
		return
	}

	// Generate JWT token
	token, err := utils.GenerateJWT(user.ID, user.Email, user.Role)
	if err != nil {
		SendInternalError(&c.Controller, map[string]interface{}{
			"message": "Failed to generate authentication token",
			"error":   err.Error(),
		})
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

// LoginRequest represents the login request body
type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Login handles POST /api/login
func (c *AuthController) Login() {
	var req LoginRequest

	// Read and parse JSON body
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &req); err != nil {
		SendBadRequest(&c.Controller, map[string]interface{}{
			"message": "Invalid JSON format",
			"error":   err.Error(),
		})
		return
	}

	// Sanitize input
	req.Email = utils.SanitizeEmail(req.Email)

	// Validate required fields
	validationErrors := make(map[string]interface{})

	if req.Email == "" {
		validationErrors["email"] = "Email is required"
	} else if !utils.ValidateEmail(req.Email) {
		validationErrors["email"] = "Invalid email format"
	}

	if req.Password == "" {
		validationErrors["password"] = "Password is required"
	}

	// If there are validation errors, return them
	if len(validationErrors) > 0 {
		SendValidationError(&c.Controller, validationErrors)
		return
	}

	// Find user by email
	var user models.User
	result := models.DB.Where("email = ?", req.Email).First(&user)
	if result.Error != nil {
		SendUnauthorized(&c.Controller, map[string]interface{}{
			"message": "Invalid email or password",
		})
		return
	}

	// Compare password
	if err := user.ComparePassword(req.Password); err != nil {
		SendUnauthorized(&c.Controller, map[string]interface{}{
			"message": "Invalid email or password",
		})
		return
	}

	// Generate JWT token
	token, err := utils.GenerateJWT(user.ID, user.Email, user.Role)
	if err != nil {
		SendInternalError(&c.Controller, map[string]interface{}{
			"message": "Failed to generate authentication token",
			"error":   err.Error(),
		})
		return
	}

	// Validate token to get expiry time
	claims, err := utils.ValidateJWT(token)
	if err != nil {
		SendInternalError(&c.Controller, map[string]interface{}{
			"message": "Failed to validate token",
			"error":   err.Error(),
		})
		return
	}

	// Store user session in Redis
	if err := utils.StoreUserSession(user.ID, claims.ExpiresAt.Time); err != nil {
		// Log error but don't fail the login
		fmt.Printf("Warning: Failed to store user session: %v\n", err)
	}

	// Sanitize user data before returning
	user.Sanitize()

	// Return success response
	SendSuccess(&c.Controller, map[string]interface{}{
		"user":  user,
		"token": token,
	})
}

// LogoutRequest represents the logout request body
type LogoutRequest struct {
	Token string `json:"token"`
}

// Logout handles POST /api/logout
func (c *AuthController) Logout() {
	// Get token from Authorization header
	authHeader := c.Ctx.Input.Header("Authorization")
	if authHeader == "" {
		SendUnauthorized(&c.Controller, map[string]interface{}{
			"message": "Authorization header required",
		})
		return
	}

	// Extract token (format: "Bearer <token>")
	var token string
	if len(authHeader) > 7 && authHeader[:7] == "Bearer " {
		token = authHeader[7:]
	} else {
		SendUnauthorized(&c.Controller, map[string]interface{}{
			"message": "Invalid authorization header format",
		})
		return
	}

	// Validate token
	claims, err := utils.ValidateJWT(token)
	if err != nil {
		SendUnauthorized(&c.Controller, map[string]interface{}{
			"message": "Invalid or expired token",
		})
		return
	}

	// Invalidate user session
	if err := utils.InvalidateUserSession(claims.UserID); err != nil {
		SendInternalError(&c.Controller, map[string]interface{}{
			"message": "Failed to logout",
			"error":   err.Error(),
		})
		return
	}

	// Return success response
	SendSuccess(&c.Controller, map[string]interface{}{
		"message": "Logged out successfully",
	})
}
