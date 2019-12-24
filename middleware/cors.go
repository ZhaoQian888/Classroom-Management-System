package middleware

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// Cors 域名控制
func Cors() gin.HandlerFunc {
	config := cors.DefaultConfig()
	config.AllowMethods = []string{"GET", "POST"}
	config.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "cookie"}
	config.AllowOrigins = []string{"http://localhost:80", "http://frontend:80", "http://47.93.193.91", "http://localhost:8080"}
	config.AllowCredentials = true
	return cors.New(config)
}
