package model

import (
	"Classroom-Management-System/information"
	"log"
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

func dealwithfail(flag bool, tx *gorm.DB) {
	if flag {
		tx.Rollback()
	}
}

// Order 完成数据库层面的订单
func (o *ClassRoomOrder) Order() (*information.Response, error) {
	var cro ClassRoom
	count := 0
	tx := DB.Begin()
	flag := false
	defer dealwithfail(flag, tx)
	err := tx.First(&cro, "room_number=?", o.RoomNumber).Count(&count).Error
	if err != nil {
		flag = true
	}
	if count == 0 {
		return &information.Response{
			Status: 30004,
			Msg:    "不存在的教室",
		}, nil
	}
	if err := tx.First(&cro, "room_number=?", o.RoomNumber).Error; err != nil {
		flag = true
		return &information.Response{
			Status: 30004,
			Msg:    "数据库错误，联系管理员",
		}, err
	}
	var rs RoomStatus
	err = tx.Model(&cro).Related(&cro.Rs, "ClassRoomRefer").Error
	if err != nil {
		flag = true
	}
	count = len(cro.Rs)
	if count == 0 {
		rs = RoomStatus{
			ClassRoomRefer: cro.ID,
			Time:           o.UseTime,
			TimeZone:       o.UseTimeZone,
			Status:         1,
		}
		tx.Create(&rs)
		tx.Commit()
		return &information.Response{
			Status: 0,
			Msg:    "教室可以使用",
		}, nil
	}
	for _, v := range cro.Rs {
		log.Println(v.Time, " ", v.TimeZone, "===", o.UseTime, " ", o.UseTimeZone)
		if v.Time.Equal(o.UseTime) && v.TimeZone == o.UseTimeZone && v.Status == 1 {
			tx.Commit()
			return &information.Response{
				Status: 70001,
				Msg:    "该时间段教室已被占用",
			}, nil
		}
	}
	rs = RoomStatus{
		ClassRoomRefer: cro.ID,
		Time:           o.UseTime,
		TimeZone:       o.UseTimeZone,
		Status:         1,
	}
	err = tx.Create(&rs).Error
	if err != nil {
		flag = true
	}
	tx.Commit()
	return &information.Response{
		Status: 0,
		Msg:    "教室可以使用",
	}, nil

}
