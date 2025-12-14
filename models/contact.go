package models

import "time"

type Contact struct {
	ID          string     `gorm:"type:char(36);primaryKey" json:"id"`
	PortfolioID string     `gorm:"type:char(36);not null;index;uniqueIndex:idx_portfolio_contact" json:"portfolio_id"`
	Email       string     `gorm:"type:varchar(255)" json:"email"`
	Phone       string     `gorm:"type:varchar(50)" json:"phone"`
	Address     string     `gorm:"type:text" json:"address"`
	City        string     `gorm:"type:varchar(100)" json:"city"`
	Country     string     `gorm:"type:varchar(100)" json:"country"`
	MapURL      string     `gorm:"type:varchar(500)" json:"map_url"`
	CreatedAt   time.Time  `gorm:"type:datetime(3)" json:"created_at"`
	UpdatedAt   time.Time  `gorm:"type:datetime(3)" json:"updated_at"`
	DeletedAt   *time.Time `gorm:"type:datetime(3);index" json:"deleted_at,omitempty"`
}

func (Contact) TableName() string {
	return "contacts"
}
