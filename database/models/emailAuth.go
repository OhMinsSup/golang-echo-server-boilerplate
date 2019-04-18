package models

import "time"

// EmailAuth 모델
type EmailAuth struct {
	ID        string `gorm:"primary_key;uuid"`
	Code      string `gorm:"size:255"`
	Email     string `gorm:"size:255"`
	Logged    bool   `gorm:"default:false"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}
