package models

import "time"

// UserProfile 모델
type UserProfile struct {
	ID          string `gorm:"primary_key;uuid"`
	DisplayName string `gorm:"size:255"`
	thumbnail   string `gorm:"size:255"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   *time.Time

	UserID string `gorm:"primary_key;uuid"`
}
