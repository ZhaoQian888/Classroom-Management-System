package model

import (
	"github.com/jinzhu/gorm"
)

// User 是用户的信息
type User struct {
	gorm.Model
	Nickname string
	Username string
	Password string
	Class    string
}
