package information

import "Classroom-Management-System/model"

// User 是保存用户的信息
type User struct {
	Nickname string `form:"nickname" json:"username" binding:"required"`
	Class    string `form:"class" json:"class" bingding:"required"`
}

// UserResponse 用户回应序列化
type UserResponse struct {
	Status int         `json:"status"`
	Data   interface{} `json:"data"`
	Msg    string      `json:"msg"`
	Error  string      `json:"error"`
	data   User
}

// BuildUser 构建用户信息
func BuildUser(u model.User) User {
	return User{
		Nickname: u.Nickname,
		Class:    u.Class,
	}
}

// CreateUserRseponse 生成用户回应信息
func CreateUserRseponse(u model.User) UserResponse {
	return UserResponse{
		Status: 1344,
		Msg:    "用户信息",
		Data:   BuildUser(u),
	}
}
