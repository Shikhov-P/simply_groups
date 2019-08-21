package models

import (
	utils "../utils"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
	"os"
)

var db *gorm.DB

func init() {
	err := godotenv.Load()
	utils.ParseError(err)

	username := os.Getenv("db_user")
	password := os.Getenv("db_pass")
	dbName := os.Getenv("db_name")
	dbHost := os.Getenv("db_host")

	connectionUri := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s",
		dbHost, username, dbName, password)

	connection, err := gorm.Open("postgres", connectionUri)
	utils.ParseError(err)

	db = connection
	//db.Model(&User{}).AddForeignKey()
	db.Debug().AutoMigrate(&User{}, &Group{})
}

func GetDB() *gorm.DB {
	return db
}