package api

import (
	"Classroom-Management-System/information"
	"Classroom-Management-System/service"
	"github.com/gin-gonic/gin"
)

// UserLogin 处理用户登陆
func UserLogin(c *gin.Context) {
	var userinfo service.UserLogInService
	if err := c.ShouldBind(&userinfo); err != nil {
		//此处放错误处理
	}
	user, err := userinfo.LogIn()
	if err != nil {
		//错误处理
	} else {
		c.SetCookie("gin_cookie", "test", 3600, "/", "localhost", false, true)
		//返回前段所需用户信息
		c.JSON(200, information.CreateUserRseponse(user))

	}

}
