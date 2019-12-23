package api

import (
	"Classroom-Management-System/model"
	"github.com/gin-gonic/gin"
)

// Myinfo 用来返回用户信息
func Myinfo(c *gin.Context) {
	u, _ := c.Get("user")
	user := u.(*model.User)
	c.JSON(200, user.Info())
}

// Roominfo 用来返回教室信息
func Roominfo(c *gin.Context) {
	c.JSON(200, model.RoomModleInfo())
}
