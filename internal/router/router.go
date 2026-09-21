package router

import (
	"data-center/internal/handler"

	"github.com/gin-gonic/gin"
)

func initUserRouter(router *gin.RouterGroup) {
	user := router.Group("/user")
	user.POST("/login", handler.UserLoginHandler)
	user.PUT("/login", handler.UserLoginKeyHandler)
	user.POST("/register", handler.UserRegisterHandler)
	user.PUT("/register", handler.UserRegisterKeyHandler)
	user.POST("/logout", handler.UserLogoutHandler)
	user.GET("/", handler.UserInfoHandler)
	user.PUT("/", handler.UserUpdateHandler)
	user.PATCH("/", handler.UserPasswordHandler)
}

func InitRouter(router *gin.Engine) {
	api := router.Group("/api")
	initUserRouter(api)
	initFrpRouter(api)
}
