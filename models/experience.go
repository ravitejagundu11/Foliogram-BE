package models

import "time"

type Experience struct {
	ID          string     `gorm:"type:char(36);primaryKey" json:"id"`
	PortfolioID string     `gorm:"type:char(36);not null;index" json:"portfolio_id"`
	Company     string     `gorm:"type:varchar(255);not null" json:"company"`
	Position    string     `gorm:"type:varchar(255)" json:"position"`
	Location    string     `gorm:"type:varchar(255)" json:"location"`
	StartDate   string     `gorm:"type:varchar(50)" json:"start_date"`
	EndDate     string     `gorm:"type:varchar(50)" json:"end_date"`
	IsCurrent   bool       `gorm:"type:tinyint(1);default:0" json:"is_current"`
	Description string     `gorm:"type:text" json:"description"`
	SortOrder   int        `gorm:"type:int;default:0" json:"sort_order"`
	CreatedAt   time.Time  `gorm:"type:datetime(3)" json:"created_at"`
	UpdatedAt   time.Time  `gorm:"type:datetime(3)" json:"updated_at"`
	DeletedAt   *time.Time `gorm:"type:datetime(3);index" json:"deleted_at,omitempty"`
}

func (Experience) TableName() string {
	return "experiences"
}
