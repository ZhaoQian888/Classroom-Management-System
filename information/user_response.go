package information

import "Classroom-Management-System/model"

type User struct {
	nickname string `form:"nickname" JSON:"username" binding:"required"`
	class    string `form:"class" JSON:"class" bingding:"required"`
}
type UserResponse struct {
	Response
	data User
}

func BuildUser(u model.User) User {
	return User{
		nickname: u.Nickname,
		class:    u.Class,
	}
}
func CreateUserRseponse(u model.User) UserResponse {
	return UserResponse{
		data: BuildUser(u),
	}
}
