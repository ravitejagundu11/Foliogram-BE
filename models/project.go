package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Project represents a project in a portfolio
type Project struct {
	ID           string     `gorm:"type:char(36);primaryKey" json:"id"`
	PortfolioID  string     `gorm:"type:char(36);not null;index" json:"portfolio_id"`
	Title        string     `gorm:"type:varchar(255);not null" json:"title"`
	Description  string     `gorm:"type:text" json:"description"`
	ImageURL     string     `gorm:"type:varchar(500)" json:"image_url"`
	GalleryURLs  string     `gorm:"type:json" json:"gallery_urls"` // JSON array of image URLs
	DemoURL      string     `gorm:"type:varchar(500)" json:"demo_url"`
	GithubURL    string     `gorm:"type:varchar(500)" json:"github_url"`
	Tags         string     `gorm:"type:json" json:"tags"`         // JSON array of tags
	Technologies string     `gorm:"type:json" json:"technologies"` // JSON array of tech stack
	StartDate    *time.Time `json:"start_date"`
	EndDate      *time.Time `json:"end_date"`
	IsFeatured   bool       `gorm:"default:false" json:"is_featured"`
	SortOrder    int        `gorm:"default:0" json:"sort_order"`

	// Association
	Portfolio Portfolio `gorm:"foreignKey:PortfolioID" json:"-"`

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

// BeforeCreate hook to generate UUID
func (p *Project) BeforeCreate(tx *gorm.DB) error {
	if p.ID == "" {
		p.ID = uuid.New().String()
	}
	return nil
}

// TableName specifies the table name
func (Project) TableName() string {
	return "projects"
}
