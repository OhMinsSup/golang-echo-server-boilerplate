package database

import (
	"fmt"
	"os"

	"github.com/OhMinsSup/pin-server/database/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// Initialize 데이터베이스 설정
func Initialize() (*gorm.DB, error) {
	dbConfig := os.Getenv("DB_CONFIG")
	db, err := gorm.Open("postgres", dbConfig)

	// logs SQL
	db.LogMode(true)
	db.Set("gorm:table_options", "charset=utf8")
	// created uuid
	db.Callback().Create().Before("gorm:create").Register("my_plugin:before_create", models.BeforeCreateUUID)
	if err != nil {
		panic(err)
	}
	fmt.Println("Connected to database")
	models.Migrate(db)

	return db, err
}
