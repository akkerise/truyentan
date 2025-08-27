package db

import (
	"fmt"

	"github.com/truyentan/backend/internal/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// Seed inserts sample data into the database for demo purposes.
func Seed(db *gorm.DB) error {
	// Create test user
	password, err := bcrypt.GenerateFromPassword([]byte("password"), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user := models.User{
		Name:         "Test User",
		Email:        "test@example.com",
		PasswordHash: string(password),
	}
	if err := db.Create(&user).Error; err != nil {
		return err
	}

	// Create genres
	genres := []models.Genre{
		{Name: "Fantasy"},
		{Name: "Sci-Fi"},
		{Name: "Romance"},
		{Name: "Action"},
		{Name: "Comedy"},
	}
	if err := db.Create(&genres).Error; err != nil {
		return err
	}

	// Create novels with chapters and associate genres
	for i := 1; i <= 5; i++ {
		novel := models.Novel{
			Title:       fmt.Sprintf("Sample Novel %d", i),
			Author:      fmt.Sprintf("Author %d", i),
			Description: "Demo novel for seeding",
			Status:      "ongoing",
			Genres: []models.Genre{
				genres[(i-1)%len(genres)],
				genres[i%len(genres)],
			},
		}
		if err := db.Create(&novel).Error; err != nil {
			return err
		}

		chapters := make([]models.Chapter, 0, 20)
		for j := 1; j <= 20; j++ {
			chapters = append(chapters, models.Chapter{
				NovelID: novel.ID,
				Number:  j,
				Title:   fmt.Sprintf("Chapter %d", j),
				Content: fmt.Sprintf("Content for chapter %d of novel %d", j, i),
			})
		}
		if err := db.Create(&chapters).Error; err != nil {
			return err
		}
	}

	return nil
}
