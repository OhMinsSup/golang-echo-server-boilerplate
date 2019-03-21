package database

import (
	"fmt"
	"os"

	"github.com/OhMinsSup/pin-server/database/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func Initialize() (*gorm.DB, error) {
	dbConfig := os.Getenv("DB_CONFIG")
	db, err := gorm.Open("postgres", dbConfig)

	// logs SQL
	db.LogMode(true)
	// created uuid
	db.Callback().Create().Before("gorm:create").Register("my_plugin:before_create", models.BeforeCreateUUID)
	if err != nil {
		panic(err)
	}
	fmt.Println("Connected to database")
	models.Migrate(db)

	return db, err
}
