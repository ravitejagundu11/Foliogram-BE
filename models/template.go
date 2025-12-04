package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Template represents a portfolio template
type Template struct {
	ID            string         `gorm:"type:char(36);primaryKey" json:"id"`
	Name          string         `gorm:"type:varchar(255);not null" json:"name"`
	Description   string         `gorm:"type:text" json:"description"`
	Category      string         `gorm:"type:varchar(100)" json:"category"` // e.g., "developer", "designer", "photographer"
	PreviewURL    string         `gorm:"type:varchar(500)" json:"preview_url"`
	ThumbnailURL  string         `gorm:"type:varchar(500)" json:"thumbnail_url"`
	IsActive      bool           `gorm:"default:true" json:"is_active"`
	IsPremium     bool           `gorm:"default:false" json:"is_premium"`
	Features      string         `gorm:"type:json" json:"features"`       // JSON array of features
	DefaultConfig string         `gorm:"type:json" json:"default_config"` // JSON object with default colors, fonts, etc.
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"-"`
}

// BeforeCreate hook to generate UUID
func (t *Template) BeforeCreate(tx *gorm.DB) error {
	if t.ID == "" {
		t.ID = uuid.New().String()
	}
	return nil
}

// TableName specifies the table name
func (Template) TableName() string {
	return "templates"
}
