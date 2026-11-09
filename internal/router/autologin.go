package router

import (
	"data-center/internal/handler/autologin"

	"github.com/gin-gonic/gin"
)

func initAutoLoginRouter(router *gin.RouterGroup) {
	autoLogin := router.Group("/autologin")
	autoLogin.GET("/", autologin.ListAutoLoginHandler)
	autoLogin.GET("/:id", autologin.GetAutoLoginHandler)
}
func initAdminAutoLoginRouter(router *gin.RouterGroup) {
	autoLogin := router.Group("/autologin")
	autoLogin.GET("/", autologin.ListAdminAutoLoginHandler)
	autoLogin.GET("/:id", autologin.GetAdminAutoLoginHandler)
	autoLogin.POST("/", autologin.CreateAdminAutoLoginHandler)
	autoLogin.PUT("/:id", autologin.UpdateAdminAutoLoginHandler)
	autoLogin.DELETE("/:id", autologin.DeleteAdminAutoLoginHandler)
}
