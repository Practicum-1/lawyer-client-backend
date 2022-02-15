package db

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	dsn := "root:password@tcp(127.0.0.1:3306)/lawyer_client?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	fmt.Println("Connected to database")
	DB = db
}

func GetDB() *gorm.DB {
	return DB
}

func ConnectDB() {
	Connect() // Connect to MySQL database.
}
