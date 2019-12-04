package service

import (
	"Classroom-Management-System/information"
	"Classroom-Management-System/model"
	"fmt"
)

// UserInfo 用户的信息
type UserInfo struct {
	Username string `form:"user_name" json:"user_name" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

// UserLogInService 用户的登陆信息
type UserLogInService struct {
	UserInfo
}

// CheckUserValid 可以检查用户名是否合法
func (u *UserLogInService) CheckUserValid() bool {
	//此处处理用户信息是否合法
	if len(u.Username) > 20 || len(u.Username) > 7 || len(u.Password) < 6 || len(u.Password) > 20 {
		return false
	}
	return true
}

// LogIn 此处返回数据库中用户信息与
func (u *UserLogInService) LogIn() (model.User, *information.Response) {
	var user model.User
	// 此处取出数据库中的用户信息，核对密码与信息然后返回一个用户模型
	if err := model.DB.Where("username = ?", u.Username).First(&user).Error; err != nil {
		fmt.Println(err)
		return user, &information.Response{
			Status: 11003,
			Msg:    "用户名或者密码错误",
		}
	}
	if user.CheckPassword(u.Password) {
		return user, &information.Response{
			Status: 0,
			Msg:    "账号密码正确",
		}
	}
	return user, &information.Response{
		Status: 11004,
		Msg:    "账号或者密码错误",
	}
}
