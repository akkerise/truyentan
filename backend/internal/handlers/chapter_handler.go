package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/truyentan/backend/internal/app/kafka"
	"github.com/truyentan/backend/internal/services"
)

// ChapterHandler handles chapter related endpoints.
type ChapterHandler struct {
	svc      *services.ChapterService
	producer *kafka.Producer
}

// NewChapterHandler creates a new ChapterHandler.
func NewChapterHandler(svc *services.ChapterService, producer *kafka.Producer) *ChapterHandler {
	return &ChapterHandler{svc: svc, producer: producer}
}

// chapterResponse represents chapter payload with navigation info.
type chapterResponse struct {
	ID      uint   `json:"id"`
	NovelID uint   `json:"novel_id"`
	Number  int    `json:"number"`
	Title   string `json:"title"`
	Content string `json:"content"`
	PrevID  *uint  `json:"prev_id,omitempty"`
	NextID  *uint  `json:"next_id,omitempty"`
}

// Get returns the chapter by id.
func (h *ChapterHandler) Get(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	ch, prevID, nextID, err := h.svc.GetByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "chapter not found"})
		return
	}
	_ = h.producer.PublishNovelRead(c.Request.Context(), ch.NovelID, ch.ID)
	res := chapterResponse{ID: ch.ID, NovelID: ch.NovelID, Number: ch.Number, Title: ch.Title, Content: ch.Content, PrevID: prevID, NextID: nextID}
	c.JSON(http.StatusOK, res)
}

// GetNext returns the next chapter.
func (h *ChapterHandler) GetNext(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	ch, prevID, nextID, err := h.svc.GetNext(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "chapter not found"})
		return
	}
	_ = h.producer.PublishNovelRead(c.Request.Context(), ch.NovelID, ch.ID)
	res := chapterResponse{ID: ch.ID, NovelID: ch.NovelID, Number: ch.Number, Title: ch.Title, Content: ch.Content, PrevID: prevID, NextID: nextID}
	c.JSON(http.StatusOK, res)
}

// GetPrev returns the previous chapter.
func (h *ChapterHandler) GetPrev(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	ch, prevID, nextID, err := h.svc.GetPrev(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "chapter not found"})
		return
	}
	_ = h.producer.PublishNovelRead(c.Request.Context(), ch.NovelID, ch.ID)
	res := chapterResponse{ID: ch.ID, NovelID: ch.NovelID, Number: ch.Number, Title: ch.Title, Content: ch.Content, PrevID: prevID, NextID: nextID}
	c.JSON(http.StatusOK, res)
}
