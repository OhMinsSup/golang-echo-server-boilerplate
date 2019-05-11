package models

import (
	"github.com/OhMinsSup/pin-server/lib"
	"time"
)

// User 모델
type User struct {
	ID          string `gorm:"primary_key;uuid"`
	Email       string `gorm:"size:255;unique_index"`
	Username    string `gorm:"size:255;unique_index"`
	DisplayName string `gorm:"size:255"`
	Thumbnail   string `gorm:"size:255"`
	Password    string `gorm:"size:255"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   *time.Time
}

// Serialize 필요한 데이터를 가져오기
func (u *User) Serialize() lib.JSON {
	return lib.JSON{
		"id":          u.ID,
		"username":    u.Username,
		"displayName": u.DisplayName,
		"email":       u.Email,
	}
}
