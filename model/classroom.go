package model

import (
	"Classroom-Management-System/information"
	"github.com/jinzhu/gorm"
)

const unused uint8 = 0
const using uint8 = 1

type entity interface {
	Ceate()
}

// Building 是一个教学楼实体
type Building struct {
	gorm.Model
	BuildingNumber uint8  `gorm:"primary_key"`
	BuildingName   string `gorm:"not null"`
}

// ClassRoom 是一个教室实体
type ClassRoom struct {
	gorm.Model
	Buildings  Building `gorm:"foreignkey:Building"`
	Floor      uint8    `gorm:"not null"`
	RoomNumber uint64   `gorm:"primary_key;not null;index:addr"`
}

// RoomStatus 代表教室当前的状态
type RoomStatus struct {
	gorm.Model
	Room     ClassRoom `gorm:"foreignkey:ClassRoomRefer"`
	Time     uint64
	TimeZone uint8
	Status   uint8 `gorm:"default:'0'"`
}

// Create 用来新建build
func (b *Building) Create() (information.Response, error) {
	count := 0
	DB.Where("BuildingNuber=?", b.BuildingNumber).Count(&count)
	if count == 0 {
		if err := DB.Create(&b).Error; err != nil {
			return information.Response{
				Status: 40010,
				Msg:    "数据库错误，请联系管理员",
			}, err
		}
		return information.Response{
			Status: 0,
			Msg:    "创建成功",
		}, nil
	}
	return information.Response{
		Status: 40011,
		Msg:    "教学楼已存在",
	}, nil

}

// Create 用来新建教室
func (c *ClassRoom) Create() *information.Response {
	count := 0
	DB.Where("BuildingNuber=?", c.RoomNumber).Count(&count)
	if count == 0 {
		if err := DB.Create(&c).Error; err != nil {
			return &information.Response{
				Status: 40013,
				Msg:    "数据库错误，请联系管理员",
			}
		}
		return &information.Response{
			Status: 0,
			Msg:    "创建成功",
		}
	}
	return &information.Response{
		Status: 40014,
		Msg:    "教室已存在",
	}

}
