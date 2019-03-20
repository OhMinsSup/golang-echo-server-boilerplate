package models

type SocicalAccountModel struct {
	id          string    `gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	socialID    string    `gorm:"size:255"`
	accessToken string    `gorm:"size:255"`
	provider    string    `gorm:"size:255"`
	user        UserModel `gorm:"foreignkey:UserID"`
	UserID      string
}
