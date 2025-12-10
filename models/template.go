package models

import (
	"time"
)

// Template represents the templates table
type Template struct {
	ID            string    `gorm:"type:char(36);primaryKey" json:"id"`
	Name          string    `gorm:"type:varchar(255);not null" json:"name"`
	Description   string    `gorm:"type:text" json:"description"`
	Category      string    `gorm:"type:varchar(50);not null" json:"category"`
	PreviewURL    string    `gorm:"type:varchar(500);column:preview_url" json:"preview_url"`
	ThumbnailURL  string    `gorm:"type:varchar(500);column:thumbnail_url" json:"thumbnail_url"`
	IsActive      bool      `gorm:"type:boolean;default:true;column:is_active" json:"is_active"`
	IsPremium     bool      `gorm:"type:boolean;default:false;column:is_premium" json:"is_premium"`
	Features      string    `gorm:"type:json" json:"features"`
	DefaultConfig string    `gorm:"type:json;column:default_config" json:"default_config"`
	CreatedAt     time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt     time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

// TableName specifies the table name for Template model
func (Template) TableName() string {
	return "templates"
}
