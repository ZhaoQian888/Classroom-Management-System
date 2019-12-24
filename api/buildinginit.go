package api

import (
	"Classroom-Management-System/information"
	"Classroom-Management-System/service"

	"github.com/gin-gonic/gin"
)

// Buildinginit 用来管理员添加教学楼
func Buildinginit(c *gin.Context) {
	var build service.BuildingInit
	if err := c.ShouldBind(&build); err != nil {
		c.JSON(200, information.Response{
			Status: 40012,
			Msg:    "序列化失败",
		})
	} else {
		info := build.Init()
		c.JSON(200, *info)
	}
}

// DeleteBuild 管理员删除教学楼
func DeleteBuild(c *gin.Context) {
	var build service.BuildingDelete
	if err := c.ShouldBind(&build); err != nil {
		c.JSON(200, information.Response{
			Status: 40020,
			Msg:    "序列化失败",
		})
	} else {
		info := build.Delete()
		c.JSON(200, *info)
	}
}
