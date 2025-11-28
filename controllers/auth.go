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
}

// RegisterResponse represents the registration response
type RegisterResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Token   string      `json:"token,omitempty"`
}

// Register handles POST /api/register
func (c *AuthController) Register() {
	var req RegisterRequest
	fmt.Print(string(c.Ctx.Input.RequestBody) + "request body\n")

	// Read and parse JSON body
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &req); err != nil {
		c.Data["json"] = RegisterResponse{
			Success: false,
			Message: "Invalid JSON format: " + err.Error(),
		}
		c.Ctx.ResponseWriter.WriteHeader(400)
		c.ServeJSON()
		return
	}

	fmt.Printf("Received request: %+v\n", req)
	// Sanitize input
	req.FullName = utils.SanitizeInput(req.FullName)
	req.Email = utils.SanitizeEmail(req.Email)

	// Validate required fields
	if req.FullName == "" {
		c.Data["json"] = RegisterResponse{
			Success: false,
			Message: "Full name is required",
		}
		c.Ctx.ResponseWriter.WriteHeader(400)
		c.ServeJSON()
		return
	}

	if req.Email == "" {
		c.Data["json"] = RegisterResponse{
			Success: false,
			Message: "Email is required",
		}
		c.Ctx.ResponseWriter.WriteHeader(400)
		c.ServeJSON()
		return
	}

	// Validate email format
	if !utils.ValidateEmail(req.Email) {
		c.Data["json"] = RegisterResponse{
			Success: false,
			Message: "Invalid email format",
		}
		c.Ctx.ResponseWriter.WriteHeader(400)
		c.ServeJSON()
		return
	}

	// Validate password
	if valid, message := utils.ValidatePassword(req.Password); !valid {
		c.Data["json"] = RegisterResponse{
			Success: false,
			Message: message,
		}
		c.Ctx.ResponseWriter.WriteHeader(400)
		c.ServeJSON()
		return
	}

	// Check if email already exists
	var existingUser models.User
	result := models.DB.Where("email = ?", req.Email).First(&existingUser)
	if result.RowsAffected > 0 {
		c.Data["json"] = RegisterResponse{
			Success: false,
			Message: "Email already registered",
		}
		c.Ctx.ResponseWriter.WriteHeader(409)
		c.ServeJSON()
		return
	}

	// Create new user
	user := models.User{
		FullName: req.FullName,
		Email:    req.Email,
		Password: req.Password, // Will be hashed in BeforeCreate hook
		Role:     "user",
	}

	// Save user to database
	if err := models.DB.Create(&user).Error; err != nil {
		c.Data["json"] = RegisterResponse{
			Success: false,
			Message: "Failed to create user account",
		}
		c.Ctx.ResponseWriter.WriteHeader(500)
		c.ServeJSON()
		return
	}

	// Generate JWT token
	token, err := utils.GenerateJWT(user.ID, user.Email, user.Role)
	if err != nil {
		c.Data["json"] = RegisterResponse{
			Success: false,
			Message: "Failed to generate authentication token",
		}
		c.Ctx.ResponseWriter.WriteHeader(500)
		c.ServeJSON()
		return
	}

	// Sanitize user data before returning
	user.Sanitize()

	// Return success response
	c.Data["json"] = RegisterResponse{
		Success: true,
		Message: "User registered successfully",
		Token:   token,
		Data: map[string]interface{}{
			"user": user,
		},
	}
	c.Ctx.ResponseWriter.WriteHeader(201)
	c.ServeJSON()
}
