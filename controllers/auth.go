package controllers

import (
	"encoding/json"
	"fmt"

	"Foliogram-BE/models"
	"Foliogram-BE/utils"

	beego "github.com/beego/beego/v2/server/web"
	"golang.org/x/crypto/bcrypt"
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
		fmt.Printf("Login failed: User not found for email %s\n", req.Email)
		SendUnauthorized(&c.Controller, "Invalid email or password")
		return
	}

	fmt.Printf("User found: %s, PasswordHash exists: %v\n", user.Email, user.PasswordHash != "")

	// Compare password
	if err := user.ComparePassword(req.Password); err != nil {
		fmt.Printf("Login failed: Password comparison failed for %s: %v\n", req.Email, err)
		SendUnauthorized(&c.Controller, "Invalid email or password")
		return
	}

	// Generate JWT token
	token, err := utils.GenerateJWT(user.ID, user.Email, user.Role)
	if err != nil {
		SendInternalServerError(&c.Controller, "Failed to generate authentication token")
		return
	}

	// Validate token to get expiry time
	claims, err := utils.ValidateJWT(token)
	if err != nil {
		SendInternalServerError(&c.Controller, "Failed to validate token")
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
		SendUnauthorized(&c.Controller, "Authorization header required")
		return
	}

	// Extract token (format: "Bearer <token>")
	var token string
	if len(authHeader) > 7 && authHeader[:7] == "Bearer " {
		token = authHeader[7:]
	} else {
		SendUnauthorized(&c.Controller, "Invalid authorization header format")
		return
	}

	// Validate token
	claims, err := utils.ValidateJWT(token)
	if err != nil {
		SendUnauthorized(&c.Controller, "Invalid or expired token")
		return
	}

	// Invalidate user session
	if err := utils.InvalidateUserSession(claims.UserID); err != nil {
		SendInternalServerError(&c.Controller, "Failed to logout")
		return
	}

	// Return success response
	SendSuccess(&c.Controller, map[string]interface{}{
		"message": "Logged out successfully",
	})
}

// ChangePasswordRequest represents the change password request body
type ChangePasswordRequest struct {
	CurrentPassword string `json:"current_password"`
	NewPassword     string `json:"new_password"`
}

// ChangeEmailRequest represents the change email request body
type ChangeEmailRequest struct {
	NewEmail        string `json:"new_email"`
	CurrentPassword string `json:"current_password"`
}

// UpdateProfileRequest represents the update profile request body
type UpdateProfileRequest struct {
	FullName        *string `json:"full_name,omitempty"`
	Email           *string `json:"email,omitempty"`
	Mobile          *string `json:"mobile,omitempty"`
	Location        *string `json:"location,omitempty"`
	ProfilePicture  *string `json:"profile_picture,omitempty"`
	CurrentPassword string  `json:"current_password"` // Required for email changes
}

// ChangePassword handles PUT /api/change-password
func (c *AuthController) ChangePassword() {
	var req ChangePasswordRequest

	// Read and parse JSON body
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
		SendUnauthorized(&c.Controller, "Authentication required")
		return
	}

	// Validate required fields
	validationErrors := make(map[string]interface{})

	if req.CurrentPassword == "" {
		validationErrors["current_password"] = "Current password is required"
	}

	if req.NewPassword == "" {
		validationErrors["new_password"] = "New password is required"
	} else if valid, message := utils.ValidatePassword(req.NewPassword); !valid {
		validationErrors["new_password"] = message
	}

	// If there are validation errors, return them
	if len(validationErrors) > 0 {
		SendValidationError(&c.Controller, validationErrors)
		return
	}

	// Get user from database
	var user models.User
	if err := models.DB.Where("id = ?", userID).First(&user).Error; err != nil {
		SendNotFound(&c.Controller, "User not found")
		return
	}

	// Verify current password
	if err := user.ComparePassword(req.CurrentPassword); err != nil {
		SendBadRequest(&c.Controller, map[string]interface{}{
			"current_password": "Current password is incorrect",
		})
		return
	}

	// Check if new password is same as current password
	if err := user.ComparePassword(req.NewPassword); err == nil {
		SendBadRequest(&c.Controller, map[string]interface{}{
			"new_password": "New password must be different from current password",
		})
		return
	}

	// Hash new password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		SendInternalServerError(&c.Controller, "Failed to update password")
		return
	}

	// Update password in database
	user.PasswordHash = string(hashedPassword)
	if err := models.DB.Save(&user).Error; err != nil {
		SendInternalServerError(&c.Controller, "Failed to update password")
		return
	}

	// Return success response
	SendSuccess(&c.Controller, map[string]interface{}{
		"message": "Password changed successfully",
	})
}

