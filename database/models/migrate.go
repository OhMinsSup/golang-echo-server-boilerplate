package models

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

func Migrate(db *gorm.DB) {
	//	db.AutoMigrate(&UserModel{}, &UserProfileModel{}, &SocicalAccountModel{}, &EmailAuthModel{})
	//	db.Model(&UserModel{}).Related(&UserProfileModel{}, "UserID")
	//	db.Model(&SocicalAccountModel{}).Related(&UserModel{})
	fmt.Println("Auto Migration has beed processed")
}
