package models

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// User represents the users table
type User struct {
	ID             string    `gorm:"type:char(36);primaryKey" json:"-"`
	FullName       string    `gorm:"type:varchar(255);not null" json:"full_name" valid:"required~Full name is required,stringlength(2|255)~Full name must be between 2 and 255 characters"`
	Email          string    `gorm:"type:varchar(255);uniqueIndex;not null" json:"email" valid:"required~Email is required,email~Invalid email format"`
	PasswordHash   string    `gorm:"type:varchar(255);not null" json:"-"`
	Password       string    `gorm:"-" json:"password,omitempty" valid:"required~Password is required,stringlength(8|100)~Password must be at least 8 characters"`
	Mobile         string    `gorm:"type:varchar(20)" json:"-" valid:"optional,matches(^[+]?[0-9]{10\\,15}$)~Invalid mobile number format"`
	Location       string    `gorm:"type:varchar(255)" json:"-"`
	ProfilePicture string    `gorm:"type:text" json:"profile_picture"`
	Role           string    `gorm:"type:varchar(20);not null;default:'user'" json:"role"`
	CreatedAt      time.Time `gorm:"type:datetime(3)" json:"-"`
	UpdatedAt      time.Time `gorm:"type:datetime(3)" json:"-"`
}

// TableName specifies the table name for User model
func (User) TableName() string {
	return "users"
}

// BeforeCreate hook to generate UUID and hash password
func (u *User) BeforeCreate(tx *gorm.DB) error {
	// Generate UUID if not set
	if u.ID == "" {
		u.ID = generateUUID()
	}

	// Set default role if not specified
	if u.Role == "" {
		u.Role = "user"
	}

	// Hash password if provided
	if u.Password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
		if err != nil {
			return err
		}
		u.PasswordHash = string(hashedPassword)
	}

	return nil
}

// ComparePassword compares the provided password with the hashed password
func (u *User) ComparePassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(u.PasswordHash), []byte(password))
}

// Sanitize removes sensitive information before returning
func (u *User) Sanitize() {
	u.PasswordHash = ""
	u.Password = ""
}

// Validate checks if user data is valid
func (u *User) Validate() error {
	if u.Email == "" {
		return errors.New("email is required")
	}
	if u.FullName == "" {
		return errors.New("full name is required")
	}
	if u.Password == "" && u.PasswordHash == "" {
		return errors.New("password is required")
	}
	if len(u.Password) > 0 && len(u.Password) < 8 {
		return errors.New("password must be at least 8 characters")
	}
	return nil
}

// generateUUID generates a UUID v4
func generateUUID() string {
	return uuid.New().String()
}
