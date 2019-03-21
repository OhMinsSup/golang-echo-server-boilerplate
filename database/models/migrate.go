package models

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/jinzhu/gorm"
	"github.com/pborman/uuid"
)

// Migrate automigrates models using ORM
func Migrate(db *gorm.DB) {
	db.AutoMigrate(&User{}, &EmailAuth{})
	//	db.Model(&UserModel{}).Related(&UserProfileModel{}, "UserID")
	//	db.Model(&SocicalAccountModel{}).Related(&UserModel{})
	fmt.Println("Auto Migration has beed processed")
}

func BeforeCreateUUID(scope *gorm.Scope) {
	reflectValue := reflect.Indirect(reflect.ValueOf(scope.Value))
	if strings.Contains(string(reflectValue.Type().Field(0).Tag), "uuid") {
		uuid.SetClockSequence(-1)
		scope.SetColumn("id", uuid.NewUUID().String())
	}
}
