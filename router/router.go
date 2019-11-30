package router

import (
	"Classroom-Management-System/api"

	"github.com/gin-gonic/gin"
)

// SetRouter 初始化路由
func SetRouter() *gin.Engine {
	router := gin.Default()
	u := router.Group("/user")
	{
		u.POST("/login", api.UserLogin)
	}

	return router
}
