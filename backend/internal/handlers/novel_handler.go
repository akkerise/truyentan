package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/truyentan/backend/internal/services"
)

// NovelHandler provides HTTP handlers for novel endpoints.
type NovelHandler struct {
	svc *services.NovelService
}

// NewNovelHandler creates a new NovelHandler.
func NewNovelHandler(svc *services.NovelService) *NovelHandler {
	return &NovelHandler{svc: svc}
}

// ListNovels godoc
// @Summary List novels
// @Tags novels
// @Param page query int false "page number"
// @Param limit query int false "page size"
// @Param query query string false "search query"
// @Param genre query string false "genre filter"
// @Param status query string false "status filter"
// @Success 200 {array} services.NovelDetail
// @Router /api/v1/novels [get]
func (h *NovelHandler) ListNovels(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	query := c.Query("query")
	genre := c.Query("genre")
	status := c.Query("status")
	novels, err := h.svc.ListNovels(page, limit, query, genre, status)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, novels)
}

// GetNovel godoc
// @Summary Get novel detail
// @Tags novels
// @Param id path int true "novel id"
// @Success 200 {object} services.NovelDetail
// @Failure 404 {object} gin.H
// @Router /api/v1/novels/{id} [get]
func (h *NovelHandler) GetNovel(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	novel, err := h.svc.GetNovel(uint(id))
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, novel)
}

// ListChapters godoc
// @Summary List chapters of a novel
// @Tags novels
// @Param id path int true "novel id"
// @Param page query int false "page number"
// @Param limit query int false "page size"
// @Success 200 {array} models.Chapter
// @Router /api/v1/novels/{id}/chapters [get]
func (h *NovelHandler) ListChapters(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	chapters, err := h.svc.ListChapters(uint(id), page, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, chapters)
}
