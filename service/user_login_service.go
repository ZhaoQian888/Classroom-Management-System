package service

import (
	"Classroom-Management-System/model"
)

// UserInfo 用户的信息
type UserInfo struct {
	Username string `form:"use_rname" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

// UserLogInService 用户的登陆信息
type UserLogInService struct {
	UserInfo
}

// CheckUserValid 可以检查用户名是否合法
func (u *UserLogInService) CheckUserValid() bool {
	//此处处理用户信息是否合法
	return true
}

// LogIn 此处返回数据库中用户信息与
func (u *UserLogInService) LogIn() (model.User, error) {
	var user model.User
	// 此处取出数据库中的用户信息，核对密码与信息然后返回一个用户模型
	return user, nil
}
