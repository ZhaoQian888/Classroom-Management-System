package service

import (
	"Classroom-Management-System/information"
	"Classroom-Management-System/model"
	"time"
)

// UserRegisterService 用户注册信息
type UserRegisterService struct {
	Username   string `form:"user_name"    json:"user_name"   binding:"required,min=7,max=20"`
	Password   string `form:"password"     json:"password"    binding:"required,min=6,max=20"`
	RePassword string `form:"repassword"   json:"repassword"  binding:"required,min=6,max=20"`
	Nickname   string `form:"nickname"     json:"nickname"    binding:"required,min=2,max=30"`
	Class      string `form:"class"        json:"class"       binding:"required,min=2,max=30"`
	Email      string `form:"email" 	   json:"email"       binding:"required,min=6,max=25"`
}

// |information|required|
// |---------|-----------|
// |user_name|必要，最短7个字符，最长20个字符|
// |passoword|必要，最短6个字符，最长20个字符|
// |nickname|必要，最短2个字符，最长30个字符|
// |class|必要，最短2个字符，最长30个字符|
// |email|必要，最短6个字符，最长25个字符|

// CreateIdentity 生成验证id
func CreateIdentity() uint64 {
	id := time.Now().Nanosecond()
	for true {
		count := 0
		model.DB.Model(&model.User{}).Where("identity=?", id).Count(&count)
		if count == 0 {
			break
		}
	}
	return uint64(id)

}

// Register 完成注册动作
func (u *UserRegisterService) Register() (model.User, *information.Response) {
	user := model.User{
		Nickname: u.Nickname,
		Username: u.Username,
		Class:    u.Class,
		Email:    u.Email,
		Identity: CreateIdentity(),
	}
	res := u.Valid()
	if res != nil {
		return user, res
	}
	if err := user.SetPassword(u.Password); err != nil {
		return user, &information.Response{
			Status: 10001,
			Msg:    "密码加密失败",
		}
	}
	if err := model.DB.Create(&user).Error; err != nil {
		return user, &information.Response{
			Status: 10002,
			Msg:    "数据库注册失败",
		}
	}
	return user, nil

}

// Valid 验证账号密码是否合法
func (u *UserRegisterService) Valid() *information.Response {
	if u.RePassword != u.Password {
		return &information.Response{
			Status: 10003,
			Msg:    "两次密码不一致",
		}
	}
	countnumber := 0
	model.DB.Model(&model.User{}).Where("username=?", u.Username).Count(&countnumber)
	if countnumber > 0 {
		return &information.Response{
			Status: 10004,
			Msg:    "用户名已被注册",
		}
	}
	return nil
}
