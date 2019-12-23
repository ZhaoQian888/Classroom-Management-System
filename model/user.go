package model

import (
	"Classroom-Management-System/information"
	"strconv"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

// User 是用户的信息
type User struct {
	gorm.Model
	Nickname string
	Username string
	Password string
	Class    string
	Email    string
	Identity uint64
}

// SetPassword 生成加密密码
func (u *User) SetPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 5)
	if err != nil {
		return err
	}
	u.Password = string(bytes)
	return nil
}

// CheckPassword 核对密码
func (u *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}

// CreateCookie 生成用户cookie
func (u *User) CreateCookie() string {
	head := u.Username
	random := time.Now().Nanosecond()
	tail := strconv.Itoa(random)
	cookie := strings.Join([]string{head, tail}, ",")
	return cookie
}

// GetUser 从cookie 提取用户信息
func GetUser(id interface{}) (User, error) {
	var u User
	err := DB.Where("identity=?", id).First(&u)
	return u, err.Error

}

//UserInfo 保存用户信息
type UserInfo struct {
	NickName string
	UserName string
	Class    string
	History  []ClassRoomOrder
	Email    string
}

// Info 返回用户信息
func (u *User) Info() *information.Response {
	var c ClassRoomOrder
	count := 0
	DB.Where("identity=?", u.Identity).Find(&c).Count(&count)
	if count == 0 {
		d := &UserInfo{
			NickName: u.Nickname,
			UserName: u.Username,
			Class:    u.Class,
			Email:    u.Email,
		}
		return &information.Response{
			Status: 0,
			Data:   d,
		}
	}
	var cs []ClassRoomOrder
	DB.Where("identity=?", u.Identity).Find(&cs)
	d := &UserInfo{
		NickName: u.Nickname,
		UserName: u.Username,
		Class:    u.Class,
		Email:    u.Email,
		History:  cs,
	}
	return &information.Response{
		Status: 0,
		Data:   d,
	}
}
