package model

import (
	"github.com/jinzhu/gorm"
)

// ClassRoom 是一个教室实体
type ClassRoom struct {
	gorm.Model
	BuildingName   string
	BuildingNumber uint8
	Floor          uint8
	RoomNumber     uint64
}
