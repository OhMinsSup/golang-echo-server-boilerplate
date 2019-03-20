package models

type UserModel struct {
	id        string           `gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	username  string           `gorm:"column:username;unique_index"`
	email     string           `gorm:"column:email;unique_index"`
	createdAt uint64           `gorm:"type:timestamp"`
	updatedAt uint64           `gorm:"type:timestamp"`
	profile   UserProfileModel `gorm:"foreignkey:UserID"`
}
