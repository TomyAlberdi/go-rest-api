package models

import "gorm.io/gorm"

type Post struct {
	ID     uint   `gorm:"primaryKey" json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
	UserID uint   `json:"user_id"` // Foreign key
}

func MigratePosts(db *gorm.DB) {
	err := db.AutoMigrate(&Post{})
	if err != nil {
		return
	}
}
