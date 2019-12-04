package api

import (
	"Classroom-Management-System/information"
	"Classroom-Management-System/service"

	"github.com/gin-gonic/gin"
)

// OrderClassroom 预定教室接口
func OrderClassroom(c *gin.Context) {
	var order service.ClassRoomOrder
	if err := c.ShouldBind(&order); err != nil {
		c.JSON(200, information.Response{
			Status: 30003,
			Msg:    "预定信息无法序列化，请检查错误",
		})
	} else {
		res := order.Order()
		c.JSON(200, res)
	}
}
