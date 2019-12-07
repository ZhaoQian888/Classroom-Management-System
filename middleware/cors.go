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
	config.AllowOrigins = []string{"http://localhost:80", "http://frontend:80", "http://47.93.193.91:80", "http://0.0.0.0:80", "http://www.monks.top:80", "http://www.monks.top:8080",
		"http://localhost:8080", "http://frontend:8080", "http://47.93.193.91:8080", "http://0.0.0.0:8080", "http://172.22.0.4:80", "http://10.129.48.1", "http://10.129.54.44"}
	config.AllowCredentials = true
	return cors.New(config)
}
