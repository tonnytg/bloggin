package database

import (
	"bloggin/config/logger"
	"bloggin/internal/entity/post"
	"bloggin/internal/infra/database/migrations"
	"bloggin/internal/infra/database/models"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq" // Import pq driver
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

// InitDatabase initializes the database connection
func InitDatabase() *gorm.DB {
	cf := Config()
	createDBIfNotExist(cf)

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

func createDBIfNotExist(cf *DBConfg) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s sslmode=%s",
		cf.Host, cf.Port, cf.User, cf.Password, cf.SSLMode)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("Error connecting to PostgreSQL: %s", err)
	}
	defer db.Close()

	// Verificar se o banco de dados já existe
	var exists bool
	checkQuery := fmt.Sprintf("SELECT EXISTS(SELECT 1 FROM pg_database WHERE datname = '%s')", cf.Name)
	err = db.QueryRow(checkQuery).Scan(&exists)
	if err != nil {
		log.Fatalf("Error checking database existence: %s", err)
	}

	if exists {
		fmt.Printf("Database %s already exists, skipping creation.\n", cf.Name)
	} else {
		// Tentativa de criação do banco de dados
		_, err = db.Exec(fmt.Sprintf("CREATE DATABASE %s", cf.Name))
		if err != nil {
			log.Fatalf("Error creating database: %s", err)
		}
		fmt.Printf("Database %s created successfully.\n", cf.Name)
	}
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
	fmt.Println("array:", len(posts))
	db.Find(&posts)
	return posts
}
