package models

import (
	"time"
)

type User struct {
	ID        string `gorm:"primary_key;uuid"`
	Username  string `gorm:"size:255;unique_index"`
	Email     string `gorm:"size:255;unique_index"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
	//	profile   UserProfile `gorm:"foreignkey:UserID"`
}
