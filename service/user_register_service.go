package service

import (
	"Classroom-Management-System/information"
	"Classroom-Management-System/model"
)

// UserRegisterService 用户注册信息
type UserRegisterService struct {
	Username   string `form:"user_name"    json:"user_name"   binding:"required,min=7,max=20"`
	Password   string `form:"password"     json:"password"    binding:"required,min=6,max=20"`
	RePassword string `form:"repassword"   json:"repassword"  binding:"required,min=6,max=20"`
	Nickname   string `form:"nickname"     json:"nickname"    binding:"required,min=2,max=30"`
	Class      string `form:"class"        json:"class"       binding:"required,min=2,max=10"`
}

// Register 完成注册动作
func (u *UserRegisterService) Register() (model.User, *information.Response) {
	user := model.User{
		Nickname: u.Nickname,
		Username: u.Username,
		Class:    u.Class,
	}
	res := u.Valid()
	if res != nil {
		return user, res
	}
	if err := user.SetPassword(u.Password); err != nil {
		return user, &information.Response{
			Status: 401,
			Msg:    "密码加密失败",
		}
	}
	if err := model.DB.Create(&user).Error; err != nil {
		return user, &information.Response{
			Status: 402,
			Msg:    "数据库注册失败",
		}
	}
	return user, nil

}

// Valid 验证账号密码是否合法
func (u *UserRegisterService) Valid() *information.Response {
	if u.RePassword != u.Password {
		return &information.Response{
			Status: 401,
			Msg:    "两次密码不一致",
		}
	}
	countnumber := 0
	model.DB.Model(&model.User{}).Where("username=?", u.Username).Count(&countnumber)
	if countnumber > 0 {
		return &information.Response{
			Status: 401,
			Msg:    "用户名已被注册",
		}
	}
	return nil
}
