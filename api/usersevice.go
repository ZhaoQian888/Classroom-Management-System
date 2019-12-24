package api

import (
	"Classroom-Management-System/information"
	"Classroom-Management-System/model"
	"Classroom-Management-System/service"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// UserLogin 处理用户登陆
func UserLogin(c *gin.Context) {
	var userinfo service.UserLogInService
	if err := c.ShouldBind(&userinfo); err != nil {
		// 这里处理前段传过来的信息不符
		if !userinfo.CheckUserValid() {
			c.JSON(200, information.Response{
				Status: 11001,
				Msg:    "用户名或用户密码不合法",
			})
		} else {
			c.JSON(200, information.Response{
				Status: 11002,
				Msg:    "用户信息未知错误",
			})
		}
	} else {
		user, err := userinfo.LogIn()
		if err != nil {
			c.JSON(200, err)
		} else {
			//此处正确，返回浏览器一个cookie
			s := sessions.Default(c)
			s.Clear()
			s.Set("user_id", user.Identity)
			s.Save()
			//并且把前段所需用户信息返回
			c.JSON(200, information.Response{
				Status: 0,
				Msg:    "登录成功",
				Data:   user,
			})
		}
	}

}

// UserRegister 处理用户注册
func UserRegister(c *gin.Context) {
	var userinfo service.UserRegisterService
	if err := c.ShouldBind(&userinfo); err != nil {
		c.JSON(200, information.Response{
			Status: 10005,
			Msg:    "用户信息不合法",
		})
	} else {
		user, err := userinfo.Register()
		if err != nil {
			c.JSON(200, err)
		} else {
			s := sessions.Default(c)
			s.Clear()
			s.Set("user_id", user.Identity)
			s.Save()

			c.JSON(200, information.Response{
				Status: 0,
				Msg:    "注册成功",
				Data:   user,
			})
		}
	}

}

// Quit 退出登录
func Quit(c *gin.Context) {
	s := sessions.Default(c)
	s.Clear()
	s.Save()
	c.JSON(200, information.Response{
		Status: 0,
		Msg:    "注销成功",
	})
}

// Status 返回当前状态
func Status(c *gin.Context) {
	u, ok := c.Get("user")
	if !ok {
		c.JSON(200, information.Response{
			Status: 80001,
			Msg:    "未登录",
		})
	} else {
		user := u.(*model.User)
		c.JSON(200, information.Response{
			Status: 0,
			Msg:    "用户以登录",
			Data:   user,
		})
	}

}
