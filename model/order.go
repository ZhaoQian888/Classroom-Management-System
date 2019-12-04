package model

import "github.com/jinzhu/gorm"

// ClassRoomOrder 是一个预定关系实体
type ClassRoomOrder struct {
	gorm.Model
	UserName     string
	RoomUser     string
	OrderTime    int64
	UseTimeStart int64
	UseTimeEnd   int64
}
