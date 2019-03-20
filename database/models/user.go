package models

type UserModel struct {
	id        string           `gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	username  string           `gorm:"size:255;unique_index"`
	email     string           `gorm:"size:255;unique_index"`
	createdAt uint64           `gorm:"type:timestamp"`
	updatedAt uint64           `gorm:"type:timestamp"`
	profile   UserProfileModel `gorm:"foreignkey:UserID"`
}
