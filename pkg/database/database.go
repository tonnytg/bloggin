package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDatabase() {
	cf := Config()
	dsn := fmt.Sprintf("host=%s port=%s sslmode=%s user=%s password=%s dbname=%s",
		cf.Host, cf.Port, cf.SSLMode, cf.User, cf.Password, cf.Name)
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}
	DB = database
}