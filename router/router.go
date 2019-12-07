package router

import (
	"Classroom-Management-System/api"
	"Classroom-Management-System/middleware"
	"os"

	"github.com/gin-gonic/gin"
)

// SetRouter 初始化路由
func SetRouter() *gin.Engine {
	router := gin.Default()

	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	// router.Use(middleware.Cors())
	router.Use(middleware.Session((os.Getenv("SESSION_SECRET"))))
	router.Use(middleware.CurrentUser())

	user := router.Group("/gin/user")
	{
		user.POST("/register", api.UserRegister)
		user.POST("/login", api.UserLogin)
	}
	order := router.Group("/gin/order")
	{
		order.Use(middleware.LoginRequired())
		order.POST("/classroom", api.OrderClassroom)
	}

	return router
}
