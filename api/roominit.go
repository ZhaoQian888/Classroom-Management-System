package api

import (
	"Classroom-Management-System/information"
	"Classroom-Management-System/model"

	"github.com/gin-gonic/gin"
)

func roominit(c *gin.Context) {
	u, _ := c.Get("user")
	user, _ := u.(model.User)
	if user.Username == "admin" {
		var r model.ClassRoom
		if err := c.ShouldBind(&r); err != nil {
			c.JSON(200, information.Response{
				Status: 40015,
				Msg:    "教室学历化失败",
			})
		} else {
			info := r.Create()
			c.JSON(200, info)
		}
	} else {
		c.JSON(200, &information.Response{
			Status: 40002,
			Msg:    "不是管理员，无法进行操作",
		})
	}

}
