package model

import "github.com/jinzhu/gorm"

// Building 是一个教学楼实体
type Building struct {
	gorm.Model
	BuildingNumber uint8 `gorm:"primary_key"`
}
