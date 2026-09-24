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
func initAdminGroupRouter(router *gin.RouterGroup) {
	group := router.Group("/group")
	group.GET("/", handler.ListAdminGroupHandler)
	group.GET("/:id", handler.GetAdminGroupHandler)
	group.POST("/", handler.CreateAdminGroupHandler)
	group.PUT("/:id", handler.UpdateAdminGroupHandler)
	group.DELETE("/:id", handler.DeleteAdminGroupHandler)
}
func initAdminPermissionRouter(router *gin.RouterGroup) {
	permission := router.Group("/permission")
	permission.GET("/", handler.ListAdminPermissionsHandler)
	permission.POST("/", handler.CreateAdminPermissionHandler)
	permission.GET("/:id", handler.GetAdminPermissionHandler)
	permission.PUT("/:id", handler.UpdateAdminPermissionHandler)
	permission.DELETE("/:id", handler.DeleteAdminPermissionHandler)
}

func initAdminRouter(router *gin.RouterGroup) {
	admin := router.Group("/admin")
	initAdminUserRouter(admin)
	initAdminGroupRouter(admin)
	initAdminPermissionRouter(admin)
}
