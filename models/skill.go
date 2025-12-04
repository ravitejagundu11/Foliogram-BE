package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Skill represents a skill in a portfolio
type Skill struct {
	ID          string `gorm:"type:char(36);primaryKey" json:"id"`
	PortfolioID string `gorm:"type:char(36);not null;index" json:"portfolio_id"`
	Name        string `gorm:"type:varchar(255);not null" json:"name"`
	Category    string `gorm:"type:varchar(100)" json:"category"` // e.g., "Frontend", "Backend", "Tools"
	Level       string `gorm:"type:varchar(50)" json:"level"`     // e.g., "Beginner", "Intermediate", "Expert"
	Percentage  int    `gorm:"type:int" json:"percentage"`        // 0-100 for visual representation
	IconURL     string `gorm:"type:varchar(500)" json:"icon_url"`
	SortOrder   int    `gorm:"default:0" json:"sort_order"`

	// Association
	Portfolio Portfolio `gorm:"foreignKey:PortfolioID" json:"-"`

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

// BeforeCreate hook to generate UUID
func (s *Skill) BeforeCreate(tx *gorm.DB) error {
	if s.ID == "" {
		s.ID = uuid.New().String()
	}
	return nil
}

// TableName specifies the table name
func (Skill) TableName() string {
	return "skills"
}
