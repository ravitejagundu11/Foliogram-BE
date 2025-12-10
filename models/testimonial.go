package models

import "time"

type Testimonial struct {
	ID          string     `gorm:"type:char(36);primaryKey" json:"id"`
	PortfolioID string     `gorm:"type:char(36);not null;index" json:"portfolio_id"`
	Name        string     `gorm:"type:varchar(255);not null" json:"name"`
	Position    string     `gorm:"type:varchar(255)" json:"position"`
	Company     string     `gorm:"type:varchar(255)" json:"company"`
	Content     string     `gorm:"type:text;not null" json:"content"`
	Rating      int64      `gorm:"type:bigint" json:"rating"`
	AvatarURL   string     `gorm:"type:varchar(500)" json:"avatar_url"`
	LinkedInURL string     `gorm:"type:varchar(500)" json:"linked_in_url"`
	IsVisible   bool       `gorm:"type:tinyint(1);default:1" json:"is_visible"`
	SortOrder   int64      `gorm:"type:bigint;default:0" json:"sort_order"`
	CreatedAt   time.Time  `gorm:"type:datetime(3);autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time  `gorm:"type:datetime(3);autoUpdateTime" json:"updated_at"`
	DeletedAt   *time.Time `gorm:"type:datetime(3);index" json:"deleted_at,omitempty"`
}

func (Testimonial) TableName() string {
	return "testimonials"
}
