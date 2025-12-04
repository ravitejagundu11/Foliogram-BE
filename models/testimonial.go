package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Testimonial represents a testimonial in a portfolio
type Testimonial struct {
	ID          string `gorm:"type:char(36);primaryKey" json:"id"`
	PortfolioID string `gorm:"type:char(36);not null;index" json:"portfolio_id"`
	Name        string `gorm:"type:varchar(255);not null" json:"name"`
	Position    string `gorm:"type:varchar(255)" json:"position"`
	Company     string `gorm:"type:varchar(255)" json:"company"`
	Content     string `gorm:"type:text;not null" json:"content"`
	Rating      int    `gorm:"type:int" json:"rating"` // 1-5 stars
	AvatarURL   string `gorm:"type:varchar(500)" json:"avatar_url"`
	LinkedInURL string `gorm:"type:varchar(500)" json:"linkedin_url"`
	IsVisible   bool   `gorm:"default:true" json:"is_visible"`
	SortOrder   int    `gorm:"default:0" json:"sort_order"`

	// Association
	Portfolio Portfolio `gorm:"foreignKey:PortfolioID" json:"-"`

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

// BeforeCreate hook to generate UUID
func (t *Testimonial) BeforeCreate(tx *gorm.DB) error {
	if t.ID == "" {
		t.ID = uuid.New().String()
	}
	return nil
}

// TableName specifies the table name
func (Testimonial) TableName() string {
	return "testimonials"
}
