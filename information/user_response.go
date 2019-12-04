package information

import "Classroom-Management-System/model"

// User 是保存用户的信息
type User struct {
	Username string `form:"user_name" json:"user_name" binding:"required"`
	Nickname string `form:"nickname" json:"nickname" binding:"required"`
	Class    string `form:"class" json:"class" bingding:"required"`
	Email    string `form:"email" json:"email" bingding:"required"`
}

// UserResponse 用户回应序列化
type UserResponse struct {
	Status int         `json:"status"`
	Data   interface{} `json:"data"`
	Msg    string      `json:"msg"`
	Error  string      `json:"error"`
}

// BuildUser 构建用户信息
func BuildUser(u model.User) User {
	return User{
		Nickname: u.Nickname,
		Class:    u.Class,
		Username: u.Username,
		Email:    u.Email,
	}
}

// CreateUserRseponse 生成用户回应信息
func CreateUserRseponse(u model.User) UserResponse {
	return UserResponse{
		Status: 0,
		Msg:    "用户信息",
		Data:   BuildUser(u),
	}
}
