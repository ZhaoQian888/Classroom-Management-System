package api

import (
	"Classroom-Management-System/information"
	"Classroom-Management-System/service"
	"fmt"
	"github.com/gin-gonic/gin"
)

// UserLogin 处理用户登陆
func UserLogin(c *gin.Context) {
	var userinfo service.UserLogInService
	if err := c.ShouldBind(&userinfo); err != nil {
		if !userinfo.CheckUserValid() {
			c.JSON(400, information.Response{
				Status: 13000,
				Msg:    "用户名或用户密码不合法",
			})
		}
		//此处放错误处理
	} else {
		user, err := userinfo.LogIn()
		if err != nil {
			fmt.Println("错误:", err)
		} else {
			fmt.Println("eee呵呵呵呵呵呵呵呵呵呵呵呵呵")
			c.SetCookie("login", "test", 3600, "/", "localhost", true, true)
			//返回前段所需用户信息
			c.JSON(200, information.CreateUserRseponse(user))

		}
	}

}

// UserRegister 处理用户注册
func UserRegister(c *gin.Context) {
	var userinfo service.UserRegisterService
	if err := c.ShouldBind(&userinfo); err != nil {
		fmt.Println(userinfo)
		c.JSON(304, information.CreateErrorResponse(err))
	} else {
		user, err := userinfo.Register()
		if err != nil {
			c.JSON(400, err)
		} else {
			c.SetCookie("login", "fdffdsfsdffsd", 3600, "/", "localhost", true, true)
			c.JSON(200, information.CreateUserRseponse(user))
		}
	}

}
