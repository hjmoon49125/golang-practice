package database

import (
	// "gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func Init() error {
	connDB()

	return nil
}

func connDB() {
	db = nil
}

func GetDB() *gorm.DB {
	return db
}
