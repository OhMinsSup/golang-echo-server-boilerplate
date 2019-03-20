package models

type EmailAuthModel struct {
	id     string `gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	code   string `gorm:"size:255"`
	email  string `gorm:"size:255"`
	logged bool   `gorm:"default:false"`
}
