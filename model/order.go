package model

import "github.com/jinzhu/gorm"

import "Classroom-Management-System/information"

// ClassRoomOrder 是一个预定关系实体
type ClassRoomOrder struct {
	gorm.Model
	RoomUser    string
	OrderTime   uint64
	UseTime     uint64
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
	err := DB.First(&cro, "RoomNumber=?", o.RoomNumber).Error
	if err != nil {
		return &information.Response{
			Status: 30004,
			Msg:    "数据错误请询问管理员",
		}, err
	}
	var rs RoomStatus
	count := 0
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
