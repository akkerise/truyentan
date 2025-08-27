package db

import (
	"github.com/truyentan/backend/internal/models"
	"gorm.io/gorm"
)

// Migrate runs database migrations using GORM.
func Migrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&models.User{},
		&models.Novel{},
		&models.Chapter{},
		&models.Genre{},
		&models.NovelGenre{},
		&models.Favorite{},
		&models.ReadingProgress{},
	)
}
