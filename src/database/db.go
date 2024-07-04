package database

import (
	"go-authentication/src/models"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	dsn := "root:truong123456@tcp(localhost:3306)/sys?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Ensure the connection is using the database 'sys'
	useSysDB := "USE sys"
	if err := DB.Exec(useSysDB).Error; err != nil {
		log.Fatal("Failed to select database 'sys':", err)
	}

	// Automatically migrate your schema
	if err := DB.AutoMigrate(&models.User{}); err != nil {
		log.Fatal("Failed to run migrations:", err)
	}

	DB.AutoMigrate(&models.User{})
	DB.AutoMigrate(&models.Message{})
	DB.AutoMigrate(&models.Channel{})
}
