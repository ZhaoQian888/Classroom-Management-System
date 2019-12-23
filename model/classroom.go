package model

import (
	"github.com/jinzhu/gorm"
)

// ClassRoom 是一个教室实体
type ClassRoom struct {
	gorm.Model
	BuildingName   string `gorm:""`
	BuildingNumber uint8  `gorm:""`
	Floor          uint8  `gorm:"not null"`
	RoomNumber     uint64 `gorm:"primary_key;not null;index:addr"`
}
