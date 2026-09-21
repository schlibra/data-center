package router

import (
	"data-center/internal/handler"

	"github.com/gin-gonic/gin"
)

func initAdminUserRouter(router *gin.RouterGroup) {
	user := router.Group("/user")
	user.GET("/", handler.ListAdminUserHandler)
	user.GET("/:id", handler.GetAdminUserHandler)
	user.POST("/", handler.CreateAdminUserHandler)
	user.PUT("/:id", handler.UpdateAdminUserHandler)
	user.PATCH("/:id", handler.PasswordAdminUserHandler)
	user.DELETE("/:id", handler.DeleteAdminUserHandler)
}

func initAdminRouter(router *gin.RouterGroup) {
	admin := router.Group("/admin")
	initAdminUserRouter(admin)
}
