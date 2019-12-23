package middleware

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

// Session 初始化
func Session(secret string) gin.HandlerFunc {
	store := cookie.NewStore([]byte(secret))
	store.Options(sessions.Options{HttpOnly: true, MaxAge: 12000, Path: "/"})
	return sessions.Sessions("gin-session", store)
}
