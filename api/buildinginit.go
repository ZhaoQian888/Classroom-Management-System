package api

import (
	"Classroom-Management-System/information"
	"Classroom-Management-System/model"

	"github.com/gin-gonic/gin"
)

func buildinginit(c *gin.Context) {
	u, _ := c.Get("user")
	user, _ := u.(model.User)
	if user.Username == "admin" {
		var build model.Building
		if err := c.ShouldBind(&build); err != nil {
			c.JSON(200, information.Response{
				Status: 40012,
				Msg:    "序列化失败",
			})
		} else {
			info, _ := build.Create()
			c.JSON(200, info)
		}
	} else {
		c.JSON(200, &information.Response{
			Status: 40002,
			Msg:    "不是管理员，无法进行操作",
		})
	}

}
