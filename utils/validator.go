package utils

import (
	"regexp"
	"strings"
)

// SanitizeInput removes potentially dangerous characters from input
func SanitizeInput(input string) string {
	// Trim whitespace
	input = strings.TrimSpace(input)
	
	// Remove null bytes
	input = strings.ReplaceAll(input, "\x00", "")
	
	return input
}

// ValidateEmail validates email format
func ValidateEmail(email string) bool {
	email = strings.TrimSpace(strings.ToLower(email))
	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,}$`)
	return emailRegex.MatchString(email)
}

// ValidatePassword checks password strength
func ValidatePassword(password string) (bool, string) {
	if len(password) < 8 {
		return false, "Password must be at least 8 characters long"
	}
	if len(password) > 100 {
		return false, "Password must not exceed 100 characters"
	}
	
	hasUpper := regexp.MustCompile(`[A-Z]`).MatchString(password)
	hasLower := regexp.MustCompile(`[a-z]`).MatchString(password)
	hasNumber := regexp.MustCompile(`[0-9]`).MatchString(password)
	
	if !hasUpper || !hasLower || !hasNumber {
		return false, "Password must contain at least one uppercase letter, one lowercase letter, and one number"
	}
	
	return true, ""
}

// ValidateMobile validates mobile number format
func ValidateMobile(mobile string) bool {
	if mobile == "" {
		return true // Optional field
	}
	mobileRegex := regexp.MustCompile(`^[+]?[0-9]{10,15}$`)
	return mobileRegex.MatchString(mobile)
}

// SanitizeEmail normalizes email format
func SanitizeEmail(email string) string {
	return strings.TrimSpace(strings.ToLower(email))
}
