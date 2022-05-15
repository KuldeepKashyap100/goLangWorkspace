package config

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func Connect() {
	connection, err := gorm.Open(sqlite.Open("mydb.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db = connection
	// db, _ = connection.DB()
	// db.SetMaxIdleConns(100)
	// db.SetMaxOpenConns(100)
	// db.SetConnMaxLifetime(time.Hour)
}

func GetDB() *gorm.DB {
	return db
}
