package app

import "github.com/gin-gonic/gin"

// NewServer creates and configures a new HTTP server.
func NewServer() *gin.Engine {
	r := gin.Default()
	return r
}
