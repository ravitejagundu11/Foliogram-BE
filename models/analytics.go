package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// AnalyticsView represents a portfolio view for analytics
type AnalyticsView struct {
	ID          string    `gorm:"type:char(36);primaryKey" json:"id"`
	PortfolioID string    `gorm:"type:char(36);not null;index" json:"portfolio_id"`
	ViewerIP    string    `gorm:"type:varchar(45)" json:"viewer_ip"`
	SessionID   string    `gorm:"type:varchar(255)" json:"session_id"`
	UserAgent   string    `gorm:"type:varchar(500)" json:"user_agent"`
	Referrer    string    `gorm:"type:varchar(500)" json:"referrer"`
	ViewedAt    time.Time `gorm:"index" json:"viewed_at"`

	// Association
	Portfolio Portfolio `gorm:"foreignKey:PortfolioID" json:"-"`

	CreatedAt time.Time `json:"created_at"`
}

// BeforeCreate hook to generate UUID
func (a *AnalyticsView) BeforeCreate(tx *gorm.DB) error {
	if a.ID == "" {
		a.ID = uuid.New().String()
	}
	if a.ViewedAt.IsZero() {
		a.ViewedAt = time.Now()
	}
	return nil
}

// TableName specifies the table name
func (AnalyticsView) TableName() string {
	return "analytics_views"
}
