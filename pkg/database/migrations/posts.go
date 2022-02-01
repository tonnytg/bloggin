package migrations

import (
	"bloggin/pkg/database/models"
	"gorm.io/gorm"
)

func StartMigration(db *gorm.DB) {

	var Posts models.Post

	// Create table for `User`
	db.Migrator().CreateTable(Posts)
}
