package models

import (
	"database/sql"
)

type UserProfile struct {
	id          string         `gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	displayName string         `gorm:"size:255"`
	thumbnail   sql.NullString `gorm:"size:255"`
	shortBio    string         `gorm:"type:text"`
	UserID      string
	createdAt   uint64 `gorm:"type:timestamp"`
	updatedAt   uint64 `gorm:"type:timestamp"`
}
