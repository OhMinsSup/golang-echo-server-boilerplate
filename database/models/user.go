package models

type User struct {
	id       string `gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	username string `gorm:"column:username;unique"`
	email    string `gorm:"column":email;unique`
}
