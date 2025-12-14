package models

import "time"

type Portfolio struct {
	ID              string     `gorm:"type:char(36);primaryKey" json:"id"`
	UserID          string     `gorm:"type:char(36);not null;index" json:"user_id"`
	TemplateID      string     `gorm:"type:char(36);not null;index" json:"template_id"`
	Title           string     `gorm:"type:varchar(255);not null" json:"title"`
	Slug            string     `gorm:"type:varchar(255);uniqueIndex" json:"slug"`
	Description     string     `gorm:"type:text" json:"description"`
	IsPublished     bool       `gorm:"type:tinyint(1);default:0" json:"is_published"`
	CustomDomain    string     `gorm:"type:varchar(255)" json:"custom_domain"`
	Config          string     `gorm:"type:json" json:"config"`
	MetaTitle       string     `gorm:"type:varchar(255)" json:"meta_title"`
	MetaDescription string     `gorm:"type:text" json:"meta_description"`
	MetaKeywords    string     `gorm:"type:varchar(500)" json:"meta_keywords"`
	ViewCount       int64      `gorm:"type:bigint;default:0" json:"view_count"`
	CreatedAt       time.Time  `gorm:"type:datetime(3)" json:"created_at"`
	UpdatedAt       time.Time  `gorm:"type:datetime(3)" json:"updated_at"`
	DeletedAt       *time.Time `gorm:"type:datetime(3);index" json:"deleted_at,omitempty"`
}

func (Portfolio) TableName() string {
	return "portfolios"
}
