package models

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Name      string    `json:"name"`
	Email     string    `gorm:"unique" json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

func MigrateUsers(db *gorm.DB) {
	err := db.AutoMigrate(&User{})
	if err != nil {
		return
	}
}
