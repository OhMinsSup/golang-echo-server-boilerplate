package models

type UserModel struct {
	id         string `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
	username   string `gorm:"column:username;unique" json:"username"`
	email      string `gorm:"column:email;unique" json:"email"`
	created_at uint64 `gorm:"type:timestamp" json:"created_at"`
	updated_at uint64 `gorm:"type:timestamp" json:"updated_at"`
}
