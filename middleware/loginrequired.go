package middleware

import (
	"Classroom-Management-System/information"
	"Classroom-Management-System/model"

	"github.com/gin-gonic/gin"
)

// LoginRequired 中间件，对于需要登陆之后才能操作的使用登陆验证
func LoginRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		user, _ := c.Get("user")
		if user != nil {
			if _, ok := user.(*model.User); ok {
				c.Next()
				return
			}
		} else {
			c.JSON(200, information.Response{
				Status: 20001,
				Msg:    "需要登陆",
			})
			c.Abort()
		}

	}

}

// AdminRequired 管理员登录中间件
func AdminRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		user, _ := c.Get("user")
		u, _ := user.(*model.User)
		if u.Username == "administrator" {
			c.Next()

		} else {
			c.JSON(200, information.Response{
				Status: 50001,
				Msg:    "需要管理员身份",
			})
			c.Abort()
		}
	}
}
