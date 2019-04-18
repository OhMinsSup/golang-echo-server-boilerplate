package models

import (
	"time"
)

// User 모델
type User struct {
	ID        string `gorm:"primary_key;uuid"`
	Username  string `gorm:"size:255;unique_index"`
	Email     string `gorm:"size:255;unique_index"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time

	// hashOne
	UserProfile UserProfile `gorm:"foreignkey:UserID;association_foreignkey:Refer"`
}
