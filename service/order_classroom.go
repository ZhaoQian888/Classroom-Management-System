package service

import (
	"Classroom-Management-System/information"
	"Classroom-Management-System/model"
	"fmt"
)

// ClassRoomOrder 前段发送的预定实体
type ClassRoomOrder struct {
	UserName     string `form:"user_name"      json:"user_name"      binding:"required"`
	RoomUser     string `form:"room_user"      json:"room_user"      binding:"required"`
	OrderTime    int64  `form:"order_time"     json:"order_time"     binding:"required"`
	UseTimeStart int64  `form:"use_time_start" json:"use_time_start" binding:"required"`
	UseTimeEnd   int64  `form:"use_time_end"   json:"use_time_end"   binding:"required"`
}

// Order 完成预定动作
func (o *ClassRoomOrder) Order(id interface{}) *information.Response {
	var user model.User
	count := 0
	model.DB.Where("identity = ?", id).First(&user).Count(&count)
	identity := user.Identity
	if count == 0 || identity == 0 {
		return &information.Response{
			Status: 30005,
			Msg:    "身份验证失败",
		}
	}
	order := model.ClassRoomOrder{
		RoomUser:     o.RoomUser,
		OrderTime:    o.OrderTime,
		UseTimeStart: o.UseTimeStart,
		UseTimeEnd:   o.UseTimeEnd,
		UserName:     o.UserName,
		Identity:     identity,
	}
	err := model.DB.Create(&order).Error
	if err != nil {
		fmt.Println(err)
		return &information.Response{
			Status: 30001,
			Msg:    "预定失败，数据库错误",
		}
	}
	return &information.Response{
		Status: 0,
		Msg:    "预定成功",
	}
}
