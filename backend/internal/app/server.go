package app

import (
	"time"

	"github.com/gin-contrib/cors"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"github.com/truyentan/backend/internal/app/cache"
	"github.com/truyentan/backend/internal/db"
	"github.com/truyentan/backend/internal/handlers"
	"github.com/truyentan/backend/internal/services"
)

// NewServer creates and configures a new HTTP server.
func NewServer() *gin.Engine {
	cfg, err := LoadConfig()
	if err != nil {
		panic(err)
	}

	logger, _ := zap.NewProduction()

	r := gin.New()
	r.Use(ginzap.Ginzap(logger, time.RFC3339, true))
	r.Use(ginzap.RecoveryWithZap(logger, true))
	r.Use(cors.Default())

	dbConn, err := db.Connect(cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName)
	if err != nil {
		logger.Fatal("failed to connect database", zap.Error(err))
	}
	redisClient := cache.NewRedis(cfg.RedisHost, cfg.RedisPort)

	authService := services.NewAuthService(dbConn, cfg.JWTSecret, cfg.AccessTokenTTL, cfg.RefreshTokenTTL)
	authHandler := handlers.NewAuthHandler(authService)
	novelRepo := services.NewNovelRepository(dbConn)
	novelService := services.NewNovelService(novelRepo, redisClient)
	novelHandler := handlers.NewNovelHandler(novelService)

	api := r.Group("/api/v1")
	auth := api.Group("/auth")
	auth.POST("/signup", authHandler.Signup)
	auth.POST("/signin", authHandler.Signin)
	auth.POST("/refresh", authHandler.Refresh)
	novels := api.Group("/novels")
	novels.GET("", novelHandler.ListNovels)
	novels.GET("/:id", novelHandler.GetNovel)
	novels.GET("/:id/chapters", novelHandler.ListChapters)

	r.GET("/healthz", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	return r
}
