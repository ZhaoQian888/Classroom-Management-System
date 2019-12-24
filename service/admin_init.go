package service

import (
	"Classroom-Management-System/information"
	"Classroom-Management-System/model"
)

// BuildingInit 序列化教学楼
type BuildingInit struct {
	BuildingName   string `form:"building_name"   json:"building_name"   binding:"required"`
	BuildingNumber uint8  `form:"building_number" json:"building_number" binding:"required"`
}

// Init 管理员用来初始化教学楼
func (b *BuildingInit) Init() *information.Response {
	i := model.Building{
		BuildingNumber: b.BuildingNumber,
		BuildingName:   b.BuildingName,
	}
	info, _ := i.Create()
	return &info

}

// RoomInit 教室初始化
type RoomInit struct {
	BuildingNumber uint8  `form:"building_number" json:"building_number"   binding:"required"`
	Floor          uint8  `form:"floor"           json:"floor"             binding:"required"`
	RoomNumber     uint64 `form:"room_number"     json:"room_number"       binding:"required"`
}

// Init 初始化教室
func (r *RoomInit) Init() *information.Response {
	count := 0
	var b model.Building
	model.DB.Where("building_number=?", r.BuildingNumber).Find(&b).Count(&count)
	if count == 0 {
		return &information.Response{
			Status: 50011,
			Msg:    "教学楼不存在",
		}
	}
	model.DB.Where("building_number=?", r.BuildingNumber).Find(&b)
	c := model.ClassRoom{
		Floor:      r.Floor,
		Buildings:  b,
		RoomNumber: r.RoomNumber,
	}
	return c.Create()
}

// RoomDelete 删除教室
type RoomDelete struct {
	RoomNumber uint64 `form:"room_number" json:"room_number" binding:"required"`
}

// Delete 删除教室
func (d *RoomDelete) Delete() *information.Response {
	var r model.ClassRoom
	count := 0
	model.DB.Where("room_number=?", d.RoomNumber).Find(&r).Count(&count)
	if count == 0 {
		return &information.Response{
			Status: 90002,
			Msg:    "教室不存在",
		}
	}
	model.DB.Where("room_number=?", d.RoomNumber).Delete(&model.ClassRoom{})
	return &information.Response{
		Status: 0,
		Msg:    "删除成功",
	}

}

// BuildingDelete 删除教学楼
type BuildingDelete struct {
	BuildingNumber uint8 `form:"building_number" json:"building_number" binding:"required"`
}

// Delete 删除教学楼
func (d *BuildingDelete) Delete() *information.Response {
	var m model.Building
	count := 0
	model.DB.Where("building_number=?", d.BuildingNumber).Find(&m).Count(&count)
	if count == 0 {
		return &information.Response{
			Status: 90003,
			Msg:    "教学楼不存在",
		}
	}
	model.DB.Where("building_number=?", d.BuildingNumber).Delete(&model.Building{})
	return &information.Response{
		Status: 0,
		Msg:    "教学楼删除成功",
	}

}
