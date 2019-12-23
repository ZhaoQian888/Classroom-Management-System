package service

import (
	"Classroom-Management-System/information"
	"Classroom-Management-System/model"
	"time"
)

// ClassRoomOrder 前段发送的预定实体
type ClassRoomOrder struct {
	RoomNumber  uint64    `form:"room_number"    json:"room_number"     binding:"required"`
	UserName    string    `form:"user_name"      json:"user_name"       binding:"required"`
	RoomUser    string    `form:"room_user"      json:"room_user"       binding:"required"`
	OrderTime   time.Time `form:"order_time"     json:"order_time"      binding:"required" time_format:"2006-01-02"`
	UseTime     time.Time `form:"use_time"       json:"use_time"        binding:"required" time_format:"2006-01-02"`
	UseTimeZone uint8     `form:"use_time_zone"  json:"use_time_zone"   binding:"required"`
}

// Order 完成预定动作
func (o *ClassRoomOrder) Order(id uint64) *information.Response {
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
	postorder := model.ClassRoomOrder{
		RoomNumber:  o.RoomNumber,
		RoomUser:    o.RoomUser,
		OrderTime:   o.OrderTime,
		UseTime:     o.UseTime,
		UseTimeZone: o.UseTimeZone,
		Identity:    identity,
	}
	info, _ := postorder.Order()
	// 这里应该调用函数预定
	return info
}
