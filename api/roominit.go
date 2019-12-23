package api

import (
	"Classroom-Management-System/information"
	"Classroom-Management-System/service"

	"github.com/gin-gonic/gin"
)

// Roominit 用来管理员添加教室
func Roominit(c *gin.Context) {
	var r service.RoomInit
	if err := c.ShouldBind(&r); err != nil {
		c.JSON(200, information.Response{
			Status: 40015,
			Msg:    "教室序列化失败",
		})
	} else {
		info := r.Init()
		c.JSON(200, info)
	}

}
