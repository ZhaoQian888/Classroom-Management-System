package middleware

import (
	"Classroom-Management-System/model"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// CurrentUser 中间件，将当前用户传下去
func CurrentUser() gin.HandlerFunc {

	return func(c *gin.Context) {
		s := sessions.Default(c)
		identity := s.Get("user_id")
		if identity != nil {
			user, err := model.GetUser(identity)
			if err == nil {
				c.Set("user", &user)
			}
		}
		c.Next()
	}

}