// ChangeEmail handles PUT /api/change-email
func (c *AuthController) ChangeEmail() {
	var req ChangeEmailRequest

	// Read and parse JSON body
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
		SendUnauthorized(&c.Controller, "Authentication required")
		return
	}

	// Sanitize input
	req.NewEmail = utils.SanitizeEmail(req.NewEmail)

	// Validate required fields
	validationErrors := make(map[string]interface{})

	if req.NewEmail == "" {
		validationErrors["new_email"] = "New email is required"
	} else if !utils.ValidateEmail(req.NewEmail) {
		validationErrors["new_email"] = "Invalid email format"
	}

	if req.CurrentPassword == "" {
		validationErrors["current_password"] = "Current password is required"
	}

	// If there are validation errors, return them
	if len(validationErrors) > 0 {
		SendValidationError(&c.Controller, validationErrors)
		return
	}

	// Get user from database
	var user models.User
	if err := models.DB.Where("id = ?", userID).First(&user).Error; err != nil {
		SendNotFound(&c.Controller, "User not found")
		return
	}

	// Verify current password
	if err := user.ComparePassword(req.CurrentPassword); err != nil {
		SendBadRequest(&c.Controller, map[string]interface{}{
			"current_password": "Current password is incorrect",
		})
		return
	}

	// Check if new email is same as current email
	if user.Email == req.NewEmail {
		SendBadRequest(&c.Controller, map[string]interface{}{
			"new_email": "New email must be different from current email",
		})
		return
	}

	// Check if email already exists
	var existingUser models.User
	result := models.DB.Where("email = ? AND id != ?", req.NewEmail, userID).First(&existingUser)
	if result.RowsAffected > 0 {
		SendConflict(&c.Controller, "Email already in use")
		return
	}

	// Update email in database
	user.Email = req.NewEmail
	if err := models.DB.Save(&user).Error; err != nil {
		SendInternalServerError(&c.Controller, "Failed to update email")
		return
	}

	// Generate new JWT token with updated email
	token, err := utils.GenerateJWT(user.ID, user.Email, user.Role)
	if err != nil {
		SendInternalServerError(&c.Controller, "Failed to generate authentication token")
		return
	}

	// Sanitize user data before returning
	user.Sanitize()

	// Return success response with new token
	SendSuccess(&c.Controller, map[string]interface{}{
		"message": "Email changed successfully",
		"user":    user,
		"token":   token,
	})
}

// UpdateProfile handles PUT /api/profile
func (c *AuthController) UpdateProfile() {
	var req UpdateProfileRequest

	// Read and parse JSON body
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
		SendUnauthorized(&c.Controller, "Authentication required")
		return
	}

	// Get user from database
	var user models.User
	if err := models.DB.Where("id = ?", userID).First(&user).Error; err != nil {
		SendNotFound(&c.Controller, "User not found")
		return
	}

	validationErrors := make(map[string]interface{})
	passwordRequired := false
	newToken := ""

	// Update full name if provided
	if req.FullName != nil {
		sanitized := utils.SanitizeInput(*req.FullName)
		if sanitized == "" {
			validationErrors["full_name"] = "Full name cannot be empty"
		} else {
			user.FullName = sanitized
		}
	}

	// Update email if provided
	if req.Email != nil {
		sanitized := utils.SanitizeEmail(*req.Email)
		if sanitized == "" {
			validationErrors["email"] = "Email cannot be empty"
		} else if !utils.ValidateEmail(sanitized) {
			validationErrors["email"] = "Invalid email format"
		} else if user.Email != sanitized {
			passwordRequired = true
			// Check if email already exists
			var existingUser models.User
			result := models.DB.Where("email = ? AND id != ?", sanitized, userID).First(&existingUser)
			if result.RowsAffected > 0 {
				validationErrors["email"] = "Email already in use"
			} else {
				user.Email = sanitized
			}
		}
	}

	// Update mobile if provided
	if req.Mobile != nil {
		sanitized := utils.SanitizeInput(*req.Mobile)
		if sanitized != "" && !utils.ValidateMobile(sanitized) {
			validationErrors["mobile"] = "Invalid mobile number format"
		} else {
			user.Mobile = sanitized
		}
	}

	// Update location if provided
	if req.Location != nil {
		user.Location = utils.SanitizeInput(*req.Location)
	}

	// Update profile picture if provided
	if req.ProfilePicture != nil {
		user.ProfilePicture = *req.ProfilePicture
	}

	// Verify password if email is being changed
	if passwordRequired {
		if req.CurrentPassword == "" {
			validationErrors["current_password"] = "Current password is required to change email"
		} else if err := user.ComparePassword(req.CurrentPassword); err != nil {
			validationErrors["current_password"] = "Current password is incorrect"
		}
	}

	// If there are validation errors, return them
	if len(validationErrors) > 0 {
		SendValidationError(&c.Controller, validationErrors)
		return
	}

	// Save updated user to database
	if err := models.DB.Save(&user).Error; err != nil {
		SendInternalServerError(&c.Controller, "Failed to update profile")
		return
	}

	// Generate new JWT token if email was changed
	if passwordRequired && req.Email != nil {
		token, err := utils.GenerateJWT(user.ID, user.Email, user.Role)
		if err != nil {
			SendInternalServerError(&c.Controller, "Failed to generate authentication token")
			return
		}
		newToken = token
	}

	// Sanitize user data before returning
	user.Sanitize()

	// Return success response
	response := map[string]interface{}{
		"message": "Profile updated successfully",
		"user":    user,
	}
	if newToken != "" {
		response["token"] = newToken
	}

	SendSuccess(&c.Controller, response)
}
