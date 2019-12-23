package model

import (
	"Classroom-Management-System/information"
	"time"

	"github.com/jinzhu/gorm"
)

// ClassRoomOrder 是一个预定关系实体
type ClassRoomOrder struct {
	gorm.Model
	RoomUser    string
	OrderTime   time.Time
	UseTime     time.Time
	UseTimeZone uint8
	Identity    uint64
	RoomNumber  uint64
}
type timezone uint8

const am timezone = 1
const pm timezone = 2
const night timezone = 3

// Order 完成数据库层面的订单
func (o *ClassRoomOrder) Order() (*information.Response, error) {
	var cro ClassRoom
	count := 0
	DB.First(&cro, "room_number=?", o.RoomNumber).Count(&count)
	if count == 0 {
		return &information.Response{
			Status: 30004,
			Msg:    "不存在的教室",
		}, nil
	}
	if err := DB.First(&cro, "room_number=?", o.RoomNumber).Find(&cro).Error; err != nil {
		return &information.Response{
			Status: 30004,
			Msg:    "数据库错误，联系管理员",
		}, err
	}
	var rs RoomStatus
	count = 0
	DB.Where("RoomNumber=? and Time=? and TimeZone=?", o.RoomNumber, o.UseTime, o.UseTimeZone).Count(&count)
	if count == 0 {
		rs = RoomStatus{
			Room:     cro,
			Time:     o.UseTime,
			TimeZone: o.UseTimeZone,
			Status:   1,
		}
		DB.Create(&rs)
		return &information.Response{
			Status: 0,
			Msg:    "教室可以使用",
		}, nil
	}
	DB.Where("RoomNumber=? and Time=? and TimeZone=?", o.RoomNumber, o.UseTime, o.UseTimeZone).Find(&rs)
	if rs.Status == 0 {
		DB.Model(&rs).Update("Status=", 1)
		return &information.Response{
			Status: 0,
			Msg:    "教室可以使用",
		}, nil
	}
	return &information.Response{
		Status: 30006,
		Msg:    "教室已被占用",
	}, nil
}
