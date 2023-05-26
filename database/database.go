package database

import (
	"JWT-authentication/models"
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var Db *gorm.DB

func Connection() (err error) {

	Db, err = gorm.Open(sqlite.Open("sqlite-database.db"), &gorm.Config{})

	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	Db.AutoMigrate(&models.User{})
	return err
}
