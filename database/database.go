package database

import (
	"fmt"
	"os"

	"letterbot-subscription/database/models"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/labstack/gommon/log"
)

// Connect : Database connect
func Connect() *gorm.DB {
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	db, err := gorm.Open("mysql", fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local", user, password, dbName))
	db.LogMode(true)

	if err != nil {
		panic(err)
	}

	models.Migrate(db)
	log.Info("Connected Database")

	return db
}
