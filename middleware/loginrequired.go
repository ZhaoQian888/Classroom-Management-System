package middleware

import (
	"github.com/gin-gonic/gin"
)

// LoginRequired 中间件，对于需要登陆之后才能操作的使用登陆验证
func LoginRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		// if logincookie, _ := c.Get("logincookie"); logincookie != nil {
		// 	if ok := CheckcookiedValid(logincookie); ok {
		// 		c.Next()
		// 		return
		// 	}
		// }
		// c.JSON(200, information.Response{
		// 	Status: 4001,
		// 	Msg:    "cookie过期或者没有，需要登陆",
		// })

	}

}

// func CheckcookiedValid(cookie string) bool {

// }
