package model

import (
	"Classroom-Management-System/information"
	"time"

	"github.com/jinzhu/gorm"
)

const unused uint8 = 0
const using uint8 = 1

type entity interface {
	Ceate()
}

// RoomInfoRes 所有的教室信息
type RoomInfoRes struct {
	Cs []ClassRoom  `json:"cs"`
	Rs []RoomStatus `json:"rs"`
}

// RoomModleInfo 返回当前所有的教室信息
func RoomModleInfo() *information.Response {
	css := []ClassRoom{}
	rss := []RoomStatus{}
	err := DB.Find(&css).Error
	if err != nil {
		return &information.Response{
			Status: 60001,
			Msg:    "数据库错误，请联系管理员",
			Data:   err,
		}
	}
	err = DB.Where("Status=?", 1).Find(&rss).Error
	if err != nil {
		return &information.Response{
			Status: 60001,
			Msg:    "数据库错误，请联系管理员",
			Data:   err,
		}
	}
	rir := RoomInfoRes{
		Cs: css,
		Rs: rss,
	}
	return &information.Response{
		Status: 0,
		Data:   &rir,
	}
}

// BuildModelInfo 返回json
func BuildModelInfo(id uint8) *information.Response {
	var b Building
	DB.Where("id=?", id).Find(&b)
	return &information.Response{
		Status: 0,
		Msg:    "教学楼信息",
		Data:   b,
	}
}

// Building 是一个教学楼实体
type Building struct {
	gorm.Model
	BuildingNumber uint8       `gorm:"index:addr;unique"`
	BuildingName   string      `gorm:"not null"`
	Cr             []ClassRoom `gorm:"foreignkey:BuildingRefer;association_foreignkey:BuildingNumber"`
}

// ClassRoom 是一个教室实体
type ClassRoom struct {
	gorm.Model
	BuildingRefer uint
	Rs            []RoomStatus `gorm:"foreignkey:ClassRoomRefer;association_foreignkey:RoomNumber" `
	Floor         uint8        `gorm:"not null"`
	RoomNumber    uint64       `gorm:"not null;index:addr;unique"`
}

// RoomStatus 代表教室当前的状态
type RoomStatus struct {
	gorm.Model
	ClassRoomRefer uint
	Time           time.Time
	TimeZone       uint8
	Status         uint8 `gorm:"default:'0'"`
}

// Create 用来新建build
func (b *Building) Create() (information.Response, error) {
	count := 0
	DB.Where("Building_number=?", b.BuildingNumber).Find(&Building{}).Count(&count)
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
	DB.Where("room_number=?", c.RoomNumber).Find(&c).Count(&count)
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
