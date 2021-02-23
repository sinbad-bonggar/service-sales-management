package model

import (
	"time"
)

// Base is gorm.Model definition
type Base struct {
	ID        uint       `json:"id" gorm:"primaryKey;column:id;type:bigserial;autoIncrement"`
	CreatedAt time.Time  `json:"createdAt" gorm:"not null;default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time  `json:"updatedAt" gorm:"not null;default:CURRENT_TIMESTAMP"`
	DeletedAt *time.Time `json:"deletedAt,omitempty" gorm:"index;default:null"`
}
