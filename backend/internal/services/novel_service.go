package services

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"

	"github.com/truyentan/backend/internal/app/cache"
	"github.com/truyentan/backend/internal/models"
)

// NovelDetail represents novel details with total chapter count.
type NovelDetail struct {
	models.Novel
	TotalChapters int `json:"total_chapters"`
}

// NovelService provides novel related operations.
type NovelService struct {
	repo  *NovelRepository
	cache *redis.Client
}

// NewNovelService creates a new NovelService.
func NewNovelService(repo *NovelRepository, cache *redis.Client) *NovelService {
	return &NovelService{repo: repo, cache: cache}
}

// ListNovels returns novels with optional filters using cache.
func (s *NovelService) ListNovels(page, limit int, query, genre, status string) ([]models.Novel, error) {
	ctx := context.Background()
	key := fmt.Sprintf("novels:%d:%d:%s:%s:%s", page, limit, query, genre, status)
	var novels []models.Novel
	if err := cache.GetJSON(ctx, s.cache, key, &novels); err == nil {
		return novels, nil
	}
	novels, err := s.repo.ListNovels(page, limit, query, genre, status)
	if err != nil {
		return nil, err
	}
	_ = cache.SetJSON(ctx, s.cache, key, novels, 300*time.Second)
	return novels, nil
}

// GetNovel returns novel detail with genres and total chapter count using cache.
func (s *NovelService) GetNovel(id uint) (*NovelDetail, error) {
	ctx := context.Background()
	key := fmt.Sprintf("novel:%d", id)
	var detail NovelDetail
	if err := cache.GetJSON(ctx, s.cache, key, &detail); err == nil {
		return &detail, nil
	}
	novel, count, err := s.repo.GetNovel(id)
	if err != nil {
		return nil, err
	}
	detail = NovelDetail{Novel: *novel, TotalChapters: count}
	_ = cache.SetJSON(ctx, s.cache, key, detail, 300*time.Second)
	return &detail, nil
}

// ListChapters returns chapters of a novel with pagination.
func (s *NovelService) ListChapters(novelID uint, page, limit int) ([]models.Chapter, error) {
	return s.repo.ListChapters(novelID, page, limit)
}
