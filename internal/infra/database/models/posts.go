package models

import (
	"time"
)

// Post struct used to store post data
type Post struct {
	ID        uint `gorm:"primaryKey; auto_increment"`
	Title     string
	Text      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
