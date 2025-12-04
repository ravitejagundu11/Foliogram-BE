package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Portfolio represents a user's portfolio
type Portfolio struct {
	ID           string `gorm:"type:char(36);primaryKey" json:"id"`
	UserID       string `gorm:"type:char(36);not null;index" json:"user_id"`
	TemplateID   string `gorm:"type:char(36);not null;index" json:"template_id"`
	Title        string `gorm:"type:varchar(255);not null" json:"title"`
	Slug         string `gorm:"type:varchar(255);uniqueIndex;not null" json:"slug"`
	Description  string `gorm:"type:text" json:"description"`
	IsPublished  bool   `gorm:"default:false" json:"is_published"`
	CustomDomain string `gorm:"type:varchar(255)" json:"custom_domain"`

	// Customization settings (stored as JSON)
	Config string `gorm:"type:json" json:"config"` // Colors, fonts, layout preferences

	// SEO fields
	MetaTitle       string `gorm:"type:varchar(255)" json:"meta_title"`
	MetaDescription string `gorm:"type:text" json:"meta_description"`
	MetaKeywords    string `gorm:"type:varchar(500)" json:"meta_keywords"`

	// Analytics
	ViewCount int `gorm:"default:0" json:"view_count"`

	// Associations
	User         User          `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"user,omitempty"`
	Template     Template      `gorm:"foreignKey:TemplateID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT" json:"template,omitempty"`
	Projects     []Project     `gorm:"foreignKey:PortfolioID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"projects,omitempty"`
	Skills       []Skill       `gorm:"foreignKey:PortfolioID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"skills,omitempty"`
	Testimonials []Testimonial `gorm:"foreignKey:PortfolioID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"testimonials,omitempty"`

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

// BeforeCreate hook to generate UUID
func (p *Portfolio) BeforeCreate(tx *gorm.DB) error {
	if p.ID == "" {
		p.ID = uuid.New().String()
	}
	return nil
}

// TableName specifies the table name
func (Portfolio) TableName() string {
	return "portfolios"
}
