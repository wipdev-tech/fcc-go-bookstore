package config

import (
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func Connect() {
	dbuser := os.Getenv("DBUSER")
	dbpass := os.Getenv("DBPASS")

	dsn := fmt.Sprintf(
		"%s:%s@tcp(127.0.0.1:3306)/simplerest?charset=utf8mb4&parseTime=True&loc=Local",
		dbuser,
		dbpass,
	)
	d, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db = d
}

func GetDB() *gorm.DB {
	return db
}
