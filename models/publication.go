package models

import "time"

type Publication struct {
	ID          string     `gorm:"type:char(36);primaryKey" json:"id"`
	PortfolioID string     `gorm:"type:char(36);not null;index" json:"portfolio_id"`
	Title       string     `gorm:"type:varchar(500);not null" json:"title"`
	Authors     string     `gorm:"type:text" json:"authors"`
	Journal     string     `gorm:"type:varchar(255)" json:"journal"`
	Conference  string     `gorm:"type:varchar(255)" json:"conference"`
	Year        string     `gorm:"type:varchar(10)" json:"year"`
	DOI         string     `gorm:"type:varchar(255)" json:"doi"`
	URL         string     `gorm:"type:varchar(500)" json:"url"`
	Abstract    string     `gorm:"type:text" json:"abstract"`
	Type        string     `gorm:"type:varchar(50)" json:"type"`
	SortOrder   int        `gorm:"type:int;default:0" json:"sort_order"`
	CreatedAt   time.Time  `gorm:"type:datetime(3)" json:"created_at"`
	UpdatedAt   time.Time  `gorm:"type:datetime(3)" json:"updated_at"`
	DeletedAt   *time.Time `gorm:"type:datetime(3);index" json:"deleted_at,omitempty"`
}

func (Publication) TableName() string {
	return "publications"
}
