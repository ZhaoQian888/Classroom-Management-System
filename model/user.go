package model

import (
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
	Identity int
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

// f err := model.DB.Where("username=?", u.Username).First(&user).Error; err != nil {
// 	return user, &information.Response{
// 		Status: 11003,
// 		Msg:    "用户名或者密码错误",
// 	}
// }
