package database

import (
	"bloggin/entity/post"
	"bloggin/pkg/database/migrations"
	"bloggin/pkg/database/models"
	"bloggin/pkg/logger"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

// InitDatabase initializes the database connection
func InitDatabase() *gorm.DB {

	// Set configuration of database
	cf := Config()
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cf.Host, cf.Port, cf.User, cf.Password, cf.Name, cf.SSLMode)

	// Connect to database and check connection errors
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {

		// Save error to log
		msg := fmt.Sprintf("Error connecting to database: %s", err.Error())
		logger.Msg("ERROR", msg)
		fmt.Println(err)
	}

	migrations.StartMigration(db)
	return db
}

func SavePost(p *post.Post) {

	db := InitDatabase()
	mp := models.Post{
		Title: p.Title,
		Text:  p.Text,
	}

	result := db.Create(&mp)
	fmt.Println(result.RowsAffected)
}

func DeletePost(id string) {

	var post models.Post
	db := InitDatabase()

	db.Where("id = ?", id).Delete(&post)
}

func GetAllArticles() []models.Post {
	db := InitDatabase()
	var posts []models.Post
	db.Find(&posts)
	return posts
}
