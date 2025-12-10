package models

import "time"

type Project struct {
	ID           string     `gorm:"type:char(36);primaryKey" json:"id"`
	PortfolioID  string     `gorm:"type:char(36);not null;index" json:"portfolio_id"`
	Title        string     `gorm:"type:varchar(255);not null" json:"title"`
	Description  string     `gorm:"type:text" json:"description"`
	ImageURL     string     `gorm:"type:varchar(500)" json:"image_url"`
	GalleryURLs  string     `gorm:"type:json" json:"gallery_urls"` // JSON array
	DemoURL      string     `gorm:"type:varchar(500)" json:"demo_url"`
	GithubURL    string     `gorm:"type:varchar(500)" json:"github_url"`
	Tags         string     `gorm:"type:json" json:"tags"`         // JSON array
	Technologies string     `gorm:"type:json" json:"technologies"` // JSON array
	StartDate    *time.Time `gorm:"type:datetime(3)" json:"start_date"`
	EndDate      *time.Time `gorm:"type:datetime(3)" json:"end_date"`
	IsFeatured   bool       `gorm:"type:tinyint(1);default:0" json:"is_featured"`
	SortOrder    int64      `gorm:"type:bigint;default:0" json:"sort_order"`
	CreatedAt    time.Time  `gorm:"type:datetime(3);autoCreateTime" json:"created_at"`
	UpdatedAt    time.Time  `gorm:"type:datetime(3);autoUpdateTime" json:"updated_at"`
	DeletedAt    *time.Time `gorm:"type:datetime(3);index" json:"deleted_at,omitempty"`
}

func (Project) TableName() string {
	return "projects"
}
