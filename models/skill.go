package models

import "time"

type Skill struct {
	ID          string     `gorm:"type:char(36);primaryKey" json:"id"`
	PortfolioID string     `gorm:"type:char(36);not null;index" json:"portfolio_id"`
	Name        string     `gorm:"type:varchar(255);not null" json:"name"`
	Category    string     `gorm:"type:varchar(100)" json:"category"`
	Level       string     `gorm:"type:varchar(50)" json:"level"`
	Percentage  int64      `gorm:"type:bigint" json:"percentage"`
	IconURL     string     `gorm:"type:varchar(500)" json:"icon_url"`
	SortOrder   int64      `gorm:"type:bigint;default:0" json:"sort_order"`
	CreatedAt   time.Time  `gorm:"type:datetime(3);autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time  `gorm:"type:datetime(3);autoUpdateTime" json:"updated_at"`
	DeletedAt   *time.Time `gorm:"type:datetime(3);index" json:"deleted_at,omitempty"`
}

func (Skill) TableName() string {
	return "skills"
}
