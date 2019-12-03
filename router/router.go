package router

import (
	"Classroom-Management-System/api"
	"Classroom-Management-System/middleware"

	"github.com/gin-gonic/gin"
)

// SetRouter 初始化路由
func SetRouter() *gin.Engine {
	router := gin.Default()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(middleware.Cors())
	user := router.Group("/user")
	user.Use(middleware.LoginRequired())
	{
		user.POST("/register", api.UserRegister)
		user.POST("/login", api.UserLogin)
	}

	return router
}
