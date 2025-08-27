package services

import (
	"github.com/truyentan/backend/internal/models"
	"gorm.io/gorm"
)

// ProgressService manages reading progress operations.
type ProgressService struct {
	db *gorm.DB
}

// NewProgressService creates a new ProgressService.
func NewProgressService(db *gorm.DB) *ProgressService {
	return &ProgressService{db: db}
}

// Save stores or updates user's reading progress.
func (s *ProgressService) Save(userID, novelID, chapterID uint, position int) error {
	var rp models.ReadingProgress
	err := s.db.Where("user_id = ? AND novel_id = ?", userID, novelID).First(&rp).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			rp = models.ReadingProgress{UserID: userID, NovelID: novelID, ChapterID: chapterID, Position: position}
			return s.db.Create(&rp).Error
		}
		return err
	}
	rp.ChapterID = chapterID
	rp.Position = position
	return s.db.Save(&rp).Error
}

// Get retrieves user's progress for a novel.
func (s *ProgressService) Get(userID, novelID uint) (*models.ReadingProgress, error) {
	var rp models.ReadingProgress
	if err := s.db.Where("user_id = ? AND novel_id = ?", userID, novelID).First(&rp).Error; err != nil {
		return nil, err
	}
	return &rp, nil
}
