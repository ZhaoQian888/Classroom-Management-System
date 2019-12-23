package api

import (
	"Classroom-Management-System/information"
	"Classroom-Management-System/model"

	"github.com/gin-gonic/gin"
)

// Roominit 用来管理员添加教室
func Roominit(c *gin.Context) {
	var r model.ClassRoom
	if err := c.ShouldBind(&r); err != nil {
		c.JSON(200, information.Response{
			Status: 40015,
			Msg:    "教室序列化失败",
		})
	} else {
		info := r.Create()
		c.JSON(200, info)
	}

}
