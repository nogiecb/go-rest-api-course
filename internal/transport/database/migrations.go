package database

import (
	"github.com/jinzhu/gorm"
	"github.com/nogiecb/go-rest-api-course/internal/transport/comment"
)

// MigrateDB - migrate database and create comment table
func MigrateDB(db *gorm.DB) error {
	if result := db.AutoMigrate(&comment.Comment{}); result.Error != nil {
		return result.Error
	}
	return nil
}
