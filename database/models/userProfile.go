package models

import (
	"database/sql"
)

type UserProfileModel struct {
	id           string         `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
	display_name string         `gorm:"type:varchar(255)" json:"display_name"`
	thumbnail    sql.NullString `gorm:"type:varchar(255);null" json:"thumbnail"`
	short_bio    string         `gorm:"type:varchar(255)" json:"short_bio"`
	created_at   uint64         `gorm:"type:timestamp" json:"created_at"`
	updated_at   uint64         `gorm:"type:timestamp" json:"updated_at"`
}
