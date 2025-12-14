package models

import "time"

type Education struct {
	ID          string     `gorm:"type:char(36);primaryKey" json:"id"`
	PortfolioID string     `gorm:"type:char(36);not null;index" json:"portfolio_id"`
	Institution string     `gorm:"type:varchar(255);not null" json:"institution"`
	Degree      string     `gorm:"type:varchar(255)" json:"degree"`
	Field       string     `gorm:"type:varchar(255)" json:"field"`
	StartDate   string     `gorm:"type:varchar(50)" json:"start_date"`
	EndDate     string     `gorm:"type:varchar(50)" json:"end_date"`
	Grade       string     `gorm:"type:varchar(50)" json:"grade"`
	Description string     `gorm:"type:text" json:"description"`
	SortOrder   int        `gorm:"type:int;default:0" json:"sort_order"`
	CreatedAt   time.Time  `gorm:"type:datetime(3)" json:"created_at"`
	UpdatedAt   time.Time  `gorm:"type:datetime(3)" json:"updated_at"`
	DeletedAt   *time.Time `gorm:"type:datetime(3);index" json:"deleted_at,omitempty"`
}

func (Education) TableName() string {
	return "education"
}
